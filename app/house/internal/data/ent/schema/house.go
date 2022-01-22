package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// House holds the schema definition for the House entity.
type House struct {
	ent.Schema
}

// Fields of the House.
func (House) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Float("Price"),
		field.String("title").MinLen(1),
		field.String("community"),
		field.String("image").Default(""),
		field.Int32("toilet_count").Default(0),
		field.Int32("kitchen_count").Default(0),
		field.Int32("floor_count").Default(0),
		field.Int32("hall_count").Default(0),
		field.Int32("room_count").Default(0),
	}
}

// Edges of the House.
func (House) Edges() []ent.Edge {
	return nil
}
