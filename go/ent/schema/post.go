package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id"),
		field.String("title"),
		field.String("content"),
		field.Uint32("version"),
		field.Time("createTime"),
		field.Time("lastUpdate"),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return nil
}
