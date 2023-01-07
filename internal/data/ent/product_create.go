// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gy_home/internal/biz"
	"gy_home/internal/data/ent/product"
	"gy_home/internal/data/ent/productcategory"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (pc *ProductCreate) SetStatus(i int) *ProductCreate {
	pc.mutation.SetStatus(i)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *ProductCreate) SetNillableStatus(i *int) *ProductCreate {
	if i != nil {
		pc.SetStatus(*i)
	}
	return pc
}

// SetRecommend sets the "recommend" field.
func (pc *ProductCreate) SetRecommend(i int) *ProductCreate {
	pc.mutation.SetRecommend(i)
	return pc
}

// SetNillableRecommend sets the "recommend" field if the given value is not nil.
func (pc *ProductCreate) SetNillableRecommend(i *int) *ProductCreate {
	if i != nil {
		pc.SetRecommend(*i)
	}
	return pc
}

// SetHistory sets the "history" field.
func (pc *ProductCreate) SetHistory(i int32) *ProductCreate {
	pc.mutation.SetHistory(i)
	return pc
}

// SetNillableHistory sets the "history" field if the given value is not nil.
func (pc *ProductCreate) SetNillableHistory(i *int32) *ProductCreate {
	if i != nil {
		pc.SetHistory(*i)
	}
	return pc
}

// SetCategoryID sets the "category_id" field.
func (pc *ProductCreate) SetCategoryID(i int32) *ProductCreate {
	pc.mutation.SetCategoryID(i)
	return pc
}

// SetNillableCategoryID sets the "category_id" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCategoryID(i *int32) *ProductCreate {
	if i != nil {
		pc.SetCategoryID(*i)
	}
	return pc
}

// SetCategoryName sets the "category_name" field.
func (pc *ProductCreate) SetCategoryName(s string) *ProductCreate {
	pc.mutation.SetCategoryName(s)
	return pc
}

// SetNillableCategoryName sets the "category_name" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCategoryName(s *string) *ProductCreate {
	if s != nil {
		pc.SetCategoryName(*s)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *ProductCreate) SetNillableName(s *string) *ProductCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetContent sets the "content" field.
func (pc *ProductCreate) SetContent(s string) *ProductCreate {
	pc.mutation.SetContent(s)
	return pc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (pc *ProductCreate) SetNillableContent(s *string) *ProductCreate {
	if s != nil {
		pc.SetContent(*s)
	}
	return pc
}

// SetUtime sets the "utime" field.
func (pc *ProductCreate) SetUtime(t time.Time) *ProductCreate {
	pc.mutation.SetUtime(t)
	return pc
}

// SetServiceLink sets the "service_link" field.
func (pc *ProductCreate) SetServiceLink(s string) *ProductCreate {
	pc.mutation.SetServiceLink(s)
	return pc
}

// SetNillableServiceLink sets the "service_link" field if the given value is not nil.
func (pc *ProductCreate) SetNillableServiceLink(s *string) *ProductCreate {
	if s != nil {
		pc.SetServiceLink(*s)
	}
	return pc
}

// SetApplyLink sets the "apply_link" field.
func (pc *ProductCreate) SetApplyLink(s string) *ProductCreate {
	pc.mutation.SetApplyLink(s)
	return pc
}

// SetNillableApplyLink sets the "apply_link" field if the given value is not nil.
func (pc *ProductCreate) SetNillableApplyLink(s *string) *ProductCreate {
	if s != nil {
		pc.SetApplyLink(*s)
	}
	return pc
}

// SetJSON sets the "json" field.
func (pc *ProductCreate) SetJSON(b biz.Product) *ProductCreate {
	pc.mutation.SetJSON(b)
	return pc
}

// SetID sets the "id" field.
func (pc *ProductCreate) SetID(i int32) *ProductCreate {
	pc.mutation.SetID(i)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ProductCreate) SetNillableID(i *int32) *ProductCreate {
	if i != nil {
		pc.SetID(*i)
	}
	return pc
}

// SetParentIDID sets the "parent_id" edge to the ProductCategory entity by ID.
func (pc *ProductCreate) SetParentIDID(id int32) *ProductCreate {
	pc.mutation.SetParentIDID(id)
	return pc
}

// SetNillableParentIDID sets the "parent_id" edge to the ProductCategory entity by ID if the given value is not nil.
func (pc *ProductCreate) SetNillableParentIDID(id *int32) *ProductCreate {
	if id != nil {
		pc = pc.SetParentIDID(*id)
	}
	return pc
}

// SetParentID sets the "parent_id" edge to the ProductCategory entity.
func (pc *ProductCreate) SetParentID(p *ProductCategory) *ProductCreate {
	return pc.SetParentIDID(p.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	var (
		err  error
		node *Product
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Product)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProductMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.Status(); !ok {
		v := product.DefaultStatus
		pc.mutation.SetStatus(v)
	}
	if _, ok := pc.mutation.Recommend(); !ok {
		v := product.DefaultRecommend
		pc.mutation.SetRecommend(v)
	}
	if _, ok := pc.mutation.History(); !ok {
		v := product.DefaultHistory
		pc.mutation.SetHistory(v)
	}
	if _, ok := pc.mutation.CategoryID(); !ok {
		v := product.DefaultCategoryID
		pc.mutation.SetCategoryID(v)
	}
	if _, ok := pc.mutation.CategoryName(); !ok {
		v := product.DefaultCategoryName
		pc.mutation.SetCategoryName(v)
	}
	if _, ok := pc.mutation.Name(); !ok {
		v := product.DefaultName
		pc.mutation.SetName(v)
	}
	if _, ok := pc.mutation.Content(); !ok {
		v := product.DefaultContent
		pc.mutation.SetContent(v)
	}
	if _, ok := pc.mutation.ServiceLink(); !ok {
		v := product.DefaultServiceLink
		pc.mutation.SetServiceLink(v)
	}
	if _, ok := pc.mutation.ApplyLink(); !ok {
		v := product.DefaultApplyLink
		pc.mutation.SetApplyLink(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := product.DefaultID
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Product.status"`)}
	}
	if _, ok := pc.mutation.Recommend(); !ok {
		return &ValidationError{Name: "recommend", err: errors.New(`ent: missing required field "Product.recommend"`)}
	}
	if _, ok := pc.mutation.History(); !ok {
		return &ValidationError{Name: "history", err: errors.New(`ent: missing required field "Product.history"`)}
	}
	if _, ok := pc.mutation.CategoryName(); !ok {
		return &ValidationError{Name: "category_name", err: errors.New(`ent: missing required field "Product.category_name"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Product.name"`)}
	}
	if _, ok := pc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Product.content"`)}
	}
	if _, ok := pc.mutation.Utime(); !ok {
		return &ValidationError{Name: "utime", err: errors.New(`ent: missing required field "Product.utime"`)}
	}
	if _, ok := pc.mutation.ServiceLink(); !ok {
		return &ValidationError{Name: "service_link", err: errors.New(`ent: missing required field "Product.service_link"`)}
	}
	if _, ok := pc.mutation.ApplyLink(); !ok {
		return &ValidationError{Name: "apply_link", err: errors.New(`ent: missing required field "Product.apply_link"`)}
	}
	if _, ok := pc.mutation.JSON(); !ok {
		return &ValidationError{Name: "json", err: errors.New(`ent: missing required field "Product.json"`)}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: product.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: product.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(product.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.Recommend(); ok {
		_spec.SetField(product.FieldRecommend, field.TypeInt, value)
		_node.Recommend = value
	}
	if value, ok := pc.mutation.History(); ok {
		_spec.SetField(product.FieldHistory, field.TypeInt32, value)
		_node.History = value
	}
	if value, ok := pc.mutation.CategoryName(); ok {
		_spec.SetField(product.FieldCategoryName, field.TypeString, value)
		_node.CategoryName = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Content(); ok {
		_spec.SetField(product.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := pc.mutation.Utime(); ok {
		_spec.SetField(product.FieldUtime, field.TypeTime, value)
		_node.Utime = value
	}
	if value, ok := pc.mutation.ServiceLink(); ok {
		_spec.SetField(product.FieldServiceLink, field.TypeString, value)
		_node.ServiceLink = value
	}
	if value, ok := pc.mutation.ApplyLink(); ok {
		_spec.SetField(product.FieldApplyLink, field.TypeString, value)
		_node.ApplyLink = value
	}
	if value, ok := pc.mutation.JSON(); ok {
		_spec.SetField(product.FieldJSON, field.TypeJSON, value)
		_node.JSON = value
	}
	if nodes := pc.mutation.ParentIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.ParentIDTable,
			Columns: []string{product.ParentIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: productcategory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CategoryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
