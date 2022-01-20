// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"github.com/starryrbs/kfan/app/history/service/internal/data/ent/history"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// History is the model entity for the History schema.
type History struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// ObjID holds the value of the "obj_id" field.
	ObjID int64 `json:"obj_id,omitempty"`
	// ObjType holds the value of the "obj_type" field.
	ObjType string `json:"obj_type,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*History) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case history.FieldID, history.FieldUserID, history.FieldObjID:
			values[i] = new(sql.NullInt64)
		case history.FieldObjType:
			values[i] = new(sql.NullString)
		case history.FieldCreatedAt, history.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type History", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the History fields.
func (h *History) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case history.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int64(value.Int64)
		case history.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				h.UserID = value.Int64
			}
		case history.FieldObjID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field obj_id", values[i])
			} else if value.Valid {
				h.ObjID = value.Int64
			}
		case history.FieldObjType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field obj_type", values[i])
			} else if value.Valid {
				h.ObjType = value.String
			}
		case history.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				h.CreatedAt = value.Time
			}
		case history.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				h.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this History.
// Note that you need to call History.Unwrap() before calling this method if this History
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *History) Update() *HistoryUpdateOne {
	return (&HistoryClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the History entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *History) Unwrap() *History {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: History is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *History) String() string {
	var builder strings.Builder
	builder.WriteString("History(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", h.UserID))
	builder.WriteString(", obj_id=")
	builder.WriteString(fmt.Sprintf("%v", h.ObjID))
	builder.WriteString(", obj_type=")
	builder.WriteString(h.ObjType)
	builder.WriteString(", created_at=")
	builder.WriteString(h.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(h.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Histories is a parsable slice of History.
type Histories []*History

func (h Histories) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
