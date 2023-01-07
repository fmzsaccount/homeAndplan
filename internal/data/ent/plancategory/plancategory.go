// Code generated by ent, DO NOT EDIT.

package plancategory

const (
	// Label holds the string label denoting the plancategory type in the database.
	Label = "plan_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSort holds the string denoting the sort field in the database.
	FieldSort = "sort"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// Table holds the table name of the plancategory in the database.
	Table = "plan_category"
)

// Columns holds all SQL columns for plancategory fields.
var Columns = []string{
	FieldID,
	FieldSort,
	FieldName,
	FieldContent,
	FieldURL,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSort holds the default value on creation for the "sort" field.
	DefaultSort int32
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultContent holds the default value on creation for the "content" field.
	DefaultContent string
	// DefaultURL holds the default value on creation for the "url" field.
	DefaultURL string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID int32
)