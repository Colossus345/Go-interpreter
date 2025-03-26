package export

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

func Exec(p Program, args ...string) (string, error) {
	arr := &object.Array{}
	for _, a := range args {
		arr.Elements = append(arr.Elements, StringToString(a))

	}
	env := object.NewEnvironment()
	env.Set("__args__", arr)
	obj := evaluator.Eval((*ast.Program)(p), env)
	if obj.Type() != object.STRING_OBJ {
		return "", fmt.Errorf("wrong return type")
	}
	return obj.Inspect(), nil
}

func StringToString(s string) *object.String {
	return &object.String{Value: s}
}
