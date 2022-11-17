// Code generated by ent, DO NOT EDIT.

package services

const (
	// Label holds the string label denoting the services type in the database.
	Label = "services"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// EdgeNamespace holds the string denoting the namespace edge name in mutations.
	EdgeNamespace = "namespace"
	// NamespaceFieldID holds the string denoting the ID field of the Namespace.
	NamespaceFieldID = "oid"
	// Table holds the table name of the services in the database.
	Table = "services"
	// NamespaceTable is the table that holds the namespace relation/edge.
	NamespaceTable = "services"
	// NamespaceInverseTable is the table name for the Namespace entity.
	// It exists in this package in order to avoid circular dependency with the "namespace" package.
	NamespaceInverseTable = "namespaces"
	// NamespaceColumn is the table column denoting the namespace relation/edge.
	NamespaceColumn = "namespace_services"
)

// Columns holds all SQL columns for services fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldData,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "services"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"namespace_services",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)
