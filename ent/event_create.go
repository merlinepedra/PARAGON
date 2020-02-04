// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/credential"
	"github.com/kcarretto/paragon/ent/event"
	"github.com/kcarretto/paragon/ent/file"
	"github.com/kcarretto/paragon/ent/job"
	"github.com/kcarretto/paragon/ent/link"
	"github.com/kcarretto/paragon/ent/service"
	"github.com/kcarretto/paragon/ent/tag"
	"github.com/kcarretto/paragon/ent/target"
	"github.com/kcarretto/paragon/ent/task"
	"github.com/kcarretto/paragon/ent/user"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	CreationTime *time.Time
	Kind         *event.Kind
	job          map[int]struct{}
	file         map[int]struct{}
	credential   map[int]struct{}
	link         map[int]struct{}
	tag          map[int]struct{}
	target       map[int]struct{}
	task         map[int]struct{}
	user         map[int]struct{}
	event        map[int]struct{}
	service      map[int]struct{}
	likers       map[int]struct{}
	owner        map[int]struct{}
	svcOwner     map[int]struct{}
}

// SetCreationTime sets the CreationTime field.
func (ec *EventCreate) SetCreationTime(t time.Time) *EventCreate {
	ec.CreationTime = &t
	return ec
}

// SetNillableCreationTime sets the CreationTime field if the given value is not nil.
func (ec *EventCreate) SetNillableCreationTime(t *time.Time) *EventCreate {
	if t != nil {
		ec.SetCreationTime(*t)
	}
	return ec
}

// SetKind sets the Kind field.
func (ec *EventCreate) SetKind(e event.Kind) *EventCreate {
	ec.Kind = &e
	return ec
}

// SetJobID sets the job edge to Job by id.
func (ec *EventCreate) SetJobID(id int) *EventCreate {
	if ec.job == nil {
		ec.job = make(map[int]struct{})
	}
	ec.job[id] = struct{}{}
	return ec
}

// SetNillableJobID sets the job edge to Job by id if the given value is not nil.
func (ec *EventCreate) SetNillableJobID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetJobID(*id)
	}
	return ec
}

// SetJob sets the job edge to Job.
func (ec *EventCreate) SetJob(j *Job) *EventCreate {
	return ec.SetJobID(j.ID)
}

// SetFileID sets the file edge to File by id.
func (ec *EventCreate) SetFileID(id int) *EventCreate {
	if ec.file == nil {
		ec.file = make(map[int]struct{})
	}
	ec.file[id] = struct{}{}
	return ec
}

// SetNillableFileID sets the file edge to File by id if the given value is not nil.
func (ec *EventCreate) SetNillableFileID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetFileID(*id)
	}
	return ec
}

// SetFile sets the file edge to File.
func (ec *EventCreate) SetFile(f *File) *EventCreate {
	return ec.SetFileID(f.ID)
}

// SetCredentialID sets the credential edge to Credential by id.
func (ec *EventCreate) SetCredentialID(id int) *EventCreate {
	if ec.credential == nil {
		ec.credential = make(map[int]struct{})
	}
	ec.credential[id] = struct{}{}
	return ec
}

// SetNillableCredentialID sets the credential edge to Credential by id if the given value is not nil.
func (ec *EventCreate) SetNillableCredentialID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetCredentialID(*id)
	}
	return ec
}

// SetCredential sets the credential edge to Credential.
func (ec *EventCreate) SetCredential(c *Credential) *EventCreate {
	return ec.SetCredentialID(c.ID)
}

// SetLinkID sets the link edge to Link by id.
func (ec *EventCreate) SetLinkID(id int) *EventCreate {
	if ec.link == nil {
		ec.link = make(map[int]struct{})
	}
	ec.link[id] = struct{}{}
	return ec
}

// SetNillableLinkID sets the link edge to Link by id if the given value is not nil.
func (ec *EventCreate) SetNillableLinkID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetLinkID(*id)
	}
	return ec
}

// SetLink sets the link edge to Link.
func (ec *EventCreate) SetLink(l *Link) *EventCreate {
	return ec.SetLinkID(l.ID)
}

// SetTagID sets the tag edge to Tag by id.
func (ec *EventCreate) SetTagID(id int) *EventCreate {
	if ec.tag == nil {
		ec.tag = make(map[int]struct{})
	}
	ec.tag[id] = struct{}{}
	return ec
}

// SetNillableTagID sets the tag edge to Tag by id if the given value is not nil.
func (ec *EventCreate) SetNillableTagID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetTagID(*id)
	}
	return ec
}

// SetTag sets the tag edge to Tag.
func (ec *EventCreate) SetTag(t *Tag) *EventCreate {
	return ec.SetTagID(t.ID)
}

// SetTargetID sets the target edge to Target by id.
func (ec *EventCreate) SetTargetID(id int) *EventCreate {
	if ec.target == nil {
		ec.target = make(map[int]struct{})
	}
	ec.target[id] = struct{}{}
	return ec
}

// SetNillableTargetID sets the target edge to Target by id if the given value is not nil.
func (ec *EventCreate) SetNillableTargetID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetTargetID(*id)
	}
	return ec
}

// SetTarget sets the target edge to Target.
func (ec *EventCreate) SetTarget(t *Target) *EventCreate {
	return ec.SetTargetID(t.ID)
}

// SetTaskID sets the task edge to Task by id.
func (ec *EventCreate) SetTaskID(id int) *EventCreate {
	if ec.task == nil {
		ec.task = make(map[int]struct{})
	}
	ec.task[id] = struct{}{}
	return ec
}

// SetNillableTaskID sets the task edge to Task by id if the given value is not nil.
func (ec *EventCreate) SetNillableTaskID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetTaskID(*id)
	}
	return ec
}

// SetTask sets the task edge to Task.
func (ec *EventCreate) SetTask(t *Task) *EventCreate {
	return ec.SetTaskID(t.ID)
}

// SetUserID sets the user edge to User by id.
func (ec *EventCreate) SetUserID(id int) *EventCreate {
	if ec.user == nil {
		ec.user = make(map[int]struct{})
	}
	ec.user[id] = struct{}{}
	return ec
}

// SetNillableUserID sets the user edge to User by id if the given value is not nil.
func (ec *EventCreate) SetNillableUserID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetUserID(*id)
	}
	return ec
}

// SetUser sets the user edge to User.
func (ec *EventCreate) SetUser(u *User) *EventCreate {
	return ec.SetUserID(u.ID)
}

// SetEventID sets the event edge to Event by id.
func (ec *EventCreate) SetEventID(id int) *EventCreate {
	if ec.event == nil {
		ec.event = make(map[int]struct{})
	}
	ec.event[id] = struct{}{}
	return ec
}

// SetNillableEventID sets the event edge to Event by id if the given value is not nil.
func (ec *EventCreate) SetNillableEventID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetEventID(*id)
	}
	return ec
}

// SetEvent sets the event edge to Event.
func (ec *EventCreate) SetEvent(e *Event) *EventCreate {
	return ec.SetEventID(e.ID)
}

// SetServiceID sets the service edge to Service by id.
func (ec *EventCreate) SetServiceID(id int) *EventCreate {
	if ec.service == nil {
		ec.service = make(map[int]struct{})
	}
	ec.service[id] = struct{}{}
	return ec
}

// SetNillableServiceID sets the service edge to Service by id if the given value is not nil.
func (ec *EventCreate) SetNillableServiceID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetServiceID(*id)
	}
	return ec
}

// SetService sets the service edge to Service.
func (ec *EventCreate) SetService(s *Service) *EventCreate {
	return ec.SetServiceID(s.ID)
}

// AddLikerIDs adds the likers edge to User by ids.
func (ec *EventCreate) AddLikerIDs(ids ...int) *EventCreate {
	if ec.likers == nil {
		ec.likers = make(map[int]struct{})
	}
	for i := range ids {
		ec.likers[ids[i]] = struct{}{}
	}
	return ec
}

// AddLikers adds the likers edges to User.
func (ec *EventCreate) AddLikers(u ...*User) *EventCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ec.AddLikerIDs(ids...)
}

// SetOwnerID sets the owner edge to User by id.
func (ec *EventCreate) SetOwnerID(id int) *EventCreate {
	if ec.owner == nil {
		ec.owner = make(map[int]struct{})
	}
	ec.owner[id] = struct{}{}
	return ec
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (ec *EventCreate) SetNillableOwnerID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetOwnerID(*id)
	}
	return ec
}

// SetOwner sets the owner edge to User.
func (ec *EventCreate) SetOwner(u *User) *EventCreate {
	return ec.SetOwnerID(u.ID)
}

// SetSvcOwnerID sets the svcOwner edge to Service by id.
func (ec *EventCreate) SetSvcOwnerID(id int) *EventCreate {
	if ec.svcOwner == nil {
		ec.svcOwner = make(map[int]struct{})
	}
	ec.svcOwner[id] = struct{}{}
	return ec
}

// SetNillableSvcOwnerID sets the svcOwner edge to Service by id if the given value is not nil.
func (ec *EventCreate) SetNillableSvcOwnerID(id *int) *EventCreate {
	if id != nil {
		ec = ec.SetSvcOwnerID(*id)
	}
	return ec
}

// SetSvcOwner sets the svcOwner edge to Service.
func (ec *EventCreate) SetSvcOwner(s *Service) *EventCreate {
	return ec.SetSvcOwnerID(s.ID)
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	if ec.CreationTime == nil {
		v := event.DefaultCreationTime()
		ec.CreationTime = &v
	}
	if ec.Kind == nil {
		return nil, errors.New("ent: missing required field \"Kind\"")
	}
	if err := event.KindValidator(*ec.Kind); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"Kind\": %v", err)
	}
	if len(ec.job) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"job\"")
	}
	if len(ec.file) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"file\"")
	}
	if len(ec.credential) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"credential\"")
	}
	if len(ec.link) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"link\"")
	}
	if len(ec.tag) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"tag\"")
	}
	if len(ec.target) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"target\"")
	}
	if len(ec.task) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"task\"")
	}
	if len(ec.user) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"user\"")
	}
	if len(ec.event) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"event\"")
	}
	if len(ec.service) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"service\"")
	}
	if len(ec.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if len(ec.svcOwner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"svcOwner\"")
	}
	return ec.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	var (
		e    = &Event{config: ec.config}
		spec = &sqlgraph.CreateSpec{
			Table: event.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: event.FieldID,
			},
		}
	)
	if value := ec.CreationTime; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: event.FieldCreationTime,
		})
		e.CreationTime = *value
	}
	if value := ec.Kind; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: event.FieldKind,
		})
		e.Kind = *value
	}
	if nodes := ec.job; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.JobTable,
			Columns: []string{event.JobColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.file; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.FileTable,
			Columns: []string{event.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.credential; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.CredentialTable,
			Columns: []string{event.CredentialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: credential.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.link; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.LinkTable,
			Columns: []string{event.LinkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: link.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.tag; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.TagTable,
			Columns: []string{event.TagColumn},
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
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.target; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.TargetTable,
			Columns: []string{event.TargetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: target.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.task; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.TaskTable,
			Columns: []string{event.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.user; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.UserTable,
			Columns: []string{event.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.event; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   event.EventTable,
			Columns: []string{event.EventColumn},
			Bidi:    true,
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
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.service; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   event.ServiceTable,
			Columns: []string{event.ServiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.likers; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   event.LikersTable,
			Columns: []string{event.LikersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.OwnerTable,
			Columns: []string{event.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := ec.svcOwner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.SvcOwnerTable,
			Columns: []string{event.SvcOwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: service.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, ec.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	e.ID = int(id)
	return e, nil
}
