// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/starryrbs/kfan/app/house/internal/data/ent/house"
)

// HouseCreate is the builder for creating a House entity.
type HouseCreate struct {
	config
	mutation *HouseMutation
	hooks    []Hook
}

// SetPrice sets the "Price" field.
func (hc *HouseCreate) SetPrice(f float64) *HouseCreate {
	hc.mutation.SetPrice(f)
	return hc
}

// SetTitle sets the "title" field.
func (hc *HouseCreate) SetTitle(s string) *HouseCreate {
	hc.mutation.SetTitle(s)
	return hc
}

// SetCommunity sets the "community" field.
func (hc *HouseCreate) SetCommunity(s string) *HouseCreate {
	hc.mutation.SetCommunity(s)
	return hc
}

// SetToiletCount sets the "toilet_count" field.
func (hc *HouseCreate) SetToiletCount(i int32) *HouseCreate {
	hc.mutation.SetToiletCount(i)
	return hc
}

// SetNillableToiletCount sets the "toilet_count" field if the given value is not nil.
func (hc *HouseCreate) SetNillableToiletCount(i *int32) *HouseCreate {
	if i != nil {
		hc.SetToiletCount(*i)
	}
	return hc
}

// SetKitchenCount sets the "kitchen_count" field.
func (hc *HouseCreate) SetKitchenCount(i int32) *HouseCreate {
	hc.mutation.SetKitchenCount(i)
	return hc
}

// SetNillableKitchenCount sets the "kitchen_count" field if the given value is not nil.
func (hc *HouseCreate) SetNillableKitchenCount(i *int32) *HouseCreate {
	if i != nil {
		hc.SetKitchenCount(*i)
	}
	return hc
}

// SetFloorCount sets the "floor_count" field.
func (hc *HouseCreate) SetFloorCount(i int32) *HouseCreate {
	hc.mutation.SetFloorCount(i)
	return hc
}

// SetNillableFloorCount sets the "floor_count" field if the given value is not nil.
func (hc *HouseCreate) SetNillableFloorCount(i *int32) *HouseCreate {
	if i != nil {
		hc.SetFloorCount(*i)
	}
	return hc
}

// SetHallCount sets the "hall_count" field.
func (hc *HouseCreate) SetHallCount(i int32) *HouseCreate {
	hc.mutation.SetHallCount(i)
	return hc
}

// SetNillableHallCount sets the "hall_count" field if the given value is not nil.
func (hc *HouseCreate) SetNillableHallCount(i *int32) *HouseCreate {
	if i != nil {
		hc.SetHallCount(*i)
	}
	return hc
}

// SetRoomCount sets the "room_count" field.
func (hc *HouseCreate) SetRoomCount(i int32) *HouseCreate {
	hc.mutation.SetRoomCount(i)
	return hc
}

// SetNillableRoomCount sets the "room_count" field if the given value is not nil.
func (hc *HouseCreate) SetNillableRoomCount(i *int32) *HouseCreate {
	if i != nil {
		hc.SetRoomCount(*i)
	}
	return hc
}

// SetID sets the "id" field.
func (hc *HouseCreate) SetID(i int64) *HouseCreate {
	hc.mutation.SetID(i)
	return hc
}

// Mutation returns the HouseMutation object of the builder.
func (hc *HouseCreate) Mutation() *HouseMutation {
	return hc.mutation
}

// Save creates the House in the database.
func (hc *HouseCreate) Save(ctx context.Context) (*House, error) {
	var (
		err  error
		node *House
	)
	hc.defaults()
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HouseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			if node, err = hc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			if hc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HouseCreate) SaveX(ctx context.Context) *House {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HouseCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HouseCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HouseCreate) defaults() {
	if _, ok := hc.mutation.ToiletCount(); !ok {
		v := house.DefaultToiletCount
		hc.mutation.SetToiletCount(v)
	}
	if _, ok := hc.mutation.KitchenCount(); !ok {
		v := house.DefaultKitchenCount
		hc.mutation.SetKitchenCount(v)
	}
	if _, ok := hc.mutation.FloorCount(); !ok {
		v := house.DefaultFloorCount
		hc.mutation.SetFloorCount(v)
	}
	if _, ok := hc.mutation.HallCount(); !ok {
		v := house.DefaultHallCount
		hc.mutation.SetHallCount(v)
	}
	if _, ok := hc.mutation.RoomCount(); !ok {
		v := house.DefaultRoomCount
		hc.mutation.SetRoomCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HouseCreate) check() error {
	if _, ok := hc.mutation.Price(); !ok {
		return &ValidationError{Name: "Price", err: errors.New(`ent: missing required field "Price"`)}
	}
	if _, ok := hc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := hc.mutation.Title(); ok {
		if err := house.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := hc.mutation.Community(); !ok {
		return &ValidationError{Name: "community", err: errors.New(`ent: missing required field "community"`)}
	}
	if _, ok := hc.mutation.ToiletCount(); !ok {
		return &ValidationError{Name: "toilet_count", err: errors.New(`ent: missing required field "toilet_count"`)}
	}
	if _, ok := hc.mutation.KitchenCount(); !ok {
		return &ValidationError{Name: "kitchen_count", err: errors.New(`ent: missing required field "kitchen_count"`)}
	}
	if _, ok := hc.mutation.FloorCount(); !ok {
		return &ValidationError{Name: "floor_count", err: errors.New(`ent: missing required field "floor_count"`)}
	}
	if _, ok := hc.mutation.HallCount(); !ok {
		return &ValidationError{Name: "hall_count", err: errors.New(`ent: missing required field "hall_count"`)}
	}
	if _, ok := hc.mutation.RoomCount(); !ok {
		return &ValidationError{Name: "room_count", err: errors.New(`ent: missing required field "room_count"`)}
	}
	return nil
}

func (hc *HouseCreate) sqlSave(ctx context.Context) (*House, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (hc *HouseCreate) createSpec() (*House, *sqlgraph.CreateSpec) {
	var (
		_node = &House{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: house.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: house.FieldID,
			},
		}
	)
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: house.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := hc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: house.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := hc.mutation.Community(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: house.FieldCommunity,
		})
		_node.Community = value
	}
	if value, ok := hc.mutation.ToiletCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: house.FieldToiletCount,
		})
		_node.ToiletCount = value
	}
	if value, ok := hc.mutation.KitchenCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: house.FieldKitchenCount,
		})
		_node.KitchenCount = value
	}
	if value, ok := hc.mutation.FloorCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: house.FieldFloorCount,
		})
		_node.FloorCount = value
	}
	if value, ok := hc.mutation.HallCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: house.FieldHallCount,
		})
		_node.HallCount = value
	}
	if value, ok := hc.mutation.RoomCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: house.FieldRoomCount,
		})
		_node.RoomCount = value
	}
	return _node, _spec
}

// HouseCreateBulk is the builder for creating many House entities in bulk.
type HouseCreateBulk struct {
	config
	builders []*HouseCreate
}

// Save creates the House entities in the database.
func (hcb *HouseCreateBulk) Save(ctx context.Context) ([]*House, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*House, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HouseMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HouseCreateBulk) SaveX(ctx context.Context) []*House {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HouseCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HouseCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
