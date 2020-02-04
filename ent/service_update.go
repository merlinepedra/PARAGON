// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/event"
	"github.com/kcarretto/paragon/ent/predicate"
	"github.com/kcarretto/paragon/ent/service"
	"github.com/kcarretto/paragon/ent/tag"
)

// ServiceUpdate is the builder for updating Service entities.
type ServiceUpdate struct {
	config
	Name          *string
	PubKey        *string
	IsActivated   *bool
	tag           map[int]struct{}
	events        map[int]struct{}
	clearedTag    bool
	removedEvents map[int]struct{}
	predicates    []predicate.Service
}

// Where adds a new predicate for the builder.
func (su *ServiceUpdate) Where(ps ...predicate.Service) *ServiceUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetName sets the Name field.
func (su *ServiceUpdate) SetName(s string) *ServiceUpdate {
	su.Name = &s
	return su
}

// SetPubKey sets the PubKey field.
func (su *ServiceUpdate) SetPubKey(s string) *ServiceUpdate {
	su.PubKey = &s
	return su
}

// SetIsActivated sets the IsActivated field.
func (su *ServiceUpdate) SetIsActivated(b bool) *ServiceUpdate {
	su.IsActivated = &b
	return su
}

// SetNillableIsActivated sets the IsActivated field if the given value is not nil.
func (su *ServiceUpdate) SetNillableIsActivated(b *bool) *ServiceUpdate {
	if b != nil {
		su.SetIsActivated(*b)
	}
	return su
}

// SetTagID sets the tag edge to Tag by id.
func (su *ServiceUpdate) SetTagID(id int) *ServiceUpdate {
	if su.tag == nil {
		su.tag = make(map[int]struct{})
	}
	su.tag[id] = struct{}{}
	return su
}

// SetTag sets the tag edge to Tag.
func (su *ServiceUpdate) SetTag(t *Tag) *ServiceUpdate {
	return su.SetTagID(t.ID)
}

// AddEventIDs adds the events edge to Event by ids.
func (su *ServiceUpdate) AddEventIDs(ids ...int) *ServiceUpdate {
	if su.events == nil {
		su.events = make(map[int]struct{})
	}
	for i := range ids {
		su.events[ids[i]] = struct{}{}
	}
	return su
}

// AddEvents adds the events edges to Event.
func (su *ServiceUpdate) AddEvents(e ...*Event) *ServiceUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.AddEventIDs(ids...)
}

// ClearTag clears the tag edge to Tag.
func (su *ServiceUpdate) ClearTag() *ServiceUpdate {
	su.clearedTag = true
	return su
}

// RemoveEventIDs removes the events edge to Event by ids.
func (su *ServiceUpdate) RemoveEventIDs(ids ...int) *ServiceUpdate {
	if su.removedEvents == nil {
		su.removedEvents = make(map[int]struct{})
	}
	for i := range ids {
		su.removedEvents[ids[i]] = struct{}{}
	}
	return su
}

// RemoveEvents removes events edges to Event.
func (su *ServiceUpdate) RemoveEvents(e ...*Event) *ServiceUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.RemoveEventIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *ServiceUpdate) Save(ctx context.Context) (int, error) {
	if su.Name != nil {
		if err := service.NameValidator(*su.Name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"Name\": %v", err)
		}
	}
	if su.PubKey != nil {
		if err := service.PubKeyValidator(*su.PubKey); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"PubKey\": %v", err)
		}
	}
	if len(su.tag) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"tag\"")
	}
	if su.clearedTag && su.tag == nil {
		return 0, errors.New("ent: clearing a unique edge \"tag\"")
	}
	return su.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (su *ServiceUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ServiceUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ServiceUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *ServiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   service.Table,
			Columns: service.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: service.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := su.Name; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: service.FieldName,
		})
	}
	if value := su.PubKey; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: service.FieldPubKey,
		})
	}
	if value := su.IsActivated; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: service.FieldIsActivated,
		})
	}
	if su.clearedTag {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   service.TagTable,
			Columns: []string{service.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := su.tag; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   service.TagTable,
			Columns: []string{service.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if nodes := su.removedEvents; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   service.EventsTable,
			Columns: []string{service.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := su.events; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   service.EventsTable,
			Columns: []string{service.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ServiceUpdateOne is the builder for updating a single Service entity.
type ServiceUpdateOne struct {
	config
	id            int
	Name          *string
	PubKey        *string
	IsActivated   *bool
	tag           map[int]struct{}
	events        map[int]struct{}
	clearedTag    bool
	removedEvents map[int]struct{}
}

// SetName sets the Name field.
func (suo *ServiceUpdateOne) SetName(s string) *ServiceUpdateOne {
	suo.Name = &s
	return suo
}

// SetPubKey sets the PubKey field.
func (suo *ServiceUpdateOne) SetPubKey(s string) *ServiceUpdateOne {
	suo.PubKey = &s
	return suo
}

// SetIsActivated sets the IsActivated field.
func (suo *ServiceUpdateOne) SetIsActivated(b bool) *ServiceUpdateOne {
	suo.IsActivated = &b
	return suo
}

// SetNillableIsActivated sets the IsActivated field if the given value is not nil.
func (suo *ServiceUpdateOne) SetNillableIsActivated(b *bool) *ServiceUpdateOne {
	if b != nil {
		suo.SetIsActivated(*b)
	}
	return suo
}

// SetTagID sets the tag edge to Tag by id.
func (suo *ServiceUpdateOne) SetTagID(id int) *ServiceUpdateOne {
	if suo.tag == nil {
		suo.tag = make(map[int]struct{})
	}
	suo.tag[id] = struct{}{}
	return suo
}

// SetTag sets the tag edge to Tag.
func (suo *ServiceUpdateOne) SetTag(t *Tag) *ServiceUpdateOne {
	return suo.SetTagID(t.ID)
}

// AddEventIDs adds the events edge to Event by ids.
func (suo *ServiceUpdateOne) AddEventIDs(ids ...int) *ServiceUpdateOne {
	if suo.events == nil {
		suo.events = make(map[int]struct{})
	}
	for i := range ids {
		suo.events[ids[i]] = struct{}{}
	}
	return suo
}

// AddEvents adds the events edges to Event.
func (suo *ServiceUpdateOne) AddEvents(e ...*Event) *ServiceUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.AddEventIDs(ids...)
}

// ClearTag clears the tag edge to Tag.
func (suo *ServiceUpdateOne) ClearTag() *ServiceUpdateOne {
	suo.clearedTag = true
	return suo
}

// RemoveEventIDs removes the events edge to Event by ids.
func (suo *ServiceUpdateOne) RemoveEventIDs(ids ...int) *ServiceUpdateOne {
	if suo.removedEvents == nil {
		suo.removedEvents = make(map[int]struct{})
	}
	for i := range ids {
		suo.removedEvents[ids[i]] = struct{}{}
	}
	return suo
}

// RemoveEvents removes events edges to Event.
func (suo *ServiceUpdateOne) RemoveEvents(e ...*Event) *ServiceUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.RemoveEventIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *ServiceUpdateOne) Save(ctx context.Context) (*Service, error) {
	if suo.Name != nil {
		if err := service.NameValidator(*suo.Name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"Name\": %v", err)
		}
	}
	if suo.PubKey != nil {
		if err := service.PubKeyValidator(*suo.PubKey); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"PubKey\": %v", err)
		}
	}
	if len(suo.tag) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"tag\"")
	}
	if suo.clearedTag && suo.tag == nil {
		return nil, errors.New("ent: clearing a unique edge \"tag\"")
	}
	return suo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ServiceUpdateOne) SaveX(ctx context.Context) *Service {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *ServiceUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ServiceUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *ServiceUpdateOne) sqlSave(ctx context.Context) (s *Service, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   service.Table,
			Columns: service.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  suo.id,
				Type:   field.TypeInt,
				Column: service.FieldID,
			},
		},
	}
	if value := suo.Name; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: service.FieldName,
		})
	}
	if value := suo.PubKey; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: service.FieldPubKey,
		})
	}
	if value := suo.IsActivated; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: service.FieldIsActivated,
		})
	}
	if suo.clearedTag {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   service.TagTable,
			Columns: []string{service.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := suo.tag; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   service.TagTable,
			Columns: []string{service.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if nodes := suo.removedEvents; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   service.EventsTable,
			Columns: []string{service.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := suo.events; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   service.EventsTable,
			Columns: []string{service.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	s = &Service{config: suo.config}
	spec.Assign = s.assignValues
	spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
