// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gy_home/internal/data/ent/plan"
	"gy_home/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlanUpdate is the builder for updating Plan entities.
type PlanUpdate struct {
	config
	hooks    []Hook
	mutation *PlanMutation
}

// Where appends a list predicates to the PlanUpdate builder.
func (pu *PlanUpdate) Where(ps ...predicate.Plan) *PlanUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetRecommend sets the "recommend" field.
func (pu *PlanUpdate) SetRecommend(i int) *PlanUpdate {
	pu.mutation.ResetRecommend()
	pu.mutation.SetRecommend(i)
	return pu
}

// SetNillableRecommend sets the "recommend" field if the given value is not nil.
func (pu *PlanUpdate) SetNillableRecommend(i *int) *PlanUpdate {
	if i != nil {
		pu.SetRecommend(*i)
	}
	return pu
}

// AddRecommend adds i to the "recommend" field.
func (pu *PlanUpdate) AddRecommend(i int) *PlanUpdate {
	pu.mutation.AddRecommend(i)
	return pu
}

// SetCategoryID sets the "category_id" field.
func (pu *PlanUpdate) SetCategoryID(i int32) *PlanUpdate {
	pu.mutation.ResetCategoryID()
	pu.mutation.SetCategoryID(i)
	return pu
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (pu *PlanUpdate) SetNillableCategoryID(i *int32) *PlanUpdate {
	if i != nil {
		pu.SetCategoryID(*i)
	}
	return pu
}

// AddCategoryID adds i to the "category_id" field.
func (pu *PlanUpdate) AddCategoryID(i int32) *PlanUpdate {
	pu.mutation.AddCategoryID(i)
	return pu
}

// SetCategory sets the "category" field.
func (pu *PlanUpdate) SetCategory(s string) *PlanUpdate {
	pu.mutation.SetCategory(s)
	return pu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (pu *PlanUpdate) SetNillableCategory(s *string) *PlanUpdate {
	if s != nil {
		pu.SetCategory(*s)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *PlanUpdate) SetName(s string) *PlanUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PlanUpdate) SetNillableName(s *string) *PlanUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetURL sets the "url" field.
func (pu *PlanUpdate) SetURL(s string) *PlanUpdate {
	pu.mutation.SetURL(s)
	return pu
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (pu *PlanUpdate) SetNillableURL(s *string) *PlanUpdate {
	if s != nil {
		pu.SetURL(*s)
	}
	return pu
}

// SetUtime sets the "utime" field.
func (pu *PlanUpdate) SetUtime(t time.Time) *PlanUpdate {
	pu.mutation.SetUtime(t)
	return pu
}

// Mutation returns the PlanMutation object of the builder.
func (pu *PlanUpdate) Mutation() *PlanMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlanUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlanUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlanUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlanUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PlanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   plan.Table,
			Columns: plan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: plan.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Recommend(); ok {
		_spec.SetField(plan.FieldRecommend, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedRecommend(); ok {
		_spec.AddField(plan.FieldRecommend, field.TypeInt, value)
	}
	if value, ok := pu.mutation.CategoryID(); ok {
		_spec.SetField(plan.FieldCategoryID, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.AddedCategoryID(); ok {
		_spec.AddField(plan.FieldCategoryID, field.TypeInt32, value)
	}
	if value, ok := pu.mutation.Category(); ok {
		_spec.SetField(plan.FieldCategory, field.TypeString, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(plan.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.URL(); ok {
		_spec.SetField(plan.FieldURL, field.TypeString, value)
	}
	if value, ok := pu.mutation.Utime(); ok {
		_spec.SetField(plan.FieldUtime, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{plan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PlanUpdateOne is the builder for updating a single Plan entity.
type PlanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlanMutation
}

// SetRecommend sets the "recommend" field.
func (puo *PlanUpdateOne) SetRecommend(i int) *PlanUpdateOne {
	puo.mutation.ResetRecommend()
	puo.mutation.SetRecommend(i)
	return puo
}

// SetNillableRecommend sets the "recommend" field if the given value is not nil.
func (puo *PlanUpdateOne) SetNillableRecommend(i *int) *PlanUpdateOne {
	if i != nil {
		puo.SetRecommend(*i)
	}
	return puo
}

// AddRecommend adds i to the "recommend" field.
func (puo *PlanUpdateOne) AddRecommend(i int) *PlanUpdateOne {
	puo.mutation.AddRecommend(i)
	return puo
}

// SetCategoryID sets the "category_id" field.
func (puo *PlanUpdateOne) SetCategoryID(i int32) *PlanUpdateOne {
	puo.mutation.ResetCategoryID()
	puo.mutation.SetCategoryID(i)
	return puo
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (puo *PlanUpdateOne) SetNillableCategoryID(i *int32) *PlanUpdateOne {
	if i != nil {
		puo.SetCategoryID(*i)
	}
	return puo
}

// AddCategoryID adds i to the "category_id" field.
func (puo *PlanUpdateOne) AddCategoryID(i int32) *PlanUpdateOne {
	puo.mutation.AddCategoryID(i)
	return puo
}

// SetCategory sets the "category" field.
func (puo *PlanUpdateOne) SetCategory(s string) *PlanUpdateOne {
	puo.mutation.SetCategory(s)
	return puo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (puo *PlanUpdateOne) SetNillableCategory(s *string) *PlanUpdateOne {
	if s != nil {
		puo.SetCategory(*s)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *PlanUpdateOne) SetName(s string) *PlanUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PlanUpdateOne) SetNillableName(s *string) *PlanUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetURL sets the "url" field.
func (puo *PlanUpdateOne) SetURL(s string) *PlanUpdateOne {
	puo.mutation.SetURL(s)
	return puo
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (puo *PlanUpdateOne) SetNillableURL(s *string) *PlanUpdateOne {
	if s != nil {
		puo.SetURL(*s)
	}
	return puo
}

// SetUtime sets the "utime" field.
func (puo *PlanUpdateOne) SetUtime(t time.Time) *PlanUpdateOne {
	puo.mutation.SetUtime(t)
	return puo
}

// Mutation returns the PlanMutation object of the builder.
func (puo *PlanUpdateOne) Mutation() *PlanMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlanUpdateOne) Select(field string, fields ...string) *PlanUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Plan entity.
func (puo *PlanUpdateOne) Save(ctx context.Context) (*Plan, error) {
	var (
		err  error
		node *Plan
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Plan)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PlanMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlanUpdateOne) SaveX(ctx context.Context) *Plan {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlanUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlanUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PlanUpdateOne) sqlSave(ctx context.Context) (_node *Plan, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   plan.Table,
			Columns: plan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: plan.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Plan.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, plan.FieldID)
		for _, f := range fields {
			if !plan.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != plan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Recommend(); ok {
		_spec.SetField(plan.FieldRecommend, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedRecommend(); ok {
		_spec.AddField(plan.FieldRecommend, field.TypeInt, value)
	}
	if value, ok := puo.mutation.CategoryID(); ok {
		_spec.SetField(plan.FieldCategoryID, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.AddedCategoryID(); ok {
		_spec.AddField(plan.FieldCategoryID, field.TypeInt32, value)
	}
	if value, ok := puo.mutation.Category(); ok {
		_spec.SetField(plan.FieldCategory, field.TypeString, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(plan.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.URL(); ok {
		_spec.SetField(plan.FieldURL, field.TypeString, value)
	}
	if value, ok := puo.mutation.Utime(); ok {
		_spec.SetField(plan.FieldUtime, field.TypeTime, value)
	}
	_node = &Plan{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{plan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
