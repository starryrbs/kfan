// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// HousesColumns holds the columns for the "houses" table.
	HousesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "title", Type: field.TypeString},
		{Name: "community", Type: field.TypeString},
		{Name: "toilet_count", Type: field.TypeInt32, Default: 0},
		{Name: "kitchen_count", Type: field.TypeInt32, Default: 0},
		{Name: "floor_count", Type: field.TypeInt32, Default: 0},
		{Name: "hall_count", Type: field.TypeInt32, Default: 0},
		{Name: "room_count", Type: field.TypeInt32, Default: 0},
	}
	// HousesTable holds the schema information for the "houses" table.
	HousesTable = &schema.Table{
		Name:       "houses",
		Columns:    HousesColumns,
		PrimaryKey: []*schema.Column{HousesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		HousesTable,
	}
)

func init() {
}
