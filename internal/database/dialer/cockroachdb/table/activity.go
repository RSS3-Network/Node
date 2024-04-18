package table

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/shopspring/decimal"
)

const (
	// ActivitySpamLimit this limit of actions in a single activity, mark as spam if exceeds.
	ActivitySpamLimit = 100
)

type Activity struct {
	ID           string          `gorm:"column:id"`
	Network      network.Network `gorm:"column:network"`
	Platform     string          `gorm:"column:platform"`
	Index        uint            `gorm:"column:index"`
	From         string          `gorm:"column:from"`
	To           string          `gorm:"column:to"`
	Tag          tag.Tag         `gorm:"column:tag"`
	Type         string          `gorm:"column:type"`
	Status       bool            `gorm:"column:status"`
	Fee          *Fee            `gorm:"column:fee;type:jsonb"`
	TotalActions uint            `gorm:"column:total_actions"`
	Actions      ActivityActions `gorm:"column:actions;type:jsonb"`
	Timestamp    time.Time       `gorm:"column:timestamp"`
	CreatedAt    time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time       `gorm:"column:updated_at;autoUpdateTime"`
}

func (f *Activity) TableName() string {
	return "activities"
}

func (f *Activity) PartitionName(_activity *activity.Activity) string {
	if _activity != nil {
		f.Timestamp = time.Unix(int64(_activity.Timestamp), 0)
		f.Network = _activity.Network
	}

	return fmt.Sprintf("%s_%s_%d_q%d", f.TableName(), f.Network,
		f.Timestamp.Year(), int(math.Ceil(float64(f.Timestamp.Month())/3)))
}

func (f *Activity) Import(_activity *activity.Activity) error {
	f.ID = _activity.ID
	f.Network = _activity.Network
	f.Platform = _activity.Platform
	f.Index = _activity.Index
	f.From = _activity.From
	f.To = _activity.To
	f.Tag = _activity.Type.Tag()
	f.Type = _activity.Type.Name()
	f.Status = _activity.Status
	f.TotalActions = uint(len(_activity.Actions))
	f.Actions = make(ActivityActions, 0)
	f.Timestamp = time.Unix(int64(_activity.Timestamp), 0)

	if _activity.Fee != nil {
		f.Fee = new(Fee)

		if err := f.Fee.Import(_activity.Fee); err != nil {
			return fmt.Errorf("invalid fee: %w", err)
		}
	}

	// spam transactions only retain last one action
	if f.TotalActions > ActivitySpamLimit {
		for i := len(_activity.Actions) - 1; i >= 0; i-- {
			item := new(ActivityAction)

			if err := item.Import(_activity.Actions[i]); err != nil {
				return err
			}

			if f.Tag.String() == item.Tag && f.Type == item.Type {
				f.Actions = append(f.Actions, item)

				return nil
			}
		}
	}

	for _, action := range _activity.Actions {
		item := new(ActivityAction)

		if err := item.Import(action); err != nil {
			return fmt.Errorf("invalid action: %w", err)
		}

		f.Actions = append(f.Actions, item)
	}

	return nil
}

func (f *Activity) Export(index *Index) (*activity.Activity, error) {
	_activity := activity.Activity{
		ID:           f.ID,
		From:         f.From,
		To:           f.To,
		Network:      f.Network,
		Platform:     f.Platform,
		Status:       f.Status,
		Actions:      make([]*activity.Action, 0, len(f.Actions)),
		TotalActions: f.TotalActions,
		Timestamp:    uint64(f.Timestamp.Unix()),
		Owner:        index.Owner,
		Direction:    index.Direction,
	}

	var err error

	if _activity.Type, err = schema.TypeString(f.Tag, f.Type); err != nil {
		return nil, err
	}

	_activity.Tag = _activity.Type.Tag()

	if _activity.Fee, err = f.Fee.Export(); err != nil {
		return nil, fmt.Errorf("invalid fee: %w", err)
	}

	for _, action := range f.Actions {
		item, err := action.Export()
		if err != nil {
			return nil, err
		}

		_activity.Actions = append(_activity.Actions, item)
	}

	return &_activity, nil
}

type Activities []*Activity

func (f *Activities) Import(activities []*activity.Activity) error {
	for _, _activity := range activities {
		item := new(Activity)

		if err := item.Import(_activity); err != nil {
			return err
		}

		*f = append(*f, item)
	}

	return nil
}

func (f *Activities) Export(indexes []*Index) ([]*activity.Activity, error) {
	activities := make(map[string]*Activity)

	for _, _activity := range *f {
		activities[_activity.ID] = _activity
	}

	result := make([]*activity.Activity, 0, len(indexes))

	for _, index := range indexes {
		if _activity, ok := activities[index.ID]; ok {
			data, err := _activity.Export(index)
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
func (f *Fee) Import(fee *activity.Fee) error {
	if fee != nil {
		f.Address = fee.Address
		f.Amount = fee.Amount
		f.Decimal = fee.Decimal
	}

	return nil
}

//goland:noinspection ALL
func (f *Fee) Export() (*activity.Fee, error) {
	if f == nil {
		return nil, nil
	}

	return &activity.Fee{
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

type ActivityAction struct {
	Tag      string          `json:"tag"`
	Type     string          `json:"type"`
	From     string          `json:"from"`
	To       string          `json:"to"`
	Platform string          `json:"platform,omitempty"`
	Metadata json.RawMessage `json:"metadata"`
}

func (f *ActivityAction) Import(action *activity.Action) (err error) {
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

func (f *ActivityAction) Export() (*activity.Action, error) {
	action := &activity.Action{
		From:     f.From,
		To:       f.To,
		Platform: f.Platform,
	}

	var err error
	if action.Tag, action.Type, err = schema.TagAndTypeString(f.Tag, f.Type); err != nil {
		return nil, err
	}

	if action.Metadata, err = metadata.Unmarshal(action.Type, f.Metadata); err != nil {
		return nil, err
	}

	return action, nil
}

var (
	_ sql.Scanner   = (*ActivityActions)(nil)
	_ driver.Valuer = (*ActivityActions)(nil)
)

type ActivityActions []*ActivityAction

//goland:noinspection ALL
func (f *ActivityActions) Scan(value any) error {
	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid type: %T", value)
	}

	return json.Unmarshal(data, &f)
}

//goland:noinspection ALL
func (f ActivityActions) Value() (driver.Value, error) {
	return json.Marshal(f)
}
