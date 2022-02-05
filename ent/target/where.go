// Code generated by entc, DO NOT EDIT.

package target

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kcarretto/paragon/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// PrimaryIP applies equality check predicate on the "PrimaryIP" field. It's identical to PrimaryIPEQ.
func PrimaryIP(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimaryIP), v))
	})
}

// MachineUUID applies equality check predicate on the "MachineUUID" field. It's identical to MachineUUIDEQ.
func MachineUUID(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMachineUUID), v))
	})
}

// PublicIP applies equality check predicate on the "PublicIP" field. It's identical to PublicIPEQ.
func PublicIP(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublicIP), v))
	})
}

// PrimaryMAC applies equality check predicate on the "PrimaryMAC" field. It's identical to PrimaryMACEQ.
func PrimaryMAC(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimaryMAC), v))
	})
}

// Hostname applies equality check predicate on the "Hostname" field. It's identical to HostnameEQ.
func Hostname(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostname), v))
	})
}

// LastSeen applies equality check predicate on the "LastSeen" field. It's identical to LastSeenEQ.
func LastSeen(v time.Time) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSeen), v))
	})
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// OSEQ applies the EQ predicate on the "OS" field.
func OSEQ(v OS) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOS), v))
	})
}

// OSNEQ applies the NEQ predicate on the "OS" field.
func OSNEQ(v OS) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOS), v))
	})
}

// OSIn applies the In predicate on the "OS" field.
func OSIn(vs ...OS) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOS), v...))
	})
}

// OSNotIn applies the NotIn predicate on the "OS" field.
func OSNotIn(vs ...OS) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOS), v...))
	})
}

// PrimaryIPEQ applies the EQ predicate on the "PrimaryIP" field.
func PrimaryIPEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimaryIP), v))
	})
}

// PrimaryIPNEQ applies the NEQ predicate on the "PrimaryIP" field.
func PrimaryIPNEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrimaryIP), v))
	})
}

// PrimaryIPIn applies the In predicate on the "PrimaryIP" field.
func PrimaryIPIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrimaryIP), v...))
	})
}

// PrimaryIPNotIn applies the NotIn predicate on the "PrimaryIP" field.
func PrimaryIPNotIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrimaryIP), v...))
	})
}

// PrimaryIPGT applies the GT predicate on the "PrimaryIP" field.
func PrimaryIPGT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrimaryIP), v))
	})
}

// PrimaryIPGTE applies the GTE predicate on the "PrimaryIP" field.
func PrimaryIPGTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrimaryIP), v))
	})
}

// PrimaryIPLT applies the LT predicate on the "PrimaryIP" field.
func PrimaryIPLT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrimaryIP), v))
	})
}

// PrimaryIPLTE applies the LTE predicate on the "PrimaryIP" field.
func PrimaryIPLTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrimaryIP), v))
	})
}

// PrimaryIPContains applies the Contains predicate on the "PrimaryIP" field.
func PrimaryIPContains(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrimaryIP), v))
	})
}

// PrimaryIPHasPrefix applies the HasPrefix predicate on the "PrimaryIP" field.
func PrimaryIPHasPrefix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrimaryIP), v))
	})
}

// PrimaryIPHasSuffix applies the HasSuffix predicate on the "PrimaryIP" field.
func PrimaryIPHasSuffix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrimaryIP), v))
	})
}

// PrimaryIPEqualFold applies the EqualFold predicate on the "PrimaryIP" field.
func PrimaryIPEqualFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrimaryIP), v))
	})
}

// PrimaryIPContainsFold applies the ContainsFold predicate on the "PrimaryIP" field.
func PrimaryIPContainsFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrimaryIP), v))
	})
}

// MachineUUIDEQ applies the EQ predicate on the "MachineUUID" field.
func MachineUUIDEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMachineUUID), v))
	})
}

// MachineUUIDNEQ applies the NEQ predicate on the "MachineUUID" field.
func MachineUUIDNEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMachineUUID), v))
	})
}

// MachineUUIDIn applies the In predicate on the "MachineUUID" field.
func MachineUUIDIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMachineUUID), v...))
	})
}

// MachineUUIDNotIn applies the NotIn predicate on the "MachineUUID" field.
func MachineUUIDNotIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMachineUUID), v...))
	})
}

// MachineUUIDGT applies the GT predicate on the "MachineUUID" field.
func MachineUUIDGT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMachineUUID), v))
	})
}

// MachineUUIDGTE applies the GTE predicate on the "MachineUUID" field.
func MachineUUIDGTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMachineUUID), v))
	})
}

// MachineUUIDLT applies the LT predicate on the "MachineUUID" field.
func MachineUUIDLT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMachineUUID), v))
	})
}

// MachineUUIDLTE applies the LTE predicate on the "MachineUUID" field.
func MachineUUIDLTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMachineUUID), v))
	})
}

// MachineUUIDContains applies the Contains predicate on the "MachineUUID" field.
func MachineUUIDContains(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMachineUUID), v))
	})
}

// MachineUUIDHasPrefix applies the HasPrefix predicate on the "MachineUUID" field.
func MachineUUIDHasPrefix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMachineUUID), v))
	})
}

// MachineUUIDHasSuffix applies the HasSuffix predicate on the "MachineUUID" field.
func MachineUUIDHasSuffix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMachineUUID), v))
	})
}

// MachineUUIDIsNil applies the IsNil predicate on the "MachineUUID" field.
func MachineUUIDIsNil() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMachineUUID)))
	})
}

// MachineUUIDNotNil applies the NotNil predicate on the "MachineUUID" field.
func MachineUUIDNotNil() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMachineUUID)))
	})
}

// MachineUUIDEqualFold applies the EqualFold predicate on the "MachineUUID" field.
func MachineUUIDEqualFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMachineUUID), v))
	})
}

// MachineUUIDContainsFold applies the ContainsFold predicate on the "MachineUUID" field.
func MachineUUIDContainsFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMachineUUID), v))
	})
}

// PublicIPEQ applies the EQ predicate on the "PublicIP" field.
func PublicIPEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPublicIP), v))
	})
}

// PublicIPNEQ applies the NEQ predicate on the "PublicIP" field.
func PublicIPNEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPublicIP), v))
	})
}

// PublicIPIn applies the In predicate on the "PublicIP" field.
func PublicIPIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPublicIP), v...))
	})
}

// PublicIPNotIn applies the NotIn predicate on the "PublicIP" field.
func PublicIPNotIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPublicIP), v...))
	})
}

// PublicIPGT applies the GT predicate on the "PublicIP" field.
func PublicIPGT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPublicIP), v))
	})
}

// PublicIPGTE applies the GTE predicate on the "PublicIP" field.
func PublicIPGTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPublicIP), v))
	})
}

// PublicIPLT applies the LT predicate on the "PublicIP" field.
func PublicIPLT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPublicIP), v))
	})
}

// PublicIPLTE applies the LTE predicate on the "PublicIP" field.
func PublicIPLTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPublicIP), v))
	})
}

// PublicIPContains applies the Contains predicate on the "PublicIP" field.
func PublicIPContains(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPublicIP), v))
	})
}

// PublicIPHasPrefix applies the HasPrefix predicate on the "PublicIP" field.
func PublicIPHasPrefix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPublicIP), v))
	})
}

// PublicIPHasSuffix applies the HasSuffix predicate on the "PublicIP" field.
func PublicIPHasSuffix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPublicIP), v))
	})
}

// PublicIPIsNil applies the IsNil predicate on the "PublicIP" field.
func PublicIPIsNil() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPublicIP)))
	})
}

// PublicIPNotNil applies the NotNil predicate on the "PublicIP" field.
func PublicIPNotNil() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPublicIP)))
	})
}

// PublicIPEqualFold applies the EqualFold predicate on the "PublicIP" field.
func PublicIPEqualFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPublicIP), v))
	})
}

// PublicIPContainsFold applies the ContainsFold predicate on the "PublicIP" field.
func PublicIPContainsFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPublicIP), v))
	})
}

// PrimaryMACEQ applies the EQ predicate on the "PrimaryMAC" field.
func PrimaryMACEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrimaryMAC), v))
	})
}

// PrimaryMACNEQ applies the NEQ predicate on the "PrimaryMAC" field.
func PrimaryMACNEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrimaryMAC), v))
	})
}

// PrimaryMACIn applies the In predicate on the "PrimaryMAC" field.
func PrimaryMACIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrimaryMAC), v...))
	})
}

// PrimaryMACNotIn applies the NotIn predicate on the "PrimaryMAC" field.
func PrimaryMACNotIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrimaryMAC), v...))
	})
}

// PrimaryMACGT applies the GT predicate on the "PrimaryMAC" field.
func PrimaryMACGT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrimaryMAC), v))
	})
}

// PrimaryMACGTE applies the GTE predicate on the "PrimaryMAC" field.
func PrimaryMACGTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrimaryMAC), v))
	})
}

// PrimaryMACLT applies the LT predicate on the "PrimaryMAC" field.
func PrimaryMACLT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrimaryMAC), v))
	})
}

// PrimaryMACLTE applies the LTE predicate on the "PrimaryMAC" field.
func PrimaryMACLTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrimaryMAC), v))
	})
}

// PrimaryMACContains applies the Contains predicate on the "PrimaryMAC" field.
func PrimaryMACContains(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPrimaryMAC), v))
	})
}

// PrimaryMACHasPrefix applies the HasPrefix predicate on the "PrimaryMAC" field.
func PrimaryMACHasPrefix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPrimaryMAC), v))
	})
}

// PrimaryMACHasSuffix applies the HasSuffix predicate on the "PrimaryMAC" field.
func PrimaryMACHasSuffix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPrimaryMAC), v))
	})
}

// PrimaryMACIsNil applies the IsNil predicate on the "PrimaryMAC" field.
func PrimaryMACIsNil() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPrimaryMAC)))
	})
}

// PrimaryMACNotNil applies the NotNil predicate on the "PrimaryMAC" field.
func PrimaryMACNotNil() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPrimaryMAC)))
	})
}

// PrimaryMACEqualFold applies the EqualFold predicate on the "PrimaryMAC" field.
func PrimaryMACEqualFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPrimaryMAC), v))
	})
}

// PrimaryMACContainsFold applies the ContainsFold predicate on the "PrimaryMAC" field.
func PrimaryMACContainsFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPrimaryMAC), v))
	})
}

// HostnameEQ applies the EQ predicate on the "Hostname" field.
func HostnameEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHostname), v))
	})
}

// HostnameNEQ applies the NEQ predicate on the "Hostname" field.
func HostnameNEQ(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHostname), v))
	})
}

// HostnameIn applies the In predicate on the "Hostname" field.
func HostnameIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHostname), v...))
	})
}

// HostnameNotIn applies the NotIn predicate on the "Hostname" field.
func HostnameNotIn(vs ...string) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHostname), v...))
	})
}

// HostnameGT applies the GT predicate on the "Hostname" field.
func HostnameGT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHostname), v))
	})
}

// HostnameGTE applies the GTE predicate on the "Hostname" field.
func HostnameGTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHostname), v))
	})
}

// HostnameLT applies the LT predicate on the "Hostname" field.
func HostnameLT(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHostname), v))
	})
}

// HostnameLTE applies the LTE predicate on the "Hostname" field.
func HostnameLTE(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHostname), v))
	})
}

// HostnameContains applies the Contains predicate on the "Hostname" field.
func HostnameContains(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHostname), v))
	})
}

// HostnameHasPrefix applies the HasPrefix predicate on the "Hostname" field.
func HostnameHasPrefix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHostname), v))
	})
}

// HostnameHasSuffix applies the HasSuffix predicate on the "Hostname" field.
func HostnameHasSuffix(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHostname), v))
	})
}

// HostnameIsNil applies the IsNil predicate on the "Hostname" field.
func HostnameIsNil() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldHostname)))
	})
}

// HostnameNotNil applies the NotNil predicate on the "Hostname" field.
func HostnameNotNil() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldHostname)))
	})
}

// HostnameEqualFold applies the EqualFold predicate on the "Hostname" field.
func HostnameEqualFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHostname), v))
	})
}

// HostnameContainsFold applies the ContainsFold predicate on the "Hostname" field.
func HostnameContainsFold(v string) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHostname), v))
	})
}

// LastSeenEQ applies the EQ predicate on the "LastSeen" field.
func LastSeenEQ(v time.Time) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastSeen), v))
	})
}

// LastSeenNEQ applies the NEQ predicate on the "LastSeen" field.
func LastSeenNEQ(v time.Time) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastSeen), v))
	})
}

// LastSeenIn applies the In predicate on the "LastSeen" field.
func LastSeenIn(vs ...time.Time) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastSeen), v...))
	})
}

// LastSeenNotIn applies the NotIn predicate on the "LastSeen" field.
func LastSeenNotIn(vs ...time.Time) predicate.Target {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Target(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastSeen), v...))
	})
}

// LastSeenGT applies the GT predicate on the "LastSeen" field.
func LastSeenGT(v time.Time) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastSeen), v))
	})
}

// LastSeenGTE applies the GTE predicate on the "LastSeen" field.
func LastSeenGTE(v time.Time) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastSeen), v))
	})
}

// LastSeenLT applies the LT predicate on the "LastSeen" field.
func LastSeenLT(v time.Time) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastSeen), v))
	})
}

// LastSeenLTE applies the LTE predicate on the "LastSeen" field.
func LastSeenLTE(v time.Time) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastSeen), v))
	})
}

// LastSeenIsNil applies the IsNil predicate on the "LastSeen" field.
func LastSeenIsNil() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastSeen)))
	})
}

// LastSeenNotNil applies the NotNil predicate on the "LastSeen" field.
func LastSeenNotNil() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastSeen)))
	})
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TasksTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TasksInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCredentials applies the HasEdge predicate on the "credentials" edge.
func HasCredentials() predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CredentialsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CredentialsTable, CredentialsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCredentialsWith applies the HasEdge predicate on the "credentials" edge with a given conditions (other predicates).
func HasCredentialsWith(preds ...predicate.Credential) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CredentialsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CredentialsTable, CredentialsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Target) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Target) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Target) predicate.Target {
	return predicate.Target(func(s *sql.Selector) {
		p(s.Not())
	})
}
