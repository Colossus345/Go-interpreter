package object

import (
	"fmt"
)

const (
	NULL_OBJ    = "NULL"
	BOOLEAN_OBJ = "BOOLEAN"

	INTEGER_OBJ = "INTEGER"
)

type ObjectType string
type Object interface {
	Type() ObjectType
	Inspect() string
}

type Boolean struct {
	Value bool
}
type Integer struct {
	Value int64
}

func (b *Integer) Type() ObjectType { return INTEGER_OBJ }
func (b *Integer) Inspect() string  { return fmt.Sprintf("%d", b.Value) }

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type NULL struct{}

func (n *NULL) Type() ObjectType { return NULL_OBJ }
func (n *NULL) Inspect() string  { return "null" }
