// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/yekta/stablecog/go-apps/database/ent/generation"
	"github.com/yekta/stablecog/go-apps/database/ent/generationg"
	"github.com/yekta/stablecog/go-apps/database/ent/predicate"
	"github.com/yekta/stablecog/go-apps/database/ent/prompt"
)

// PromptQuery is the builder for querying Prompt entities.
type PromptQuery struct {
	config
	limit           *int
	offset          *int
	unique          *bool
	order           []OrderFunc
	fields          []string
	inters          []Interceptor
	predicates      []predicate.Prompt
	withGeneration  *GenerationQuery
	withGenerationG *GenerationGQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PromptQuery builder.
func (pq *PromptQuery) Where(ps ...predicate.Prompt) *PromptQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PromptQuery) Limit(limit int) *PromptQuery {
	pq.limit = &limit
	return pq
}

// Offset to start from.
func (pq *PromptQuery) Offset(offset int) *PromptQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PromptQuery) Unique(unique bool) *PromptQuery {
	pq.unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PromptQuery) Order(o ...OrderFunc) *PromptQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryGeneration chains the current query on the "generation" edge.
func (pq *PromptQuery) QueryGeneration() *GenerationQuery {
	query := (&GenerationClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(prompt.Table, prompt.FieldID, selector),
			sqlgraph.To(generation.Table, generation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, prompt.GenerationTable, prompt.GenerationColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGenerationG chains the current query on the "generation_g" edge.
func (pq *PromptQuery) QueryGenerationG() *GenerationGQuery {
	query := (&GenerationGClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(prompt.Table, prompt.FieldID, selector),
			sqlgraph.To(generationg.Table, generationg.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, prompt.GenerationGTable, prompt.GenerationGColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Prompt entity from the query.
// Returns a *NotFoundError when no Prompt was found.
func (pq *PromptQuery) First(ctx context.Context) (*Prompt, error) {
	nodes, err := pq.Limit(1).All(newQueryContext(ctx, TypePrompt, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{prompt.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PromptQuery) FirstX(ctx context.Context) *Prompt {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Prompt ID from the query.
// Returns a *NotFoundError when no Prompt ID was found.
func (pq *PromptQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(newQueryContext(ctx, TypePrompt, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{prompt.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PromptQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Prompt entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Prompt entity is found.
// Returns a *NotFoundError when no Prompt entities are found.
func (pq *PromptQuery) Only(ctx context.Context) (*Prompt, error) {
	nodes, err := pq.Limit(2).All(newQueryContext(ctx, TypePrompt, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{prompt.Label}
	default:
		return nil, &NotSingularError{prompt.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PromptQuery) OnlyX(ctx context.Context) *Prompt {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Prompt ID in the query.
// Returns a *NotSingularError when more than one Prompt ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PromptQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(newQueryContext(ctx, TypePrompt, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{prompt.Label}
	default:
		err = &NotSingularError{prompt.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PromptQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Prompts.
func (pq *PromptQuery) All(ctx context.Context) ([]*Prompt, error) {
	ctx = newQueryContext(ctx, TypePrompt, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Prompt, *PromptQuery]()
	return withInterceptors[[]*Prompt](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PromptQuery) AllX(ctx context.Context) []*Prompt {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Prompt IDs.
func (pq *PromptQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = newQueryContext(ctx, TypePrompt, "IDs")
	if err := pq.Select(prompt.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PromptQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PromptQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypePrompt, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PromptQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PromptQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PromptQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypePrompt, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PromptQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PromptQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PromptQuery) Clone() *PromptQuery {
	if pq == nil {
		return nil
	}
	return &PromptQuery{
		config:          pq.config,
		limit:           pq.limit,
		offset:          pq.offset,
		order:           append([]OrderFunc{}, pq.order...),
		inters:          append([]Interceptor{}, pq.inters...),
		predicates:      append([]predicate.Prompt{}, pq.predicates...),
		withGeneration:  pq.withGeneration.Clone(),
		withGenerationG: pq.withGenerationG.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithGeneration tells the query-builder to eager-load the nodes that are connected to
// the "generation" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PromptQuery) WithGeneration(opts ...func(*GenerationQuery)) *PromptQuery {
	query := (&GenerationClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withGeneration = query
	return pq
}

// WithGenerationG tells the query-builder to eager-load the nodes that are connected to
// the "generation_g" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PromptQuery) WithGenerationG(opts ...func(*GenerationGQuery)) *PromptQuery {
	query := (&GenerationGClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withGenerationG = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Prompt.Query().
//		GroupBy(prompt.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PromptQuery) GroupBy(field string, fields ...string) *PromptGroupBy {
	pq.fields = append([]string{field}, fields...)
	grbuild := &PromptGroupBy{build: pq}
	grbuild.flds = &pq.fields
	grbuild.label = prompt.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Prompt.Query().
//		Select(prompt.FieldText).
//		Scan(ctx, &v)
func (pq *PromptQuery) Select(fields ...string) *PromptSelect {
	pq.fields = append(pq.fields, fields...)
	sbuild := &PromptSelect{PromptQuery: pq}
	sbuild.label = prompt.Label
	sbuild.flds, sbuild.scan = &pq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PromptSelect configured with the given aggregations.
func (pq *PromptQuery) Aggregate(fns ...AggregateFunc) *PromptSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PromptQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.fields {
		if !prompt.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PromptQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Prompt, error) {
	var (
		nodes       = []*Prompt{}
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withGeneration != nil,
			pq.withGenerationG != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Prompt).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Prompt{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withGeneration; query != nil {
		if err := pq.loadGeneration(ctx, query, nodes,
			func(n *Prompt) { n.Edges.Generation = []*Generation{} },
			func(n *Prompt, e *Generation) { n.Edges.Generation = append(n.Edges.Generation, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withGenerationG; query != nil {
		if err := pq.loadGenerationG(ctx, query, nodes,
			func(n *Prompt) { n.Edges.GenerationG = []*GenerationG{} },
			func(n *Prompt, e *GenerationG) { n.Edges.GenerationG = append(n.Edges.GenerationG, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PromptQuery) loadGeneration(ctx context.Context, query *GenerationQuery, nodes []*Prompt, init func(*Prompt), assign func(*Prompt, *Generation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Prompt)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Generation(func(s *sql.Selector) {
		s.Where(sql.InValues(prompt.GenerationColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PromptID
		if fk == nil {
			return fmt.Errorf(`foreign-key "prompt_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "prompt_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PromptQuery) loadGenerationG(ctx context.Context, query *GenerationGQuery, nodes []*Prompt, init func(*Prompt), assign func(*Prompt, *GenerationG)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Prompt)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.GenerationG(func(s *sql.Selector) {
		s.Where(sql.InValues(prompt.GenerationGColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PromptID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "prompt_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PromptQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PromptQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prompt.Table,
			Columns: prompt.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: prompt.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prompt.FieldID)
		for i := range fields {
			if fields[i] != prompt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PromptQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(prompt.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = prompt.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PromptGroupBy is the group-by builder for Prompt entities.
type PromptGroupBy struct {
	selector
	build *PromptQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PromptGroupBy) Aggregate(fns ...AggregateFunc) *PromptGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PromptGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypePrompt, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PromptQuery, *PromptGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PromptGroupBy) sqlScan(ctx context.Context, root *PromptQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PromptSelect is the builder for selecting fields of Prompt entities.
type PromptSelect struct {
	*PromptQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PromptSelect) Aggregate(fns ...AggregateFunc) *PromptSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PromptSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypePrompt, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PromptQuery, *PromptSelect](ctx, ps.PromptQuery, ps, ps.inters, v)
}

func (ps *PromptSelect) sqlScan(ctx context.Context, root *PromptQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
