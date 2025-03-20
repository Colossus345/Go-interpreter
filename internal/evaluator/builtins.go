package evaluator

import (
	"fmt"
	"inter-median/internal/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}

			}
			return newError("argument to `len` not supported got %s",
				args[0].Type())
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) <= 1 {
				return newError("wrong number of arguments. got=%d want>1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("arguments to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElems := make([]object.Object, length+len(args)-1)
			copy(newElems, arr.Elements)
			for i := length; i < len(newElems); i++ {
				newElems[i] = args[i-length+1]
			}
			return &object.Array{Elements: newElems}

		},
	},
}
