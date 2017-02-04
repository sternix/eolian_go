package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type Value struct {
	obj C.Eolian_Value
}

func NewValue(obj C.Eolian_Value) *Value {
	return &Value{
		obj: obj,
	}
}

func (p *Value) ToLiteral() string {
	return GoString(C.eolian_expression_value_to_literal(&p.obj), true)
}
