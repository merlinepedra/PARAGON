// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/service"
	"github.com/kcarretto/paragon/ent/tag"
)

// Service is the model entity for the Service schema.
type Service struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	// The name displayed for the service
	Name string `json:"Name,omitempty"`
	// PubKey holds the value of the "PubKey" field.
	// The ed25519 public key for the service (stored in Base64 of DER format)
	PubKey string `json:"PubKey,omitempty"`
	// Config holds the value of the "Config" field.
	// The configuration script of the service (usually a Renegade Script)
	Config string `json:"Config,omitempty"`
	// IsActivated holds the value of the "IsActivated" field.
	// True iff the service is active and able to authenticate
	IsActivated bool `json:"IsActivated,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ServiceQuery when eager-loading is set.
	Edges       ServiceEdges `json:"edges"`
	service_tag *int
}

// ServiceEdges holds the relations/edges for other nodes in the graph.
type ServiceEdges struct {
	// Tag holds the value of the tag edge.
	Tag *Tag `json:"tag,omitempty"`
	// Events holds the value of the events edge.
	Events []*Event `json:"events,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ServiceEdges) TagOrErr() (*Tag, error) {
	if e.loadedTypes[0] {
		if e.Tag == nil {
			// The edge tag was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: tag.Label}
		}
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e ServiceEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[1] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Service) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case service.FieldIsActivated:
			values[i] = new(sql.NullBool)
		case service.FieldID:
			values[i] = new(sql.NullInt64)
		case service.FieldName, service.FieldPubKey, service.FieldConfig:
			values[i] = new(sql.NullString)
		case service.ForeignKeys[0]: // service_tag
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Service", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Service fields.
func (s *Service) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case service.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case service.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case service.FieldPubKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PubKey", values[i])
			} else if value.Valid {
				s.PubKey = value.String
			}
		case service.FieldConfig:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Config", values[i])
			} else if value.Valid {
				s.Config = value.String
			}
		case service.FieldIsActivated:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field IsActivated", values[i])
			} else if value.Valid {
				s.IsActivated = value.Bool
			}
		case service.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field service_tag", value)
			} else if value.Valid {
				s.service_tag = new(int)
				*s.service_tag = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTag queries the "tag" edge of the Service entity.
func (s *Service) QueryTag() *TagQuery {
	return (&ServiceClient{config: s.config}).QueryTag(s)
}

// QueryEvents queries the "events" edge of the Service entity.
func (s *Service) QueryEvents() *EventQuery {
	return (&ServiceClient{config: s.config}).QueryEvents(s)
}

// Update returns a builder for updating this Service.
// Note that you need to call Service.Unwrap() before calling this method if this Service
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Service) Update() *ServiceUpdateOne {
	return (&ServiceClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Service entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Service) Unwrap() *Service {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Service is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Service) String() string {
	var builder strings.Builder
	builder.WriteString("Service(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", Name=")
	builder.WriteString(s.Name)
	builder.WriteString(", PubKey=")
	builder.WriteString(s.PubKey)
	builder.WriteString(", Config=")
	builder.WriteString(s.Config)
	builder.WriteString(", IsActivated=")
	builder.WriteString(fmt.Sprintf("%v", s.IsActivated))
	builder.WriteByte(')')
	return builder.String()
}

// Services is a parsable slice of Service.
type Services []*Service

func (s Services) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
