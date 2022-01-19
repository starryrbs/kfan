// Code generated by entc, DO NOT EDIT.

package account

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// Table holds the table name of the account in the database.
	Table = "accounts"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAge,
	FieldEmail,
	FieldSex,
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultAge holds the default value on creation for the "age" field.
	DefaultAge int
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)
