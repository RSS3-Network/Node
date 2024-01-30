package table

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/shopspring/decimal"
)

const (
	// FeedSpamLimit The restriction conditions for judging the feed to be spam cannot be changed at will.
	FeedSpamLimit = 100
)

type Feed struct {
	ID           string           `gorm:"column:id"`
	Network      filter.Network   `gorm:"column:network"`
	Platform     *filter.Platform `gorm:"column:platform"`
	Index        uint             `gorm:"column:index"`
	From         string           `gorm:"column:from"`
	To           string           `gorm:"column:to"`
	Tag          filter.Tag       `gorm:"column:tag"`
	Type         string           `gorm:"column:type"`
	Status       bool             `gorm:"column:status"`
	Fee          *Fee             `gorm:"column:fee;type:jsonb"`
	TotalActions uint             `gorm:"column:total_actions"`
	Actions      FeedActions      `gorm:"column:actions;type:jsonb"`
	Timestamp    time.Time        `gorm:"column:timestamp"`
	CreatedAt    time.Time        `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time        `gorm:"column:updated_at;autoUpdateTime"`
}

func (f *Feed) TableName() string {
	return "feeds"
}

func (f *Feed) PartitionName(feed *schema.Feed) string {
	if feed != nil {
		f.Timestamp = time.Unix(int64(feed.Timestamp), 0)
		f.Network = feed.Network
	}

	return fmt.Sprintf("%s_%s_%d_q%d", f.TableName(), f.Network,
		f.Timestamp.Year(), int(math.Ceil(float64(f.Timestamp.Month())/3)))
}

func (f *Feed) Import(feed *schema.Feed) error {
	f.ID = feed.ID
	f.Network = feed.Network
	f.Platform = feed.Platform
	f.Index = feed.Index
	f.From = feed.From
	f.To = feed.To
	f.Tag = feed.Type.Tag()
	f.Type = feed.Type.Name()
	f.Status = feed.Status
	f.TotalActions = uint(len(feed.Actions))
	f.Actions = make(FeedActions, 0)
	f.Timestamp = time.Unix(int64(feed.Timestamp), 0)

	if feed.Fee != nil {
		f.Fee = new(Fee)

		if err := f.Fee.Import(feed.Fee); err != nil {
			return fmt.Errorf("invalid fee: %w", err)
		}
	}

	// spam transactions only retain last one action
	if f.TotalActions > FeedSpamLimit {
		for i := len(feed.Actions) - 1; i >= 0; i-- {
			item := new(FeedAction)

			if err := item.Import(feed.Actions[i]); err != nil {
				return err
			}

			if f.Tag.String() == item.Tag && f.Type == item.Type {
				f.Actions = append(f.Actions, item)

				return nil
			}
		}
	}

	for _, action := range feed.Actions {
		item := new(FeedAction)

		if err := item.Import(action); err != nil {
			return fmt.Errorf("invalid action: %w", err)
		}

		f.Actions = append(f.Actions, item)
	}

	return nil
}

func (f *Feed) Export(index *Index) (*schema.Feed, error) {
	feed := schema.Feed{
		ID:           f.ID,
		From:         f.From,
		To:           f.To,
		Network:      f.Network,
		Platform:     f.Platform,
		Status:       f.Status,
		Actions:      make([]*schema.Action, 0, len(f.Actions)),
		TotalActions: f.TotalActions,
		Timestamp:    uint64(f.Timestamp.Unix()),
		Owner:        index.Owner,
		Direction:    index.Direction,
	}

	var err error

	if feed.Type, err = filter.TypeString(f.Tag, f.Type); err != nil {
		return nil, err
	}

	feed.Tag = feed.Type.Tag()

	if feed.Fee, err = f.Fee.Export(); err != nil {
		return nil, fmt.Errorf("invalid fee: %w", err)
	}

	for _, action := range f.Actions {
		item, err := action.Export()
		if err != nil {
			return nil, err
		}

		feed.Actions = append(feed.Actions, item)
	}

	return &feed, nil
}

type Feeds []*Feed

func (f *Feeds) Import(feeds []*schema.Feed) error {
	for _, feed := range feeds {
		item := new(Feed)

		if err := item.Import(feed); err != nil {
			return err
		}

		*f = append(*f, item)
	}

	return nil
}

func (f *Feeds) Export(indexes []*Index) ([]*schema.Feed, error) {
	feeds := make(map[string]*Feed)

	for _, feed := range *f {
		feeds[feed.ID] = feed
	}

	result := make([]*schema.Feed, 0, len(indexes))

	for _, index := range indexes {
		if feed, ok := feeds[index.ID]; ok {
			data, err := feed.Export(index)
			if err != nil {
				return nil, err
			}

			result = append(result, data)
		}
	}

	return result, nil
}

type Fee struct {
	Address *string         `json:"address,omitempty"`
	Amount  decimal.Decimal `json:"amount"`
	Decimal uint            `json:"decimal"`
}

//goland:noinspection ALL
func (f *Fee) Import(fee *schema.Fee) error {
	if fee != nil {
		f.Address = fee.Address
		f.Amount = fee.Amount
		f.Decimal = fee.Decimal
	}

	return nil
}

//goland:noinspection ALL
func (f *Fee) Export() (*schema.Fee, error) {
	if f == nil {
		return nil, nil
	}

	return &schema.Fee{
		Address: f.Address,
		Amount:  f.Amount,
		Decimal: f.Decimal,
	}, nil
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

type FeedAction struct {
	Tag      string          `json:"tag"`
	Type     string          `json:"type"`
	From     string          `json:"from"`
	To       string          `json:"to"`
	Platform string          `json:"platform,omitempty"`
	Metadata json.RawMessage `json:"metadata"`
}

func (f *FeedAction) Import(action *schema.Action) (err error) {
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

func (f *FeedAction) Export() (*schema.Action, error) {
	action := &schema.Action{
		From:     f.From,
		To:       f.To,
		Platform: f.Platform,
	}

	var err error
	if action.Tag, action.Type, err = filter.TagAndTypeString(f.Tag, f.Type); err != nil {
		return nil, err
	}

	if action.Metadata, err = metadata.Unmarshal(action.Type, f.Metadata); err != nil {
		return nil, err
	}

	return action, nil
}

var (
	_ sql.Scanner   = (*FeedActions)(nil)
	_ driver.Valuer = (*FeedActions)(nil)
)

type FeedActions []*FeedAction

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
