// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sectran_admin/ent/department"
	"sectran_admin/ent/device"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Device is the model entity for the Device schema.
type Device struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// The name of the device.|设备名称
	Name string `json:"name,omitempty"`
	// ID of the device's department.|设备所属部门
	DepartmentID uint64 `json:"department_id,omitempty"`
	// login host|设备地址
	Host string `json:"host,omitempty"`
	// type of the device.|设备类型
	Type string `json:"type,omitempty"`
	// Description of the device.|设备描述
	Description string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceQuery when eager-loading is set.
	Edges        DeviceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DeviceEdges holds the relations/edges for other nodes in the graph.
type DeviceEdges struct {
	// Departments holds the value of the departments edge.
	Departments *Department `json:"departments,omitempty"`
	// Accounts holds the value of the accounts edge.
	Accounts []*Account `json:"accounts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DepartmentsOrErr returns the Departments value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceEdges) DepartmentsOrErr() (*Department, error) {
	if e.Departments != nil {
		return e.Departments, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: department.Label}
	}
	return nil, &NotLoadedError{edge: "departments"}
}

// AccountsOrErr returns the Accounts value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) AccountsOrErr() ([]*Account, error) {
	if e.loadedTypes[1] {
		return e.Accounts, nil
	}
	return nil, &NotLoadedError{edge: "accounts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case device.FieldID, device.FieldDepartmentID:
			values[i] = new(sql.NullInt64)
		case device.FieldName, device.FieldHost, device.FieldType, device.FieldDescription:
			values[i] = new(sql.NullString)
		case device.FieldCreatedAt, device.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case device.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = uint64(value.Int64)
		case device.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Time
			}
		case device.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Time
			}
		case device.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case device.FieldDepartmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field department_id", values[i])
			} else if value.Valid {
				d.DepartmentID = uint64(value.Int64)
			}
		case device.FieldHost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host", values[i])
			} else if value.Valid {
				d.Host = value.String
			}
		case device.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				d.Type = value.String
			}
		case device.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				d.Description = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Device.
// This includes values selected through modifiers, order, etc.
func (d *Device) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryDepartments queries the "departments" edge of the Device entity.
func (d *Device) QueryDepartments() *DepartmentQuery {
	return NewDeviceClient(d.config).QueryDepartments(d)
}

// QueryAccounts queries the "accounts" edge of the Device entity.
func (d *Device) QueryAccounts() *AccountQuery {
	return NewDeviceClient(d.config).QueryAccounts(d)
}

// Update returns a builder for updating this Device.
// Note that you need to call Device.Unwrap() before calling this method if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return NewDeviceClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Device entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Device is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("created_at=")
	builder.WriteString(d.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(d.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	builder.WriteString("department_id=")
	builder.WriteString(fmt.Sprintf("%v", d.DepartmentID))
	builder.WriteString(", ")
	builder.WriteString("host=")
	builder.WriteString(d.Host)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(d.Type)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(d.Description)
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device
