// Code generated by ent, DO NOT EDIT.

package accesspolicy

import (
	"sectran_admin/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldName, v))
}

// Power applies equality check predicate on the "power" field. It's identical to PowerEQ.
func Power(v int32) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldPower, v))
}

// DepartmentID applies equality check predicate on the "department_id" field. It's identical to DepartmentIDEQ.
func DepartmentID(v uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldDepartmentID, v))
}

// Users applies equality check predicate on the "users" field. It's identical to UsersEQ.
func Users(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldUsers, v))
}

// Accounts applies equality check predicate on the "accounts" field. It's identical to AccountsEQ.
func Accounts(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldAccounts, v))
}

// EffecteTimeStart applies equality check predicate on the "effecte_time_start" field. It's identical to EffecteTimeStartEQ.
func EffecteTimeStart(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldEffecteTimeStart, v))
}

// EffecteTimeEnd applies equality check predicate on the "effecte_time_end" field. It's identical to EffecteTimeEndEQ.
func EffecteTimeEnd(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldEffecteTimeEnd, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldContainsFold(FieldName, v))
}

// PowerEQ applies the EQ predicate on the "power" field.
func PowerEQ(v int32) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldPower, v))
}

// PowerNEQ applies the NEQ predicate on the "power" field.
func PowerNEQ(v int32) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNEQ(FieldPower, v))
}

// PowerIn applies the In predicate on the "power" field.
func PowerIn(vs ...int32) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIn(FieldPower, vs...))
}

// PowerNotIn applies the NotIn predicate on the "power" field.
func PowerNotIn(vs ...int32) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotIn(FieldPower, vs...))
}

// PowerGT applies the GT predicate on the "power" field.
func PowerGT(v int32) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGT(FieldPower, v))
}

// PowerGTE applies the GTE predicate on the "power" field.
func PowerGTE(v int32) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGTE(FieldPower, v))
}

// PowerLT applies the LT predicate on the "power" field.
func PowerLT(v int32) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLT(FieldPower, v))
}

// PowerLTE applies the LTE predicate on the "power" field.
func PowerLTE(v int32) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLTE(FieldPower, v))
}

// PowerIsNil applies the IsNil predicate on the "power" field.
func PowerIsNil() predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIsNull(FieldPower))
}

// PowerNotNil applies the NotNil predicate on the "power" field.
func PowerNotNil() predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotNull(FieldPower))
}

// DepartmentIDEQ applies the EQ predicate on the "department_id" field.
func DepartmentIDEQ(v uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldDepartmentID, v))
}

// DepartmentIDNEQ applies the NEQ predicate on the "department_id" field.
func DepartmentIDNEQ(v uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNEQ(FieldDepartmentID, v))
}

// DepartmentIDIn applies the In predicate on the "department_id" field.
func DepartmentIDIn(vs ...uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIn(FieldDepartmentID, vs...))
}

// DepartmentIDNotIn applies the NotIn predicate on the "department_id" field.
func DepartmentIDNotIn(vs ...uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotIn(FieldDepartmentID, vs...))
}

// DepartmentIDGT applies the GT predicate on the "department_id" field.
func DepartmentIDGT(v uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGT(FieldDepartmentID, v))
}

// DepartmentIDGTE applies the GTE predicate on the "department_id" field.
func DepartmentIDGTE(v uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGTE(FieldDepartmentID, v))
}

// DepartmentIDLT applies the LT predicate on the "department_id" field.
func DepartmentIDLT(v uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLT(FieldDepartmentID, v))
}

// DepartmentIDLTE applies the LTE predicate on the "department_id" field.
func DepartmentIDLTE(v uint64) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLTE(FieldDepartmentID, v))
}

// UsersEQ applies the EQ predicate on the "users" field.
func UsersEQ(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldUsers, v))
}

// UsersNEQ applies the NEQ predicate on the "users" field.
func UsersNEQ(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNEQ(FieldUsers, v))
}

// UsersIn applies the In predicate on the "users" field.
func UsersIn(vs ...string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIn(FieldUsers, vs...))
}

// UsersNotIn applies the NotIn predicate on the "users" field.
func UsersNotIn(vs ...string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotIn(FieldUsers, vs...))
}

// UsersGT applies the GT predicate on the "users" field.
func UsersGT(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGT(FieldUsers, v))
}

// UsersGTE applies the GTE predicate on the "users" field.
func UsersGTE(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGTE(FieldUsers, v))
}

// UsersLT applies the LT predicate on the "users" field.
func UsersLT(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLT(FieldUsers, v))
}

// UsersLTE applies the LTE predicate on the "users" field.
func UsersLTE(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLTE(FieldUsers, v))
}

// UsersContains applies the Contains predicate on the "users" field.
func UsersContains(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldContains(FieldUsers, v))
}

// UsersHasPrefix applies the HasPrefix predicate on the "users" field.
func UsersHasPrefix(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldHasPrefix(FieldUsers, v))
}

// UsersHasSuffix applies the HasSuffix predicate on the "users" field.
func UsersHasSuffix(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldHasSuffix(FieldUsers, v))
}

// UsersEqualFold applies the EqualFold predicate on the "users" field.
func UsersEqualFold(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEqualFold(FieldUsers, v))
}

// UsersContainsFold applies the ContainsFold predicate on the "users" field.
func UsersContainsFold(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldContainsFold(FieldUsers, v))
}

// AccountsEQ applies the EQ predicate on the "accounts" field.
func AccountsEQ(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldAccounts, v))
}

// AccountsNEQ applies the NEQ predicate on the "accounts" field.
func AccountsNEQ(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNEQ(FieldAccounts, v))
}

// AccountsIn applies the In predicate on the "accounts" field.
func AccountsIn(vs ...string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIn(FieldAccounts, vs...))
}

// AccountsNotIn applies the NotIn predicate on the "accounts" field.
func AccountsNotIn(vs ...string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotIn(FieldAccounts, vs...))
}

// AccountsGT applies the GT predicate on the "accounts" field.
func AccountsGT(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGT(FieldAccounts, v))
}

// AccountsGTE applies the GTE predicate on the "accounts" field.
func AccountsGTE(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGTE(FieldAccounts, v))
}

// AccountsLT applies the LT predicate on the "accounts" field.
func AccountsLT(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLT(FieldAccounts, v))
}

// AccountsLTE applies the LTE predicate on the "accounts" field.
func AccountsLTE(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLTE(FieldAccounts, v))
}

// AccountsContains applies the Contains predicate on the "accounts" field.
func AccountsContains(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldContains(FieldAccounts, v))
}

// AccountsHasPrefix applies the HasPrefix predicate on the "accounts" field.
func AccountsHasPrefix(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldHasPrefix(FieldAccounts, v))
}

// AccountsHasSuffix applies the HasSuffix predicate on the "accounts" field.
func AccountsHasSuffix(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldHasSuffix(FieldAccounts, v))
}

// AccountsEqualFold applies the EqualFold predicate on the "accounts" field.
func AccountsEqualFold(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEqualFold(FieldAccounts, v))
}

// AccountsContainsFold applies the ContainsFold predicate on the "accounts" field.
func AccountsContainsFold(v string) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldContainsFold(FieldAccounts, v))
}

// EffecteTimeStartEQ applies the EQ predicate on the "effecte_time_start" field.
func EffecteTimeStartEQ(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldEffecteTimeStart, v))
}

// EffecteTimeStartNEQ applies the NEQ predicate on the "effecte_time_start" field.
func EffecteTimeStartNEQ(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNEQ(FieldEffecteTimeStart, v))
}

// EffecteTimeStartIn applies the In predicate on the "effecte_time_start" field.
func EffecteTimeStartIn(vs ...time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIn(FieldEffecteTimeStart, vs...))
}

// EffecteTimeStartNotIn applies the NotIn predicate on the "effecte_time_start" field.
func EffecteTimeStartNotIn(vs ...time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotIn(FieldEffecteTimeStart, vs...))
}

// EffecteTimeStartGT applies the GT predicate on the "effecte_time_start" field.
func EffecteTimeStartGT(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGT(FieldEffecteTimeStart, v))
}

// EffecteTimeStartGTE applies the GTE predicate on the "effecte_time_start" field.
func EffecteTimeStartGTE(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGTE(FieldEffecteTimeStart, v))
}

// EffecteTimeStartLT applies the LT predicate on the "effecte_time_start" field.
func EffecteTimeStartLT(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLT(FieldEffecteTimeStart, v))
}

// EffecteTimeStartLTE applies the LTE predicate on the "effecte_time_start" field.
func EffecteTimeStartLTE(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLTE(FieldEffecteTimeStart, v))
}

// EffecteTimeStartIsNil applies the IsNil predicate on the "effecte_time_start" field.
func EffecteTimeStartIsNil() predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIsNull(FieldEffecteTimeStart))
}

// EffecteTimeStartNotNil applies the NotNil predicate on the "effecte_time_start" field.
func EffecteTimeStartNotNil() predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotNull(FieldEffecteTimeStart))
}

// EffecteTimeEndEQ applies the EQ predicate on the "effecte_time_end" field.
func EffecteTimeEndEQ(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldEQ(FieldEffecteTimeEnd, v))
}

// EffecteTimeEndNEQ applies the NEQ predicate on the "effecte_time_end" field.
func EffecteTimeEndNEQ(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNEQ(FieldEffecteTimeEnd, v))
}

// EffecteTimeEndIn applies the In predicate on the "effecte_time_end" field.
func EffecteTimeEndIn(vs ...time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIn(FieldEffecteTimeEnd, vs...))
}

// EffecteTimeEndNotIn applies the NotIn predicate on the "effecte_time_end" field.
func EffecteTimeEndNotIn(vs ...time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotIn(FieldEffecteTimeEnd, vs...))
}

// EffecteTimeEndGT applies the GT predicate on the "effecte_time_end" field.
func EffecteTimeEndGT(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGT(FieldEffecteTimeEnd, v))
}

// EffecteTimeEndGTE applies the GTE predicate on the "effecte_time_end" field.
func EffecteTimeEndGTE(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldGTE(FieldEffecteTimeEnd, v))
}

// EffecteTimeEndLT applies the LT predicate on the "effecte_time_end" field.
func EffecteTimeEndLT(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLT(FieldEffecteTimeEnd, v))
}

// EffecteTimeEndLTE applies the LTE predicate on the "effecte_time_end" field.
func EffecteTimeEndLTE(v time.Time) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldLTE(FieldEffecteTimeEnd, v))
}

// EffecteTimeEndIsNil applies the IsNil predicate on the "effecte_time_end" field.
func EffecteTimeEndIsNil() predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldIsNull(FieldEffecteTimeEnd))
}

// EffecteTimeEndNotNil applies the NotNil predicate on the "effecte_time_end" field.
func EffecteTimeEndNotNil() predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.FieldNotNull(FieldEffecteTimeEnd))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccessPolicy) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccessPolicy) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccessPolicy) predicate.AccessPolicy {
	return predicate.AccessPolicy(sql.NotPredicates(p))
}
