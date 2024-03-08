// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"sectran_admin/ent/policyauth"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PolicyAuth is the model entity for the PolicyAuth schema.
type PolicyAuth struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// policy name|策略名称
	Name string `json:"name,omitempty"`
	// policy power|策略优先级
	Power int32 `json:"power,omitempty"`
	// ID of the policy's department.|策略所属部门
	DepartmentID uint64 `json:"department_id,omitempty"`
	// 策略关联用户
	Users string `json:"users,omitempty"`
	// 策略关联账号
	Accounts     string `json:"accounts,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PolicyAuth) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case policyauth.FieldID, policyauth.FieldPower, policyauth.FieldDepartmentID:
			values[i] = new(sql.NullInt64)
		case policyauth.FieldName, policyauth.FieldUsers, policyauth.FieldAccounts:
			values[i] = new(sql.NullString)
		case policyauth.FieldCreatedAt, policyauth.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PolicyAuth fields.
func (pa *PolicyAuth) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case policyauth.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pa.ID = uint64(value.Int64)
		case policyauth.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pa.CreatedAt = value.Time
			}
		case policyauth.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pa.UpdatedAt = value.Time
			}
		case policyauth.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pa.Name = value.String
			}
		case policyauth.FieldPower:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field power", values[i])
			} else if value.Valid {
				pa.Power = int32(value.Int64)
			}
		case policyauth.FieldDepartmentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field department_id", values[i])
			} else if value.Valid {
				pa.DepartmentID = uint64(value.Int64)
			}
		case policyauth.FieldUsers:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field users", values[i])
			} else if value.Valid {
				pa.Users = value.String
			}
		case policyauth.FieldAccounts:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field accounts", values[i])
			} else if value.Valid {
				pa.Accounts = value.String
			}
		default:
			pa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PolicyAuth.
// This includes values selected through modifiers, order, etc.
func (pa *PolicyAuth) Value(name string) (ent.Value, error) {
	return pa.selectValues.Get(name)
}

// Update returns a builder for updating this PolicyAuth.
// Note that you need to call PolicyAuth.Unwrap() before calling this method if this PolicyAuth
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *PolicyAuth) Update() *PolicyAuthUpdateOne {
	return NewPolicyAuthClient(pa.config).UpdateOne(pa)
}

// Unwrap unwraps the PolicyAuth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pa *PolicyAuth) Unwrap() *PolicyAuth {
	_tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: PolicyAuth is not a transactional entity")
	}
	pa.config.driver = _tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *PolicyAuth) String() string {
	var builder strings.Builder
	builder.WriteString("PolicyAuth(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pa.ID))
	builder.WriteString("created_at=")
	builder.WriteString(pa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(pa.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pa.Name)
	builder.WriteString(", ")
	builder.WriteString("power=")
	builder.WriteString(fmt.Sprintf("%v", pa.Power))
	builder.WriteString(", ")
	builder.WriteString("department_id=")
	builder.WriteString(fmt.Sprintf("%v", pa.DepartmentID))
	builder.WriteString(", ")
	builder.WriteString("users=")
	builder.WriteString(pa.Users)
	builder.WriteString(", ")
	builder.WriteString("accounts=")
	builder.WriteString(pa.Accounts)
	builder.WriteByte(')')
	return builder.String()
}

// PolicyAuths is a parsable slice of PolicyAuth.
type PolicyAuths []*PolicyAuth