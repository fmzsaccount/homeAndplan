// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gy_home/internal/data/ent/site"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SiteCreate is the builder for creating a Site entity.
type SiteCreate struct {
	config
	mutation *SiteMutation
	hooks    []Hook
}

// SetSort sets the "sort" field.
func (sc *SiteCreate) SetSort(i int32) *SiteCreate {
	sc.mutation.SetSort(i)
	return sc
}

// SetStatus sets the "status" field.
func (sc *SiteCreate) SetStatus(i int) *SiteCreate {
	sc.mutation.SetStatus(i)
	return sc
}

// SetName sets the "name" field.
func (sc *SiteCreate) SetName(s string) *SiteCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetLink sets the "link" field.
func (sc *SiteCreate) SetLink(s string) *SiteCreate {
	sc.mutation.SetLink(s)
	return sc
}

// SetCategory sets the "category" field.
func (sc *SiteCreate) SetCategory(s string) *SiteCreate {
	sc.mutation.SetCategory(s)
	return sc
}

// SetID sets the "id" field.
func (sc *SiteCreate) SetID(i int32) *SiteCreate {
	sc.mutation.SetID(i)
	return sc
}

// Mutation returns the SiteMutation object of the builder.
func (sc *SiteCreate) Mutation() *SiteMutation {
	return sc.mutation
}

// Save creates the Site in the database.
func (sc *SiteCreate) Save(ctx context.Context) (*Site, error) {
	var (
		err  error
		node *Site
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SiteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Site)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SiteMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SiteCreate) SaveX(ctx context.Context) *Site {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SiteCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SiteCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SiteCreate) check() error {
	if _, ok := sc.mutation.Sort(); !ok {
		return &ValidationError{Name: "sort", err: errors.New(`ent: missing required field "Site.sort"`)}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Site.status"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Site.name"`)}
	}
	if _, ok := sc.mutation.Link(); !ok {
		return &ValidationError{Name: "link", err: errors.New(`ent: missing required field "Site.link"`)}
	}
	if _, ok := sc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Site.category"`)}
	}
	return nil
}

func (sc *SiteCreate) sqlSave(ctx context.Context) (*Site, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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

func (sc *SiteCreate) createSpec() (*Site, *sqlgraph.CreateSpec) {
	var (
		_node = &Site{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: site.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: site.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Sort(); ok {
		_spec.SetField(site.FieldSort, field.TypeInt32, value)
		_node.Sort = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(site.FieldStatus, field.TypeInt, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(site.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Link(); ok {
		_spec.SetField(site.FieldLink, field.TypeString, value)
		_node.Link = value
	}
	if value, ok := sc.mutation.Category(); ok {
		_spec.SetField(site.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	return _node, _spec
}

// SiteCreateBulk is the builder for creating many Site entities in bulk.
type SiteCreateBulk struct {
	config
	builders []*SiteCreate
}

// Save creates the Site entities in the database.
func (scb *SiteCreateBulk) Save(ctx context.Context) ([]*Site, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Site, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SiteMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SiteCreateBulk) SaveX(ctx context.Context) []*Site {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SiteCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SiteCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}