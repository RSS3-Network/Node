package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Index holds the schema definition for the Index entity.
type Index struct {
	ent.Schema
}

// Fields of the Index.
func (Index) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("uuid").StructTag(`json:"uuid"`),
		field.String("feed_id").StructTag(`json:"id"`),
		field.String("owner").NotEmpty(),
		field.String("chain").NotEmpty(),
		field.String("platform"),
		field.String("tag").NotEmpty(),
		field.String("type").NotEmpty(),
		field.String("status").NotEmpty(),
		field.String("direction").NotEmpty(),
		field.Uint("index").Default(0),
		field.Time("timestamp"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Index of the Index Table.
func (Index) Index() []ent.Index {
	return []ent.Index{
		index.Fields("timestamp", "index"),
		index.Fields("platform", "timestamp", "index"),
		index.Fields("tag", "timestamp", "index", "type"),
		index.Fields("chain", "timestamp", "index"),
		index.Fields("owner", "timestamp", "index", "direction"),
	}
}
