// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sectran_admin/ent/account"
	"sectran_admin/ent/device"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Account is the model entity for the Account schema.
type Account struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// account username|账号名称
	Username string `json:"username,omitempty"`
	// account port|端口
	Port uint32 `json:"port,omitempty"`
	// protocol of the this account.|账号协议
	Protocol uint8 `json:"protocol,omitempty"`
	// account password|账号密码
	Password string `json:"password,omitempty"`
	// private_key of the this account.|账号私钥
	PrivateKey string `json:"private_key,omitempty"`
	// private_key password of the this account.|私钥口令
	PrivateKeyPassword string `json:"private_key_password,omitempty"`
	// account belong to|账号所属设备
	DeviceID uint64 `json:"device_id,omitempty"`
	// account belong to|账号所属部门
	DepartmentID uint64 `json:"department_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AccountQuery when eager-loading is set.
	Edges        AccountEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AccountEdges holds the relations/edges for other nodes in the graph.
type AccountEdges struct {
	// Devices holds the value of the devices edge.
	Devices *Device `json:"devices,omitempty"`
	// Departments holds the value of the departments edge.
	Departments *Device `json:"departments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DevicesOrErr returns the Devices value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) DevicesOrErr() (*Device, error) {
	if e.Devices != nil {
		return e.Devices, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: device.Label}
	}
	return nil, &NotLoadedError{edge: "devices"}
}

// DepartmentsOrErr returns the Departments value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AccountEdges) DepartmentsOrErr() (*Device, error) {
	if e.Departments != nil {
		return e.Departments, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: device.Label}
	}
	return nil, &NotLoadedError{edge: "departments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Account) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case account.FieldID, account.FieldPort, account.FieldProtocol, account.FieldDeviceID, account.FieldDepartmentID:
			values[i] = new(sql.NullInt64)
		case account.FieldUsername, account.FieldPassword, account.FieldPrivateKey, account.FieldPrivateKeyPassword:
			values[i] = new(sql.NullString)
		case account.FieldCreatedAt, account.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Account fields.
func (a *Account) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case account.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = uint64(value.Int64)
		case account.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case account.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case account.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				a.Username = value.String
			}
		case account.FieldPort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field port", values[i])
			} else if value.Valid {
				a.Port = uint32(value.Int64)
			}
		case account.FieldProtocol:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field protocol", values[i])
			} else if value.Valid {
				a.Protocol = uint8(value.Int64)
			}
		case account.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				a.Password = value.String
			}
		case account.FieldPrivateKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_key", values[i])
			} else if value.Valid {
				a.PrivateKey = value.String
			}
		case account.FieldPrivateKeyPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field private_key_password", values[i])
			} else if value.Valid {
				a.PrivateKeyPassword = value.String
			}
		case account.FieldDeviceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field device_id", values[i])
			} else if value.Valid {
				a.DeviceID = uint64(value.Int64)
			}
		case account.FieldDepartmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field department_id", values[i])
			} else if value.Valid {
				a.DepartmentID = uint64(value.Int64)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Account.
// This includes values selected through modifiers, order, etc.
func (a *Account) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryDevices queries the "devices" edge of the Account entity.
func (a *Account) QueryDevices() *DeviceQuery {
	return NewAccountClient(a.config).QueryDevices(a)
}

// QueryDepartments queries the "departments" edge of the Account entity.
func (a *Account) QueryDepartments() *DeviceQuery {
	return NewAccountClient(a.config).QueryDepartments(a)
}

// Update returns a builder for updating this Account.
// Note that you need to call Account.Unwrap() before calling this method if this Account
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Account) Update() *AccountUpdateOne {
	return NewAccountClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Account entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Account) Unwrap() *Account {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Account is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Account) String() string {
	var builder strings.Builder
	builder.WriteString("Account(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(a.Username)
	builder.WriteString(", ")
	builder.WriteString("port=")
	builder.WriteString(fmt.Sprintf("%v", a.Port))
	builder.WriteString(", ")
	builder.WriteString("protocol=")
	builder.WriteString(fmt.Sprintf("%v", a.Protocol))
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(a.Password)
	builder.WriteString(", ")
	builder.WriteString("private_key=")
	builder.WriteString(a.PrivateKey)
	builder.WriteString(", ")
	builder.WriteString("private_key_password=")
	builder.WriteString(a.PrivateKeyPassword)
	builder.WriteString(", ")
	builder.WriteString("device_id=")
	builder.WriteString(fmt.Sprintf("%v", a.DeviceID))
	builder.WriteString(", ")
	builder.WriteString("department_id=")
	builder.WriteString(fmt.Sprintf("%v", a.DepartmentID))
	builder.WriteByte(')')
	return builder.String()
}

// Accounts is a parsable slice of Account.
type Accounts []*Account
