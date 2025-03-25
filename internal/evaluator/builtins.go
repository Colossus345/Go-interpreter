package evaluator

import (
	"fmt"
	"github.com/Colossus345/Go-interpreter/internal/object"
	"os"
)

var builtins = map[string]*object.Builtin{
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL

		},
	},
	"fputs": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 {
				return newError("wrong number of arguments. got=%d, want at least 2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first parameter must be PATH string. got=%s", args[0].Type())
			}
			file, err := os.OpenFile(args[0].Inspect(),
				os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			defer file.Close()
			if err != nil {
				return newError("Open file error:%s", err)
			}
			prnt := []interface{}{}
			for i := 1; i < len(args); i++ {
				prnt = append(prnt, args[i].Inspect())
			}

			fmt.Fprintln(file, prnt...)
			return NULL
		},
	},
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
