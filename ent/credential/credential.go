// Code generated by entc, DO NOT EDIT.

package credential

import (
	"fmt"
)

const (
	// Label holds the string label denoting the credential type in the database.
	Label = "credential"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrincipal holds the string denoting the principal field in the database.
	FieldPrincipal = "principal"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// FieldKind holds the string denoting the kind field in the database.
	FieldKind = "kind"
	// FieldFails holds the string denoting the fails field in the database.
	FieldFails = "fails"
	// EdgeTarget holds the string denoting the target edge name in mutations.
	EdgeTarget = "target"
	// Table holds the table name of the credential in the database.
	Table = "credentials"
	// TargetTable is the table that holds the target relation/edge.
	TargetTable = "credentials"
	// TargetInverseTable is the table name for the Target entity.
	// It exists in this package in order to avoid circular dependency with the "target" package.
	TargetInverseTable = "targets"
	// TargetColumn is the table column denoting the target relation/edge.
	TargetColumn = "target_credentials"
)

// Columns holds all SQL columns for credential fields.
var Columns = []string{
	FieldID,
	FieldPrincipal,
	FieldSecret,
	FieldKind,
	FieldFails,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "credentials"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"target_credentials",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// SecretValidator is a validator for the "secret" field. It is called by the builders before save.
	SecretValidator func(string) error
	// DefaultFails holds the default value on creation for the "fails" field.
	DefaultFails int
	// FailsValidator is a validator for the "fails" field. It is called by the builders before save.
	FailsValidator func(int) error
)

// Kind defines the type for the "kind" enum field.
type Kind string

// Kind values.
const (
	KindPassword    Kind = "password"
	KindKey         Kind = "key"
	KindCertificate Kind = "certificate"
)

func (k Kind) String() string {
	return string(k)
}

// KindValidator is a validator for the "kind" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindPassword, KindKey, KindCertificate:
		return nil
	default:
		return fmt.Errorf("credential: invalid enum value for kind field: %q", k)
	}
}
