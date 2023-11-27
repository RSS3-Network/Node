package table

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/shopspring/decimal"
)

var _ schema.FeedTransformer = (*Feed)(nil)

type Feed struct {
	ID           string           `gorm:"column:id"`
	Chain        string           `gorm:"column:chain"`
	Platform     *filter.Platform `gorm:"column:platform"`
	Index        uint             `gorm:"column:index"`
	From         string           `gorm:"column:from"`
	To           string           `gorm:"column:to"`
	Tag          filter.Tag       `gorm:"column:tag"`
	Type         string           `gorm:"column:type"`
	Status       bool             `gorm:"column:status"`
	Fee          Fee              `gorm:"column:fee;jsonb"`
	TotalActions int              `gorm:"column:total_actions"`
	Actions      FeedActions      `gorm:"column:actions;type:jsonb"`
	Timestamp    time.Time        `gorm:"column:timestamp"`
	Version      string           `gorm:"column:version"`
	CreatedAt    time.Time        `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time        `gorm:"column:updated_at;autoUpdateTime"`
}

func (f *Feed) TableName() string {
	return "feeds"
}

func (f *Feed) PartitionName(feed *schema.Feed) string {
	if feed != nil {
		f.Timestamp = time.Unix(int64(feed.Timestamp), 0)
	}

	return fmt.Sprintf("%s_%s_%d_q%d", f.TableName(), strings.Replace(f.Chain, ".", "_", -1),
		f.Timestamp.Year(), int(math.Ceil(float64(f.Timestamp.Month())/3)))
}

func (f *Feed) Import(feed schema.Feed) error {
	f.ID = feed.ID
	f.Chain = feed.Chain.FullName()
	f.Platform = feed.Platform
	f.Index = feed.Index
	f.From = feed.From
	f.To = feed.To
	f.Tag = feed.Type.Tag()
	f.Type = feed.Type.Name()
	f.Status = feed.Status
	f.TotalActions = feed.TotalActions
	f.Actions = make(FeedActions, 0)
	f.Timestamp = time.Unix(int64(feed.Timestamp), 0)
	f.Version = feed.Version

	// TODO: spam

	if err := f.Fee.Import(feed.Fee); err != nil {
		return fmt.Errorf("invalid fee: %w", err)
	}

	for _, action := range feed.Actions {
		var item FeedAction
		if err := item.Import(action); err != nil {
			return fmt.Errorf("invalid action: %w", err)
		}

		f.Actions = append(f.Actions, item)
	}

	return nil
}

var _ schema.FeedsTransformer = (*Feeds)(nil)

type Feeds []Feed

func (f *Feeds) Import(feeds []schema.Feed) error {
	for _, feed := range feeds {
		var item Feed
		if err := item.Import(feed); err != nil {
			return err
		}

		*f = append(*f, item)
	}

	return nil
}

var _ schema.FeeTransformer = (*Fee)(nil)

type Fee struct {
	Address *string         `json:"address,omitempty"`
	Amount  decimal.Decimal `json:"amount"`
	Decimal uint            `json:"decimal"`
}

func (f *Fee) Import(fee schema.Fee) error {
	f.Address = fee.Address
	f.Amount = fee.Amount
	f.Decimal = fee.Decimal

	return nil
}

var (
	_ sql.Scanner   = (*Fee)(nil)
	_ driver.Valuer = (*Fee)(nil)
)

//goland:noinspection ALL
func (f *Fee) Scan(value any) error {
	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid type: %T", value)
	}

	return json.Unmarshal(data, &f)
}

//goland:noinspection ALL
func (f Fee) Value() (driver.Value, error) {
	return json.Marshal(f)
}

var _ schema.ActionTransformer = (*FeedAction)(nil)

type FeedAction struct {
	Tag      string          `json:"tag"`
	Type     string          `json:"type"`
	From     string          `json:"from"`
	To       string          `json:"to"`
	Platform string          `json:"platform,omitempty"`
	Metadata json.RawMessage `json:"metadata"`
}

func (f *FeedAction) Import(action schema.Action) (err error) {
	f.Tag = action.Type.Tag().String()
	f.Type = action.Type.Name()
	f.From = action.From
	f.To = action.To
	f.Platform = action.Platform

	if f.Metadata, err = json.Marshal(action.Metadata); err != nil {
		return fmt.Errorf("invalid metadata: %w", err)
	}

	return nil
}

var (
	_ sql.Scanner   = (*FeedActions)(nil)
	_ driver.Valuer = (*FeedActions)(nil)
)

type FeedActions []FeedAction

//goland:noinspection ALL
func (f *FeedActions) Scan(value any) error {
	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid type: %T", value)
	}

	return json.Unmarshal(data, &f)
}

//goland:noinspection ALL
func (f FeedActions) Value() (driver.Value, error) {
	return json.Marshal(f)
}
