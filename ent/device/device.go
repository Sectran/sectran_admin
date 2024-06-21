// Code generated by ent, DO NOT EDIT.

package device

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the device type in the database.
	Label = "device"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDepartmentID holds the string denoting the department_id field in the database.
	FieldDepartmentID = "department_id"
	// FieldHost holds the string denoting the host field in the database.
	FieldHost = "host"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLables holds the string denoting the lables field in the database.
	FieldLables = "lables"
	// EdgeDepartments holds the string denoting the departments edge name in mutations.
	EdgeDepartments = "departments"
	// EdgeAccounts holds the string denoting the accounts edge name in mutations.
	EdgeAccounts = "accounts"
	// Table holds the table name of the device in the database.
	Table = "devices"
	// DepartmentsTable is the table that holds the departments relation/edge.
	DepartmentsTable = "devices"
	// DepartmentsInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentsInverseTable = "departments"
	// DepartmentsColumn is the table column denoting the departments relation/edge.
	DepartmentsColumn = "department_id"
	// AccountsTable is the table that holds the accounts relation/edge.
	AccountsTable = "accounts"
	// AccountsInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	AccountsInverseTable = "accounts"
	// AccountsColumn is the table column denoting the accounts relation/edge.
	AccountsColumn = "device_id"
)

// Columns holds all SQL columns for device fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDepartmentID,
	FieldHost,
	FieldType,
	FieldDescription,
	FieldLables,
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
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	DepartmentIDValidator func(uint64) error
	// HostValidator is a validator for the "host" field. It is called by the builders before save.
	HostValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)

// OrderOption defines the ordering options for the Device queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDepartmentID orders the results by the department_id field.
func ByDepartmentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartmentID, opts...).ToFunc()
}

// ByHost orders the results by the host field.
func ByHost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHost, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByLables orders the results by the lables field.
func ByLables(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLables, opts...).ToFunc()
}

// ByDepartmentsField orders the results by departments field.
func ByDepartmentsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDepartmentsStep(), sql.OrderByField(field, opts...))
	}
}

// ByAccountsCount orders the results by accounts count.
func ByAccountsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAccountsStep(), opts...)
	}
}

// ByAccounts orders the results by accounts terms.
func ByAccounts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccountsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDepartmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DepartmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DepartmentsTable, DepartmentsColumn),
	)
}
func newAccountsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccountsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, AccountsTable, AccountsColumn),
	)
}
