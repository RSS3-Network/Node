package table

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
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
	Calldata     *Calldata       `gorm:"column:calldata;type:jsonb"`
	TotalActions uint            `gorm:"column:total_actions"`
	Actions      ActivityActions `gorm:"column:actions;type:jsonb"`
	Timestamp    time.Time       `gorm:"column:timestamp"`
	CreatedAt    time.Time       `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time       `gorm:"column:updated_at;autoUpdateTime"`
}

func (f *Activity) TableName() string {
	return "activities"
}

func (f *Activity) PartitionName(activity *activityx.Activity) string {
	if activity != nil {
		f.Timestamp = time.Unix(int64(activity.Timestamp), 0)
		f.Network = activity.Network
	}

	return fmt.Sprintf("%s_%s_%d_q%d", f.TableName(), f.Network,
		f.Timestamp.Year(), int(math.Ceil(float64(f.Timestamp.Month())/3)))
}

func (f *Activity) Import(activity *activityx.Activity) error {
	f.ID = activity.ID
	f.Network = activity.Network
	f.Platform = activity.Platform
	f.Index = activity.Index
	f.From = activity.From
	f.To = activity.To
	f.Tag = activity.Type.Tag()
	f.Type = activity.Type.Name()
	f.Status = activity.Status
	f.TotalActions = uint(len(activity.Actions))
	f.Actions = make(ActivityActions, 0)
	f.Timestamp = time.Unix(int64(activity.Timestamp), 0)

	if activity.Fee != nil {
		f.Fee = new(Fee)

		if err := f.Fee.Import(activity.Fee); err != nil {
			return fmt.Errorf("invalid fee: %w", err)
		}
	}

	if activity.Calldata != nil {
		f.Calldata = new(Calldata)

		if err := f.Calldata.Import(activity.Calldata); err != nil {
			return fmt.Errorf("invalid calldata: %w", err)
		}
	}

	// spam transactions only retain last one action
	if f.TotalActions > ActivitySpamLimit {
		for i := len(activity.Actions) - 1; i >= 0; i-- {
			item := new(ActivityAction)

			if err := item.Import(activity.Actions[i]); err != nil {
				return err
			}

			if f.Tag.String() == item.Tag && f.Type == item.Type {
				f.Actions = append(f.Actions, item)

				return nil
			}
		}
	}

	for _, action := range activity.Actions {
		item := new(ActivityAction)

		if err := item.Import(action); err != nil {
			return fmt.Errorf("invalid action: %w", err)
		}

		f.Actions = append(f.Actions, item)
	}

	return nil
}

func (f *Activity) Export(index *Index) (*activityx.Activity, error) {
	activity := activityx.Activity{
		ID:           f.ID,
		From:         f.From,
		To:           f.To,
		Network:      f.Network,
		Platform:     f.Platform,
		Status:       f.Status,
		Actions:      make([]*activityx.Action, 0, len(f.Actions)),
		TotalActions: f.TotalActions,
		Timestamp:    uint64(f.Timestamp.Unix()),
		Owner:        index.Owner,
		Direction:    index.Direction,
	}

	var err error

	if activity.Type, err = schema.ParseTypeFromString(f.Tag, f.Type); err != nil {
		return nil, err
	}

	activity.Tag = activity.Type.Tag()

	if activity.Fee, err = f.Fee.Export(); err != nil {
		return nil, fmt.Errorf("invalid fee: %w", err)
	}

	if activity.Calldata, err = f.Calldata.Export(); err != nil {
		return nil, fmt.Errorf("invalid calldata: %w", err)
	}

	for _, action := range f.Actions {
		item, err := action.Export()
		if err != nil {
			return nil, err
		}

		activity.Actions = append(activity.Actions, item)
	}

	return &activity, nil
}

type Activities []*Activity

func (f *Activities) Import(activities []*activityx.Activity) error {
	for _, activity := range activities {
		item := new(Activity)

		if err := item.Import(activity); err != nil {
			return err
		}

		*f = append(*f, item)
	}

	return nil
}

func (f *Activities) Export(indexes []*Index) ([]*activityx.Activity, error) {
	activities := make(map[string]*Activity)

	for _, activity := range *f {
		activities[activity.ID] = activity
	}

	result := make([]*activityx.Activity, 0, len(indexes))

	for _, index := range indexes {
		if activity, ok := activities[index.ID]; ok {
			data, err := activity.Export(index)
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
func (f *Fee) Import(fee *activityx.Fee) error {
	if fee != nil {
		f.Address = fee.Address
		f.Amount = fee.Amount
		f.Decimal = fee.Decimal
	}

	return nil
}

//goland:noinspection ALL
func (f *Fee) Export() (*activityx.Fee, error) {
	if f == nil {
		return nil, nil
	}

	return &activityx.Fee{
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

type Calldata struct {
	Raw            string `json:"raw,omitempty"`
	FunctionHash   string `json:"function_hash,omitempty"`
	ParsedFunction string `json:"parsed_function,omitempty"`
}

//goland:noinspection ALL
func (c *Calldata) Import(calldata *activityx.Calldata) error {
	if calldata != nil {
		c.Raw = calldata.Raw
		c.FunctionHash = calldata.FunctionHash
		c.ParsedFunction = calldata.ParsedFunction
	}

	return nil
}

//goland:noinspection ALL
func (c *Calldata) Export() (*activityx.Calldata, error) {
	if c == nil {
		return nil, nil
	}

	return &activityx.Calldata{
		Raw:            c.Raw,
		FunctionHash:   c.FunctionHash,
		ParsedFunction: c.ParsedFunction,
	}, nil
}

var (
	_ sql.Scanner   = (*Calldata)(nil)
	_ driver.Valuer = (*Calldata)(nil)
)

//goland:noinspection ALL
func (c *Calldata) Scan(value any) error {
	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("invalid type: %T", value)
	}

	return json.Unmarshal(data, &c)
}

//goland:noinspection ALL
func (c Calldata) Value() (driver.Value, error) {
	return json.Marshal(c)
}

type ActivityAction struct {
	Tag      string          `json:"tag"`
	Type     string          `json:"type"`
	From     string          `json:"from"`
	To       string          `json:"to"`
	Platform string          `json:"platform,omitempty"`
	Metadata json.RawMessage `json:"metadata"`
}

func (f *ActivityAction) Import(action *activityx.Action) (err error) {
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

func (f *ActivityAction) Export() (*activityx.Action, error) {
	action := &activityx.Action{
		From:     f.From,
		To:       f.To,
		Platform: f.Platform,
	}

	var err error
	if action.Tag, action.Type, err = schema.ParseTagAndTypeFromString(f.Tag, f.Type); err != nil {
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
