// Code generated by entc, DO NOT EDIT.

package model

const (
	// Label holds the string label denoting the model type in the database.
	Label = "model"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldFax holds the string denoting the fax field in the database.
	FieldFax = "fax"
	// FieldWeb holds the string denoting the web field in the database.
	FieldWeb = "web"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldRight holds the string denoting the right field in the database.
	FieldRight = "right"
	// FieldCounter holds the string denoting the counter field in the database.
	FieldCounter = "counter"
	// Table holds the table name of the model in the database.
	Table = "models"
)

// Columns holds all SQL columns for model fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldTitle,
	FieldFax,
	FieldWeb,
	FieldAge,
	FieldRight,
	FieldCounter,
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
