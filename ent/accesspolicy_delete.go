// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"sectran_admin/ent/accesspolicy"
	"sectran_admin/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessPolicyDelete is the builder for deleting a AccessPolicy entity.
type AccessPolicyDelete struct {
	config
	hooks    []Hook
	mutation *AccessPolicyMutation
}

// Where appends a list predicates to the AccessPolicyDelete builder.
func (apd *AccessPolicyDelete) Where(ps ...predicate.AccessPolicy) *AccessPolicyDelete {
	apd.mutation.Where(ps...)
	return apd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (apd *AccessPolicyDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, apd.sqlExec, apd.mutation, apd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (apd *AccessPolicyDelete) ExecX(ctx context.Context) int {
	n, err := apd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (apd *AccessPolicyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(accesspolicy.Table, sqlgraph.NewFieldSpec(accesspolicy.FieldID, field.TypeUint64))
	if ps := apd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, apd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	apd.mutation.done = true
	return affected, err
}

// AccessPolicyDeleteOne is the builder for deleting a single AccessPolicy entity.
type AccessPolicyDeleteOne struct {
	apd *AccessPolicyDelete
}

// Where appends a list predicates to the AccessPolicyDelete builder.
func (apdo *AccessPolicyDeleteOne) Where(ps ...predicate.AccessPolicy) *AccessPolicyDeleteOne {
	apdo.apd.mutation.Where(ps...)
	return apdo
}

// Exec executes the deletion query.
func (apdo *AccessPolicyDeleteOne) Exec(ctx context.Context) error {
	n, err := apdo.apd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accesspolicy.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (apdo *AccessPolicyDeleteOne) ExecX(ctx context.Context) {
	if err := apdo.Exec(ctx); err != nil {
		panic(err)
	}
}
