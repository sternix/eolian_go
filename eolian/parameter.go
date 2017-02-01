package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type Parameter struct {
	obj *C.Eolian_Function_Parameter
}

func NewParameter(obj *C.Eolian_Function_Parameter) *Parameter {
	return &Parameter{
		obj: obj,
	}
}

func (p *Parameter) IsValid() bool {
	return p.obj != nil
}

func (p *Parameter) Direction() ParameterDir {
	return ParameterDir(C.eolian_parameter_direction_get(p.obj))
}

func (p *Parameter) Type() *Type {
	return NewType(C.eolian_parameter_type_get(p.obj))
}

func (p *Parameter) DefaultValue() *Expression {
	return NewExpression(C.eolian_parameter_default_value_get(p.obj))
}

func (p *Parameter) Name() string {
	return GoString(C.eolian_parameter_name_get(p.obj))
}

func (p *Parameter) IsNoNull() bool {
	return C.eolian_parameter_is_nonull(p.obj) == ETrue
}

func (p *Parameter) IsNullable() bool {
	return C.eolian_parameter_is_nullable(p.obj) == ETrue
}

func (p *Parameter) IsOptional() bool {
	return C.eolian_parameter_is_optional(p.obj) == ETrue
}

func (p *Parameter) Documentation() *Documentation {
	return NewDocumentation(C.eolian_parameter_documentation_get(p.obj))
}
