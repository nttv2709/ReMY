// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"remy/pkg/ent/event"
	"remy/pkg/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks     []Hook
	mutation  *EventMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetTitle sets the "title" field.
func (eu *EventUpdate) SetTitle(s string) *EventUpdate {
	eu.mutation.SetTitle(s)
	return eu
}

// SetStart sets the "start" field.
func (eu *EventUpdate) SetStart(t time.Time) *EventUpdate {
	eu.mutation.SetStart(t)
	return eu
}

// SetEnd sets the "end" field.
func (eu *EventUpdate) SetEnd(t time.Time) *EventUpdate {
	eu.mutation.SetEnd(t)
	return eu
}

// SetXPos sets the "xPos" field.
func (eu *EventUpdate) SetXPos(f float32) *EventUpdate {
	eu.mutation.ResetXPos()
	eu.mutation.SetXPos(f)
	return eu
}

// AddXPos adds f to the "xPos" field.
func (eu *EventUpdate) AddXPos(f float32) *EventUpdate {
	eu.mutation.AddXPos(f)
	return eu
}

// SetYPos sets the "yPos" field.
func (eu *EventUpdate) SetYPos(f float32) *EventUpdate {
	eu.mutation.ResetYPos()
	eu.mutation.SetYPos(f)
	return eu
}

// AddYPos adds f to the "yPos" field.
func (eu *EventUpdate) AddYPos(f float32) *EventUpdate {
	eu.mutation.AddYPos(f)
	return eu
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (eu *EventUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EventUpdate {
	eu.modifiers = append(eu.modifiers, modifiers...)
	return eu
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldTitle,
		})
	}
	if value, ok := eu.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldStart,
		})
	}
	if value, ok := eu.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldEnd,
		})
	}
	if value, ok := eu.mutation.XPos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: event.FieldXPos,
		})
	}
	if value, ok := eu.mutation.AddedXPos(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: event.FieldXPos,
		})
	}
	if value, ok := eu.mutation.YPos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: event.FieldYPos,
		})
	}
	if value, ok := eu.mutation.AddedYPos(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: event.FieldYPos,
		})
	}
	_spec.Modifiers = eu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *EventMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetTitle sets the "title" field.
func (euo *EventUpdateOne) SetTitle(s string) *EventUpdateOne {
	euo.mutation.SetTitle(s)
	return euo
}

// SetStart sets the "start" field.
func (euo *EventUpdateOne) SetStart(t time.Time) *EventUpdateOne {
	euo.mutation.SetStart(t)
	return euo
}

// SetEnd sets the "end" field.
func (euo *EventUpdateOne) SetEnd(t time.Time) *EventUpdateOne {
	euo.mutation.SetEnd(t)
	return euo
}

// SetXPos sets the "xPos" field.
func (euo *EventUpdateOne) SetXPos(f float32) *EventUpdateOne {
	euo.mutation.ResetXPos()
	euo.mutation.SetXPos(f)
	return euo
}

// AddXPos adds f to the "xPos" field.
func (euo *EventUpdateOne) AddXPos(f float32) *EventUpdateOne {
	euo.mutation.AddXPos(f)
	return euo
}

// SetYPos sets the "yPos" field.
func (euo *EventUpdateOne) SetYPos(f float32) *EventUpdateOne {
	euo.mutation.ResetYPos()
	euo.mutation.SetYPos(f)
	return euo
}

// AddYPos adds f to the "yPos" field.
func (euo *EventUpdateOne) AddYPos(f float32) *EventUpdateOne {
	euo.mutation.AddYPos(f)
	return euo
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, euo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Event)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (euo *EventUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *EventUpdateOne {
	euo.modifiers = append(euo.modifiers, modifiers...)
	return euo
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Event.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldTitle,
		})
	}
	if value, ok := euo.mutation.Start(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldStart,
		})
	}
	if value, ok := euo.mutation.End(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldEnd,
		})
	}
	if value, ok := euo.mutation.XPos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: event.FieldXPos,
		})
	}
	if value, ok := euo.mutation.AddedXPos(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: event.FieldXPos,
		})
	}
	if value, ok := euo.mutation.YPos(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: event.FieldYPos,
		})
	}
	if value, ok := euo.mutation.AddedYPos(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: event.FieldYPos,
		})
	}
	_spec.Modifiers = euo.modifiers
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
