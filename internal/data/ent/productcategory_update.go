// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gy_home/internal/data/ent/predicate"
	"gy_home/internal/data/ent/product"
	"gy_home/internal/data/ent/productcategory"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCategoryUpdate is the builder for updating ProductCategory entities.
type ProductCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *ProductCategoryMutation
}

// Where appends a list predicates to the ProductCategoryUpdate builder.
func (pcu *ProductCategoryUpdate) Where(ps ...predicate.ProductCategory) *ProductCategoryUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetStatus sets the "status" field.
func (pcu *ProductCategoryUpdate) SetStatus(i int) *ProductCategoryUpdate {
	pcu.mutation.ResetStatus()
	pcu.mutation.SetStatus(i)
	return pcu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pcu *ProductCategoryUpdate) SetNillableStatus(i *int) *ProductCategoryUpdate {
	if i != nil {
		pcu.SetStatus(*i)
	}
	return pcu
}

// AddStatus adds i to the "status" field.
func (pcu *ProductCategoryUpdate) AddStatus(i int) *ProductCategoryUpdate {
	pcu.mutation.AddStatus(i)
	return pcu
}

// SetLevel sets the "level" field.
func (pcu *ProductCategoryUpdate) SetLevel(i int32) *ProductCategoryUpdate {
	pcu.mutation.ResetLevel()
	pcu.mutation.SetLevel(i)
	return pcu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (pcu *ProductCategoryUpdate) SetNillableLevel(i *int32) *ProductCategoryUpdate {
	if i != nil {
		pcu.SetLevel(*i)
	}
	return pcu
}

// AddLevel adds i to the "level" field.
func (pcu *ProductCategoryUpdate) AddLevel(i int32) *ProductCategoryUpdate {
	pcu.mutation.AddLevel(i)
	return pcu
}

// SetParentID sets the "parent_id" field.
func (pcu *ProductCategoryUpdate) SetParentID(i int32) *ProductCategoryUpdate {
	pcu.mutation.SetParentID(i)
	return pcu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pcu *ProductCategoryUpdate) SetNillableParentID(i *int32) *ProductCategoryUpdate {
	if i != nil {
		pcu.SetParentID(*i)
	}
	return pcu
}

// ClearParentID clears the value of the "parent_id" field.
func (pcu *ProductCategoryUpdate) ClearParentID() *ProductCategoryUpdate {
	pcu.mutation.ClearParentID()
	return pcu
}

// SetSort sets the "sort" field.
func (pcu *ProductCategoryUpdate) SetSort(i int32) *ProductCategoryUpdate {
	pcu.mutation.ResetSort()
	pcu.mutation.SetSort(i)
	return pcu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (pcu *ProductCategoryUpdate) SetNillableSort(i *int32) *ProductCategoryUpdate {
	if i != nil {
		pcu.SetSort(*i)
	}
	return pcu
}

// AddSort adds i to the "sort" field.
func (pcu *ProductCategoryUpdate) AddSort(i int32) *ProductCategoryUpdate {
	pcu.mutation.AddSort(i)
	return pcu
}

// SetName sets the "name" field.
func (pcu *ProductCategoryUpdate) SetName(s string) *ProductCategoryUpdate {
	pcu.mutation.SetName(s)
	return pcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pcu *ProductCategoryUpdate) SetNillableName(s *string) *ProductCategoryUpdate {
	if s != nil {
		pcu.SetName(*s)
	}
	return pcu
}

// SetParent sets the "parent" edge to the ProductCategory entity.
func (pcu *ProductCategoryUpdate) SetParent(p *ProductCategory) *ProductCategoryUpdate {
	return pcu.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the ProductCategory entity by IDs.
func (pcu *ProductCategoryUpdate) AddChildIDs(ids ...int32) *ProductCategoryUpdate {
	pcu.mutation.AddChildIDs(ids...)
	return pcu
}

// AddChildren adds the "children" edges to the ProductCategory entity.
func (pcu *ProductCategoryUpdate) AddChildren(p ...*ProductCategory) *ProductCategoryUpdate {
	ids := make([]int32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcu.AddChildIDs(ids...)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (pcu *ProductCategoryUpdate) AddProductIDs(ids ...int32) *ProductCategoryUpdate {
	pcu.mutation.AddProductIDs(ids...)
	return pcu
}

// AddProducts adds the "products" edges to the Product entity.
func (pcu *ProductCategoryUpdate) AddProducts(p ...*Product) *ProductCategoryUpdate {
	ids := make([]int32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcu.AddProductIDs(ids...)
}

// Mutation returns the ProductCategoryMutation object of the builder.
func (pcu *ProductCategoryUpdate) Mutation() *ProductCategoryMutation {
	return pcu.mutation
}

// ClearParent clears the "parent" edge to the ProductCategory entity.
func (pcu *ProductCategoryUpdate) ClearParent() *ProductCategoryUpdate {
	pcu.mutation.ClearParent()
	return pcu
}

// ClearChildren clears all "children" edges to the ProductCategory entity.
func (pcu *ProductCategoryUpdate) ClearChildren() *ProductCategoryUpdate {
	pcu.mutation.ClearChildren()
	return pcu
}

// RemoveChildIDs removes the "children" edge to ProductCategory entities by IDs.
func (pcu *ProductCategoryUpdate) RemoveChildIDs(ids ...int32) *ProductCategoryUpdate {
	pcu.mutation.RemoveChildIDs(ids...)
	return pcu
}

// RemoveChildren removes "children" edges to ProductCategory entities.
func (pcu *ProductCategoryUpdate) RemoveChildren(p ...*ProductCategory) *ProductCategoryUpdate {
	ids := make([]int32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcu.RemoveChildIDs(ids...)
}

// ClearProducts clears all "products" edges to the Product entity.
func (pcu *ProductCategoryUpdate) ClearProducts() *ProductCategoryUpdate {
	pcu.mutation.ClearProducts()
	return pcu
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (pcu *ProductCategoryUpdate) RemoveProductIDs(ids ...int32) *ProductCategoryUpdate {
	pcu.mutation.RemoveProductIDs(ids...)
	return pcu
}

// RemoveProducts removes "products" edges to Product entities.
func (pcu *ProductCategoryUpdate) RemoveProducts(p ...*Product) *ProductCategoryUpdate {
	ids := make([]int32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *ProductCategoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pcu.hooks) == 0 {
		affected, err = pcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pcu.mutation = mutation
			affected, err = pcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pcu.hooks) - 1; i >= 0; i-- {
			if pcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *ProductCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *ProductCategoryUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *ProductCategoryUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcu *ProductCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productcategory.Table,
			Columns: productcategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: productcategory.FieldID,
			},
		},
	}
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.Status(); ok {
		_spec.SetField(productcategory.FieldStatus, field.TypeInt, value)
	}
	if value, ok := pcu.mutation.AddedStatus(); ok {
		_spec.AddField(productcategory.FieldStatus, field.TypeInt, value)
	}
	if value, ok := pcu.mutation.Level(); ok {
		_spec.SetField(productcategory.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := pcu.mutation.AddedLevel(); ok {
		_spec.AddField(productcategory.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := pcu.mutation.Sort(); ok {
		_spec.SetField(productcategory.FieldSort, field.TypeInt32, value)
	}
	if value, ok := pcu.mutation.AddedSort(); ok {
		_spec.AddField(productcategory.FieldSort, field.TypeInt32, value)
	}
	if value, ok := pcu.mutation.Name(); ok {
		_spec.SetField(productcategory.FieldName, field.TypeString, value)
	}
	if pcu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcategory.ParentTable,
			Columns: []string{productcategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: productcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcategory.ParentTable,
			Columns: []string{productcategory.ParentColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ChildrenTable,
			Columns: []string{productcategory.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: productcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !pcu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ChildrenTable,
			Columns: []string{productcategory.ChildrenColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ChildrenTable,
			Columns: []string{productcategory.ChildrenColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ProductsTable,
			Columns: []string{productcategory.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RemovedProductsIDs(); len(nodes) > 0 && !pcu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ProductsTable,
			Columns: []string{productcategory.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ProductsTable,
			Columns: []string{productcategory.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ProductCategoryUpdateOne is the builder for updating a single ProductCategory entity.
type ProductCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductCategoryMutation
}

// SetStatus sets the "status" field.
func (pcuo *ProductCategoryUpdateOne) SetStatus(i int) *ProductCategoryUpdateOne {
	pcuo.mutation.ResetStatus()
	pcuo.mutation.SetStatus(i)
	return pcuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pcuo *ProductCategoryUpdateOne) SetNillableStatus(i *int) *ProductCategoryUpdateOne {
	if i != nil {
		pcuo.SetStatus(*i)
	}
	return pcuo
}

// AddStatus adds i to the "status" field.
func (pcuo *ProductCategoryUpdateOne) AddStatus(i int) *ProductCategoryUpdateOne {
	pcuo.mutation.AddStatus(i)
	return pcuo
}

// SetLevel sets the "level" field.
func (pcuo *ProductCategoryUpdateOne) SetLevel(i int32) *ProductCategoryUpdateOne {
	pcuo.mutation.ResetLevel()
	pcuo.mutation.SetLevel(i)
	return pcuo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (pcuo *ProductCategoryUpdateOne) SetNillableLevel(i *int32) *ProductCategoryUpdateOne {
	if i != nil {
		pcuo.SetLevel(*i)
	}
	return pcuo
}

// AddLevel adds i to the "level" field.
func (pcuo *ProductCategoryUpdateOne) AddLevel(i int32) *ProductCategoryUpdateOne {
	pcuo.mutation.AddLevel(i)
	return pcuo
}

// SetParentID sets the "parent_id" field.
func (pcuo *ProductCategoryUpdateOne) SetParentID(i int32) *ProductCategoryUpdateOne {
	pcuo.mutation.SetParentID(i)
	return pcuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pcuo *ProductCategoryUpdateOne) SetNillableParentID(i *int32) *ProductCategoryUpdateOne {
	if i != nil {
		pcuo.SetParentID(*i)
	}
	return pcuo
}

// ClearParentID clears the value of the "parent_id" field.
func (pcuo *ProductCategoryUpdateOne) ClearParentID() *ProductCategoryUpdateOne {
	pcuo.mutation.ClearParentID()
	return pcuo
}

// SetSort sets the "sort" field.
func (pcuo *ProductCategoryUpdateOne) SetSort(i int32) *ProductCategoryUpdateOne {
	pcuo.mutation.ResetSort()
	pcuo.mutation.SetSort(i)
	return pcuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (pcuo *ProductCategoryUpdateOne) SetNillableSort(i *int32) *ProductCategoryUpdateOne {
	if i != nil {
		pcuo.SetSort(*i)
	}
	return pcuo
}

// AddSort adds i to the "sort" field.
func (pcuo *ProductCategoryUpdateOne) AddSort(i int32) *ProductCategoryUpdateOne {
	pcuo.mutation.AddSort(i)
	return pcuo
}

// SetName sets the "name" field.
func (pcuo *ProductCategoryUpdateOne) SetName(s string) *ProductCategoryUpdateOne {
	pcuo.mutation.SetName(s)
	return pcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pcuo *ProductCategoryUpdateOne) SetNillableName(s *string) *ProductCategoryUpdateOne {
	if s != nil {
		pcuo.SetName(*s)
	}
	return pcuo
}

// SetParent sets the "parent" edge to the ProductCategory entity.
func (pcuo *ProductCategoryUpdateOne) SetParent(p *ProductCategory) *ProductCategoryUpdateOne {
	return pcuo.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the ProductCategory entity by IDs.
func (pcuo *ProductCategoryUpdateOne) AddChildIDs(ids ...int32) *ProductCategoryUpdateOne {
	pcuo.mutation.AddChildIDs(ids...)
	return pcuo
}

// AddChildren adds the "children" edges to the ProductCategory entity.
func (pcuo *ProductCategoryUpdateOne) AddChildren(p ...*ProductCategory) *ProductCategoryUpdateOne {
	ids := make([]int32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcuo.AddChildIDs(ids...)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (pcuo *ProductCategoryUpdateOne) AddProductIDs(ids ...int32) *ProductCategoryUpdateOne {
	pcuo.mutation.AddProductIDs(ids...)
	return pcuo
}

// AddProducts adds the "products" edges to the Product entity.
func (pcuo *ProductCategoryUpdateOne) AddProducts(p ...*Product) *ProductCategoryUpdateOne {
	ids := make([]int32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcuo.AddProductIDs(ids...)
}

// Mutation returns the ProductCategoryMutation object of the builder.
func (pcuo *ProductCategoryUpdateOne) Mutation() *ProductCategoryMutation {
	return pcuo.mutation
}

// ClearParent clears the "parent" edge to the ProductCategory entity.
func (pcuo *ProductCategoryUpdateOne) ClearParent() *ProductCategoryUpdateOne {
	pcuo.mutation.ClearParent()
	return pcuo
}

// ClearChildren clears all "children" edges to the ProductCategory entity.
func (pcuo *ProductCategoryUpdateOne) ClearChildren() *ProductCategoryUpdateOne {
	pcuo.mutation.ClearChildren()
	return pcuo
}

// RemoveChildIDs removes the "children" edge to ProductCategory entities by IDs.
func (pcuo *ProductCategoryUpdateOne) RemoveChildIDs(ids ...int32) *ProductCategoryUpdateOne {
	pcuo.mutation.RemoveChildIDs(ids...)
	return pcuo
}

// RemoveChildren removes "children" edges to ProductCategory entities.
func (pcuo *ProductCategoryUpdateOne) RemoveChildren(p ...*ProductCategory) *ProductCategoryUpdateOne {
	ids := make([]int32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcuo.RemoveChildIDs(ids...)
}

// ClearProducts clears all "products" edges to the Product entity.
func (pcuo *ProductCategoryUpdateOne) ClearProducts() *ProductCategoryUpdateOne {
	pcuo.mutation.ClearProducts()
	return pcuo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (pcuo *ProductCategoryUpdateOne) RemoveProductIDs(ids ...int32) *ProductCategoryUpdateOne {
	pcuo.mutation.RemoveProductIDs(ids...)
	return pcuo
}

// RemoveProducts removes "products" edges to Product entities.
func (pcuo *ProductCategoryUpdateOne) RemoveProducts(p ...*Product) *ProductCategoryUpdateOne {
	ids := make([]int32, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcuo.RemoveProductIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *ProductCategoryUpdateOne) Select(field string, fields ...string) *ProductCategoryUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated ProductCategory entity.
func (pcuo *ProductCategoryUpdateOne) Save(ctx context.Context) (*ProductCategory, error) {
	var (
		err  error
		node *ProductCategory
	)
	if len(pcuo.hooks) == 0 {
		node, err = pcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProductCategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pcuo.mutation = mutation
			node, err = pcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pcuo.hooks) - 1; i >= 0; i-- {
			if pcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ProductCategory)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProductCategoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *ProductCategoryUpdateOne) SaveX(ctx context.Context) *ProductCategory {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *ProductCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *ProductCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcuo *ProductCategoryUpdateOne) sqlSave(ctx context.Context) (_node *ProductCategory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   productcategory.Table,
			Columns: productcategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: productcategory.FieldID,
			},
		},
	}
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProductCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productcategory.FieldID)
		for _, f := range fields {
			if !productcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.Status(); ok {
		_spec.SetField(productcategory.FieldStatus, field.TypeInt, value)
	}
	if value, ok := pcuo.mutation.AddedStatus(); ok {
		_spec.AddField(productcategory.FieldStatus, field.TypeInt, value)
	}
	if value, ok := pcuo.mutation.Level(); ok {
		_spec.SetField(productcategory.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := pcuo.mutation.AddedLevel(); ok {
		_spec.AddField(productcategory.FieldLevel, field.TypeInt32, value)
	}
	if value, ok := pcuo.mutation.Sort(); ok {
		_spec.SetField(productcategory.FieldSort, field.TypeInt32, value)
	}
	if value, ok := pcuo.mutation.AddedSort(); ok {
		_spec.AddField(productcategory.FieldSort, field.TypeInt32, value)
	}
	if value, ok := pcuo.mutation.Name(); ok {
		_spec.SetField(productcategory.FieldName, field.TypeString, value)
	}
	if pcuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcategory.ParentTable,
			Columns: []string{productcategory.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: productcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcategory.ParentTable,
			Columns: []string{productcategory.ParentColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ChildrenTable,
			Columns: []string{productcategory.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: productcategory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !pcuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ChildrenTable,
			Columns: []string{productcategory.ChildrenColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ChildrenTable,
			Columns: []string{productcategory.ChildrenColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ProductsTable,
			Columns: []string{productcategory.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: product.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !pcuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ProductsTable,
			Columns: []string{productcategory.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productcategory.ProductsTable,
			Columns: []string{productcategory.ProductsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProductCategory{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
