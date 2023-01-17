// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yekta/stablecog/go-apps/database/ent/predicate"
	"github.com/yekta/stablecog/go-apps/database/ent/server"
)

// ServerDelete is the builder for deleting a Server entity.
type ServerDelete struct {
	config
	hooks    []Hook
	mutation *ServerMutation
}

// Where appends a list predicates to the ServerDelete builder.
func (sd *ServerDelete) Where(ps ...predicate.Server) *ServerDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *ServerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ServerMutation](ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *ServerDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *ServerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: server.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: server.FieldID,
			},
		},
	}
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// ServerDeleteOne is the builder for deleting a single Server entity.
type ServerDeleteOne struct {
	sd *ServerDelete
}

// Exec executes the deletion query.
func (sdo *ServerDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{server.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *ServerDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}
