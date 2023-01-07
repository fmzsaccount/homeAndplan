// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"gy_home/internal/data/ent/productcategory"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// ProductCategory is the model entity for the ProductCategory schema.
type ProductCategory struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Level holds the value of the "level" field.
	Level int32 `json:"level,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID int32 `json:"parent_id,omitempty"`
	// Sort holds the value of the "sort" field.
	Sort int32 `json:"sort,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductCategoryQuery when eager-loading is set.
	Edges ProductCategoryEdges `json:"edges"`
}

// ProductCategoryEdges holds the relations/edges for other nodes in the graph.
type ProductCategoryEdges struct {
	// Parent holds the value of the parent edge.
	Parent *ProductCategory `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*ProductCategory `json:"children,omitempty"`
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductCategoryEdges) ParentOrErr() (*ProductCategory, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: productcategory.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e ProductCategoryEdges) ChildrenOrErr() ([]*ProductCategory, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e ProductCategoryEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[2] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductCategory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case productcategory.FieldID, productcategory.FieldStatus, productcategory.FieldLevel, productcategory.FieldParentID, productcategory.FieldSort:
			values[i] = new(sql.NullInt64)
		case productcategory.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProductCategory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductCategory fields.
func (pc *ProductCategory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productcategory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = int32(value.Int64)
		case productcategory.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pc.Status = int(value.Int64)
			}
		case productcategory.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				pc.Level = int32(value.Int64)
			}
		case productcategory.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				pc.ParentID = int32(value.Int64)
			}
		case productcategory.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				pc.Sort = int32(value.Int64)
			}
		case productcategory.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pc.Name = value.String
			}
		}
	}
	return nil
}

// QueryParent queries the "parent" edge of the ProductCategory entity.
func (pc *ProductCategory) QueryParent() *ProductCategoryQuery {
	return (&ProductCategoryClient{config: pc.config}).QueryParent(pc)
}

// QueryChildren queries the "children" edge of the ProductCategory entity.
func (pc *ProductCategory) QueryChildren() *ProductCategoryQuery {
	return (&ProductCategoryClient{config: pc.config}).QueryChildren(pc)
}

// QueryProducts queries the "products" edge of the ProductCategory entity.
func (pc *ProductCategory) QueryProducts() *ProductQuery {
	return (&ProductCategoryClient{config: pc.config}).QueryProducts(pc)
}

// Update returns a builder for updating this ProductCategory.
// Note that you need to call ProductCategory.Unwrap() before calling this method if this ProductCategory
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *ProductCategory) Update() *ProductCategoryUpdateOne {
	return (&ProductCategoryClient{config: pc.config}).UpdateOne(pc)
}

// Unwrap unwraps the ProductCategory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *ProductCategory) Unwrap() *ProductCategory {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductCategory is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *ProductCategory) String() string {
	var builder strings.Builder
	builder.WriteString("ProductCategory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pc.Status))
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", pc.Level))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", pc.ParentID))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", pc.Sort))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pc.Name)
	builder.WriteByte(')')
	return builder.String()
}

// ProductCategories is a parsable slice of ProductCategory.
type ProductCategories []*ProductCategory

func (pc ProductCategories) config(cfg config) {
	for _i := range pc {
		pc[_i].config = cfg
	}
}
