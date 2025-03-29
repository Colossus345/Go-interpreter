package interpreter

import (
	"fmt"

	"github.com/Colossus345/Go-interpreter/internal/ast"
	"github.com/Colossus345/Go-interpreter/internal/evaluator"
	"github.com/Colossus345/Go-interpreter/internal/lexer"
	"github.com/Colossus345/Go-interpreter/internal/object"
	"github.com/Colossus345/Go-interpreter/internal/parser"
)

type Program *ast.Program

func Compile(s string) (Program, error) {
	l := lexer.New(s)
	if l == nil {
		return nil, fmt.Errorf("lexer nil")
	}
	p := parser.New(l)
	if p == nil {
		return nil, fmt.Errorf("parser nil")
	}

	program := p.ParseProgram()
	if program == nil {
		return nil, fmt.Errorf("program nil")
	}
	return program, nil
}

func Exec(p Program, args map[string]interface{}) (string, error) {
	arr := &object.Hash{Pairs: make(map[object.HashKey]object.HashPair)}
	for key, val := range args {
		strObj := StringToString(key)
		arr.Pairs[strObj.HashKey()] = object.HashPair{Key: strObj, Value: InterfaceToObject(val)}

	}
	env := object.NewEnvironment()
	for _, val := range arr.Pairs {
		env.Set(val.Key.Inspect(), val.Value)
	}
	obj := evaluator.Eval((*ast.Program)(p), env)
	if obj.Type() != object.STRING_OBJ {
		return "", fmt.Errorf("wrong return type")
	}
	return obj.Inspect(), nil
}

func StringToString(s string) *object.String {
	return &object.String{Value: s}
}
func InterfaceToObject(inter interface{}) object.Object {
	switch v := inter.(type) {
	case int:
		return &object.Integer{Value: int64(v)}
	case string:
		return &object.String{Value: v}
	case bool:
		return &object.Boolean{Value: v}
	default:
		return nil
	}
}
