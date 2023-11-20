package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/shopspring/decimal"
)

// Feed holds the schema definition for the Feed entity.
type Feed struct {
	ent.Schema
}

// Fields of the Feed.
func (Feed) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique(),
		field.String("chain").NotEmpty(),
		field.String("platform"),
		field.String("from").NotEmpty(),
		field.String("to").NotEmpty(),
		field.String("tag").NotEmpty(),
		field.String("type").NotEmpty(),
		field.String("status").NotEmpty(),
		field.Int("index").Default(0),
		field.Int("total_actions").Default(0),
		field.JSON("actions", []schema.Action{}),
		field.Float("fee_value").GoType(decimal.Decimal{}),
		field.String("fee_token"),
		field.Time("timestamp"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Indexes of the Feed.
func (Feed) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("platform"),
		index.Fields("tag", "type"),
		index.Fields("timestamp"),
		index.Fields("total_actions"),
	}
}
