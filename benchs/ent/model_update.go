// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/efectn/go-orm-benchmarks/benchs/ent/model"
	"github.com/efectn/go-orm-benchmarks/benchs/ent/predicate"
)

// ModelUpdate is the builder for updating Model entities.
type ModelUpdate struct {
	config
	hooks    []Hook
	mutation *ModelMutation
}

// Where appends a list predicates to the ModelUpdate builder.
func (mu *ModelUpdate) Where(ps ...predicate.Model) *ModelUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetName sets the "name" field.
func (mu *ModelUpdate) SetName(s string) *ModelUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetTitle sets the "title" field.
func (mu *ModelUpdate) SetTitle(s string) *ModelUpdate {
	mu.mutation.SetTitle(s)
	return mu
}

// SetFax sets the "fax" field.
func (mu *ModelUpdate) SetFax(s string) *ModelUpdate {
	mu.mutation.SetFax(s)
	return mu
}

// SetWeb sets the "web" field.
func (mu *ModelUpdate) SetWeb(s string) *ModelUpdate {
	mu.mutation.SetWeb(s)
	return mu
}

// SetAge sets the "age" field.
func (mu *ModelUpdate) SetAge(i int) *ModelUpdate {
	mu.mutation.ResetAge()
	mu.mutation.SetAge(i)
	return mu
}

// AddAge adds i to the "age" field.
func (mu *ModelUpdate) AddAge(i int) *ModelUpdate {
	mu.mutation.AddAge(i)
	return mu
}

// SetRight sets the "right" field.
func (mu *ModelUpdate) SetRight(b bool) *ModelUpdate {
	mu.mutation.SetRight(b)
	return mu
}

// SetCounter sets the "counter" field.
func (mu *ModelUpdate) SetCounter(i int64) *ModelUpdate {
	mu.mutation.ResetCounter()
	mu.mutation.SetCounter(i)
	return mu
}

// AddCounter adds i to the "counter" field.
func (mu *ModelUpdate) AddCounter(i int64) *ModelUpdate {
	mu.mutation.AddCounter(i)
	return mu
}

// Mutation returns the ModelMutation object of the builder.
func (mu *ModelUpdate) Mutation() *ModelMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *ModelUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *ModelUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *ModelUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *ModelUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *ModelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   model.Table,
			Columns: model.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: model.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: model.FieldName,
		})
	}
	if value, ok := mu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: model.FieldTitle,
		})
	}
	if value, ok := mu.mutation.Fax(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: model.FieldFax,
		})
	}
	if value, ok := mu.mutation.Web(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: model.FieldWeb,
		})
	}
	if value, ok := mu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: model.FieldAge,
		})
	}
	if value, ok := mu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: model.FieldAge,
		})
	}
	if value, ok := mu.mutation.Right(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: model.FieldRight,
		})
	}
	if value, ok := mu.mutation.Counter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: model.FieldCounter,
		})
	}
	if value, ok := mu.mutation.AddedCounter(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: model.FieldCounter,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{model.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ModelUpdateOne is the builder for updating a single Model entity.
type ModelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ModelMutation
}

// SetName sets the "name" field.
func (muo *ModelUpdateOne) SetName(s string) *ModelUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetTitle sets the "title" field.
func (muo *ModelUpdateOne) SetTitle(s string) *ModelUpdateOne {
	muo.mutation.SetTitle(s)
	return muo
}

// SetFax sets the "fax" field.
func (muo *ModelUpdateOne) SetFax(s string) *ModelUpdateOne {
	muo.mutation.SetFax(s)
	return muo
}

// SetWeb sets the "web" field.
func (muo *ModelUpdateOne) SetWeb(s string) *ModelUpdateOne {
	muo.mutation.SetWeb(s)
	return muo
}

// SetAge sets the "age" field.
func (muo *ModelUpdateOne) SetAge(i int) *ModelUpdateOne {
	muo.mutation.ResetAge()
	muo.mutation.SetAge(i)
	return muo
}

// AddAge adds i to the "age" field.
func (muo *ModelUpdateOne) AddAge(i int) *ModelUpdateOne {
	muo.mutation.AddAge(i)
	return muo
}

// SetRight sets the "right" field.
func (muo *ModelUpdateOne) SetRight(b bool) *ModelUpdateOne {
	muo.mutation.SetRight(b)
	return muo
}

// SetCounter sets the "counter" field.
func (muo *ModelUpdateOne) SetCounter(i int64) *ModelUpdateOne {
	muo.mutation.ResetCounter()
	muo.mutation.SetCounter(i)
	return muo
}

// AddCounter adds i to the "counter" field.
func (muo *ModelUpdateOne) AddCounter(i int64) *ModelUpdateOne {
	muo.mutation.AddCounter(i)
	return muo
}

// Mutation returns the ModelMutation object of the builder.
func (muo *ModelUpdateOne) Mutation() *ModelMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *ModelUpdateOne) Select(field string, fields ...string) *ModelUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Model entity.
func (muo *ModelUpdateOne) Save(ctx context.Context) (*Model, error) {
	var (
		err  error
		node *Model
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *ModelUpdateOne) SaveX(ctx context.Context) *Model {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *ModelUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *ModelUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *ModelUpdateOne) sqlSave(ctx context.Context) (_node *Model, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   model.Table,
			Columns: model.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: model.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Model.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, model.FieldID)
		for _, f := range fields {
			if !model.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != model.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: model.FieldName,
		})
	}
	if value, ok := muo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: model.FieldTitle,
		})
	}
	if value, ok := muo.mutation.Fax(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: model.FieldFax,
		})
	}
	if value, ok := muo.mutation.Web(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: model.FieldWeb,
		})
	}
	if value, ok := muo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: model.FieldAge,
		})
	}
	if value, ok := muo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: model.FieldAge,
		})
	}
	if value, ok := muo.mutation.Right(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: model.FieldRight,
		})
	}
	if value, ok := muo.mutation.Counter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: model.FieldCounter,
		})
	}
	if value, ok := muo.mutation.AddedCounter(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: model.FieldCounter,
		})
	}
	_node = &Model{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{model.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
