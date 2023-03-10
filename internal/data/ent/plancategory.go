// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"gy_home/internal/data/ent/plancategory"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// PlanCategory is the model entity for the PlanCategory schema.
type PlanCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Sort holds the value of the "sort" field.
	Sort int32 `json:"sort,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PlanCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case plancategory.FieldID, plancategory.FieldSort:
			values[i] = new(sql.NullInt64)
		case plancategory.FieldName, plancategory.FieldContent, plancategory.FieldURL:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PlanCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PlanCategory fields.
func (pc *PlanCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case plancategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = int32(value.Int64)
		case plancategory.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				pc.Sort = int32(value.Int64)
			}
		case plancategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pc.Name = value.String
			}
		case plancategory.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				pc.Content = value.String
			}
		case plancategory.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				pc.URL = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PlanCategory.
// Note that you need to call PlanCategory.Unwrap() before calling this method if this PlanCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *PlanCategory) Update() *PlanCategoryUpdateOne {
	return (&PlanCategoryClient{config: pc.config}).UpdateOne(pc)
}

// Unwrap unwraps the PlanCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *PlanCategory) Unwrap() *PlanCategory {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PlanCategory is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *PlanCategory) String() string {
	var builder strings.Builder
	builder.WriteString("PlanCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", pc.Sort))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pc.Name)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(pc.Content)
	builder.WriteString(", ")
	builder.WriteString("url=")
	builder.WriteString(pc.URL)
	builder.WriteByte(')')
	return builder.String()
}

// PlanCategories is a parsable slice of PlanCategory.
type PlanCategories []*PlanCategory

func (pc PlanCategories) config(cfg config) {
	for _i := range pc {
		pc[_i].config = cfg
	}
}
