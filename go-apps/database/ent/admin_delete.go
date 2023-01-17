// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yekta/stablecog/go-apps/database/ent/admin"
	"github.com/yekta/stablecog/go-apps/database/ent/predicate"
)

// AdminDelete is the builder for deleting a Admin entity.
type AdminDelete struct {
	config
	hooks    []Hook
	mutation *AdminMutation
}

// Where appends a list predicates to the AdminDelete builder.
func (ad *AdminDelete) Where(ps ...predicate.Admin) *AdminDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AdminDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, AdminMutation](ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AdminDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AdminDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: admin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: admin.FieldID,
			},
		},
	}
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AdminDeleteOne is the builder for deleting a single Admin entity.
type AdminDeleteOne struct {
	ad *AdminDelete
}

// Exec executes the deletion query.
func (ado *AdminDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{admin.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AdminDeleteOne) ExecX(ctx context.Context) {
	ado.ad.ExecX(ctx)
}
