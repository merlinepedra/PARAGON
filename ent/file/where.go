// Code generated by entc, DO NOT EDIT.

package file

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kcarretto/paragon/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
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
func IDGT(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// CreationTime applies equality check predicate on the "CreationTime" field. It's identical to CreationTimeEQ.
func CreationTime(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreationTime), v))
	})
}

// LastModifiedTime applies equality check predicate on the "LastModifiedTime" field. It's identical to LastModifiedTimeEQ.
func LastModifiedTime(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastModifiedTime), v))
	})
}

// Size applies equality check predicate on the "Size" field. It's identical to SizeEQ.
func Size(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// Content applies equality check predicate on the "Content" field. It's identical to ContentEQ.
func Content(v []byte) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// Hash applies equality check predicate on the "Hash" field. It's identical to HashEQ.
func Hash(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// ContentType applies equality check predicate on the "ContentType" field. It's identical to ContentTypeEQ.
func ContentType(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentType), v))
	})
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
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
func NameNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
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
func NameGT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// CreationTimeEQ applies the EQ predicate on the "CreationTime" field.
func CreationTimeEQ(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreationTime), v))
	})
}

// CreationTimeNEQ applies the NEQ predicate on the "CreationTime" field.
func CreationTimeNEQ(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreationTime), v))
	})
}

// CreationTimeIn applies the In predicate on the "CreationTime" field.
func CreationTimeIn(vs ...time.Time) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreationTime), v...))
	})
}

// CreationTimeNotIn applies the NotIn predicate on the "CreationTime" field.
func CreationTimeNotIn(vs ...time.Time) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreationTime), v...))
	})
}

// CreationTimeGT applies the GT predicate on the "CreationTime" field.
func CreationTimeGT(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreationTime), v))
	})
}

// CreationTimeGTE applies the GTE predicate on the "CreationTime" field.
func CreationTimeGTE(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreationTime), v))
	})
}

// CreationTimeLT applies the LT predicate on the "CreationTime" field.
func CreationTimeLT(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreationTime), v))
	})
}

// CreationTimeLTE applies the LTE predicate on the "CreationTime" field.
func CreationTimeLTE(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreationTime), v))
	})
}

// LastModifiedTimeEQ applies the EQ predicate on the "LastModifiedTime" field.
func LastModifiedTimeEQ(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastModifiedTime), v))
	})
}

// LastModifiedTimeNEQ applies the NEQ predicate on the "LastModifiedTime" field.
func LastModifiedTimeNEQ(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastModifiedTime), v))
	})
}

// LastModifiedTimeIn applies the In predicate on the "LastModifiedTime" field.
func LastModifiedTimeIn(vs ...time.Time) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastModifiedTime), v...))
	})
}

// LastModifiedTimeNotIn applies the NotIn predicate on the "LastModifiedTime" field.
func LastModifiedTimeNotIn(vs ...time.Time) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastModifiedTime), v...))
	})
}

// LastModifiedTimeGT applies the GT predicate on the "LastModifiedTime" field.
func LastModifiedTimeGT(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastModifiedTime), v))
	})
}

// LastModifiedTimeGTE applies the GTE predicate on the "LastModifiedTime" field.
func LastModifiedTimeGTE(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastModifiedTime), v))
	})
}

// LastModifiedTimeLT applies the LT predicate on the "LastModifiedTime" field.
func LastModifiedTimeLT(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastModifiedTime), v))
	})
}

// LastModifiedTimeLTE applies the LTE predicate on the "LastModifiedTime" field.
func LastModifiedTimeLTE(v time.Time) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastModifiedTime), v))
	})
}

// SizeEQ applies the EQ predicate on the "Size" field.
func SizeEQ(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSize), v))
	})
}

// SizeNEQ applies the NEQ predicate on the "Size" field.
func SizeNEQ(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSize), v))
	})
}

// SizeIn applies the In predicate on the "Size" field.
func SizeIn(vs ...int) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSize), v...))
	})
}

// SizeNotIn applies the NotIn predicate on the "Size" field.
func SizeNotIn(vs ...int) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSize), v...))
	})
}

// SizeGT applies the GT predicate on the "Size" field.
func SizeGT(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSize), v))
	})
}

// SizeGTE applies the GTE predicate on the "Size" field.
func SizeGTE(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSize), v))
	})
}

// SizeLT applies the LT predicate on the "Size" field.
func SizeLT(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSize), v))
	})
}

// SizeLTE applies the LTE predicate on the "Size" field.
func SizeLTE(v int) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSize), v))
	})
}

// ContentEQ applies the EQ predicate on the "Content" field.
func ContentEQ(v []byte) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "Content" field.
func ContentNEQ(v []byte) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "Content" field.
func ContentIn(vs ...[]byte) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "Content" field.
func ContentNotIn(vs ...[]byte) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "Content" field.
func ContentGT(v []byte) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "Content" field.
func ContentGTE(v []byte) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "Content" field.
func ContentLT(v []byte) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "Content" field.
func ContentLTE(v []byte) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// HashEQ applies the EQ predicate on the "Hash" field.
func HashEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// HashNEQ applies the NEQ predicate on the "Hash" field.
func HashNEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHash), v))
	})
}

// HashIn applies the In predicate on the "Hash" field.
func HashIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHash), v...))
	})
}

// HashNotIn applies the NotIn predicate on the "Hash" field.
func HashNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHash), v...))
	})
}

// HashGT applies the GT predicate on the "Hash" field.
func HashGT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHash), v))
	})
}

// HashGTE applies the GTE predicate on the "Hash" field.
func HashGTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHash), v))
	})
}

// HashLT applies the LT predicate on the "Hash" field.
func HashLT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHash), v))
	})
}

// HashLTE applies the LTE predicate on the "Hash" field.
func HashLTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHash), v))
	})
}

// HashContains applies the Contains predicate on the "Hash" field.
func HashContains(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHash), v))
	})
}

// HashHasPrefix applies the HasPrefix predicate on the "Hash" field.
func HashHasPrefix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHash), v))
	})
}

// HashHasSuffix applies the HasSuffix predicate on the "Hash" field.
func HashHasSuffix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHash), v))
	})
}

// HashEqualFold applies the EqualFold predicate on the "Hash" field.
func HashEqualFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHash), v))
	})
}

// HashContainsFold applies the ContainsFold predicate on the "Hash" field.
func HashContainsFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHash), v))
	})
}

// ContentTypeEQ applies the EQ predicate on the "ContentType" field.
func ContentTypeEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContentType), v))
	})
}

// ContentTypeNEQ applies the NEQ predicate on the "ContentType" field.
func ContentTypeNEQ(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContentType), v))
	})
}

// ContentTypeIn applies the In predicate on the "ContentType" field.
func ContentTypeIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContentType), v...))
	})
}

// ContentTypeNotIn applies the NotIn predicate on the "ContentType" field.
func ContentTypeNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.File(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContentType), v...))
	})
}

// ContentTypeGT applies the GT predicate on the "ContentType" field.
func ContentTypeGT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContentType), v))
	})
}

// ContentTypeGTE applies the GTE predicate on the "ContentType" field.
func ContentTypeGTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContentType), v))
	})
}

// ContentTypeLT applies the LT predicate on the "ContentType" field.
func ContentTypeLT(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContentType), v))
	})
}

// ContentTypeLTE applies the LTE predicate on the "ContentType" field.
func ContentTypeLTE(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContentType), v))
	})
}

// ContentTypeContains applies the Contains predicate on the "ContentType" field.
func ContentTypeContains(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContentType), v))
	})
}

// ContentTypeHasPrefix applies the HasPrefix predicate on the "ContentType" field.
func ContentTypeHasPrefix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContentType), v))
	})
}

// ContentTypeHasSuffix applies the HasSuffix predicate on the "ContentType" field.
func ContentTypeHasSuffix(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContentType), v))
	})
}

// ContentTypeEqualFold applies the EqualFold predicate on the "ContentType" field.
func ContentTypeEqualFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContentType), v))
	})
}

// ContentTypeContainsFold applies the ContainsFold predicate on the "ContentType" field.
func ContentTypeContainsFold(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContentType), v))
	})
}

// HasLinks applies the HasEdge predicate on the "links" edge.
func HasLinks() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LinksTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LinksTable, LinksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLinksWith applies the HasEdge predicate on the "links" edge with a given conditions (other predicates).
func HasLinksWith(preds ...predicate.Link) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(LinksInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, LinksTable, LinksColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
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
func Not(p predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		p(s.Not())
	})
}
