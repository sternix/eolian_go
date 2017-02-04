package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type Expression struct {
	obj *C.Eolian_Expression
}

func NewExpression(obj *C.Eolian_Expression) *Expression {
	return &Expression{
		obj: obj,
	}
}
func (p *Expression) IsValid() bool {
	return p.obj != nil
}

func (p *Expression) Eval(m ExpressionMask) *Value {
	return NewValue(C.eolian_expression_eval(p.obj, C.Eolian_Expression_Mask(m)))
}

func (p *Expression) EvalType(t *Type) *Value {
	return NewValue(C.eolian_expression_eval_type(p.obj, t.obj))
}

func (p *Expression) Serialize() string {
	return GoString(C.eolian_expression_serialize(p.obj), true)
}

func (p *Expression) Type() ExpressionType {
	return ExpressionType(C.eolian_expression_type_get(p.obj))
}

func (p *Expression) BinaryOperator() BinaryOperator {
	return BinaryOperator(C.eolian_expression_binary_operator_get(p.obj))
}

func (p *Expression) BinaryLhs() *Expression {
	return NewExpression(C.eolian_expression_binary_lhs_get(p.obj))
}

func (p *Expression) BinaryRhs() *Expression {
	return NewExpression(C.eolian_expression_binary_rhs_get(p.obj))
}

func (p *Expression) UnaryOperator() UnaryOperator {
	return UnaryOperator(C.eolian_expression_unary_operator_get(p.obj))
}

func (p *Expression) UnaryExpression() *Expression {
	return NewExpression(C.eolian_expression_unary_expression_get(p.obj))
}

func (p *Expression) Value() *Value {
	return NewValue(C.eolian_expression_value_get(p.obj))
}
