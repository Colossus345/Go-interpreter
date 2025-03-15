package evaluator

import (
	"inter-median/internal/ast"
	"inter-median/internal/object"
)

var (
	NULL  = &object.NULL{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func Eval(node ast.Node) object.Object {

	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		right := Eval(node.Right)
		left := Eval(node.Left)
		return evalInfixExpression(node.Operator, left, right)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	}

	return nil

}

func evalInfixExpression(oper string, left, right object.Object) object.Object {

	switch {
	case left.Type() == right.Type() && left.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(oper, left, right)
	case oper == "==":
		return nativeBoolToBooleanObject(left == right)
	case oper == "!=":
		return nativeBoolToBooleanObject(left != right)
	}

	return NULL
}
func evalIntegerInfixExpression(oper string,
	left,
	right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value
	switch oper {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	}
	return NULL
}

func evalPrefixExpression(oper string, right object.Object) object.Object {

	switch oper {
	case "!":
		return evalBangOperator(right)
	case "-":
		return evalMinusPrefixOperator(right)

	}
	return NULL
}

func evalBangOperator(right object.Object) object.Object {
	switch right {
	case FALSE:
		return TRUE
	case NULL:
		return TRUE

	case TRUE:
		return FALSE

	default:
		return FALSE
	}
}
func evalMinusPrefixOperator(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return NULL
	}
	value := right.(*object.Integer).Value

	return &object.Integer{Value: -value}
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
	}
	return result
}
