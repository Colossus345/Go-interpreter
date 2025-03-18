package object

import (
	"bytes"
	"fmt"
	"inter-median/internal/ast"
	"strings"
)

const (
	NULL_OBJ    = "NULL"
	BOOLEAN_OBJ = "BOOLEAN"

	FUNCTION_OBJ = "FUNCTION_OBJ"

	RETURN_VALUE_OBJ = "RETURN_VALUE_OBJ"
	INTEGER_OBJ      = "INTEGER"
	ERROR_OBJ        = "ERROR"
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
type Error struct {
	Message string
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}

	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

func (b *Integer) Type() ObjectType { return INTEGER_OBJ }
func (b *Integer) Inspect() string  { return fmt.Sprintf("%d", b.Value) }

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type NULL struct{}

func (n *NULL) Type() ObjectType { return NULL_OBJ }
func (n *NULL) Inspect() string  { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
