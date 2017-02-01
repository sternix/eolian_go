package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type Function struct {
	obj *C.Eolian_Function
}

func NewFunction(obj *C.Eolian_Function) *Function {
	return &Function{
		obj: obj,
	}
}

func (p *Function) IsValid() bool {
	return p != nil
}

func (p *Function) Type() FunctionType {
	return FunctionType(C.eolian_function_type_get(p.obj))
}

func (p *Function) Scope(t FunctionType) ObjectScope {
	return ObjectScope(C.eolian_function_scope_get(p.obj, C.Eolian_Function_Type(t)))
}

func (p *Function) Name() string {
	return GoString(C.eolian_function_name_get(p.obj))
}

func (p *Function) FullCName(t FunctionType, legacy bool) string {
	return GoString(C.eolian_function_full_c_name_get(p.obj, C.Eolian_Function_Type(t), EBool(legacy)))
}

func (p *Function) Legacy(t FunctionType) string {
	return GoString(C.eolian_function_legacy_get(p.obj, C.Eolian_Function_Type(t)))
}

func (p *Function) IsVirtualPure(t FunctionType) bool {
	return C.eolian_function_is_virtual_pure(p.obj, C.Eolian_Function_Type(t)) == ETrue
}

func (p *Function) IsAuto(t FunctionType) bool {
	return C.eolian_function_is_auto(p.obj, C.Eolian_Function_Type(t)) == ETrue
}

func (p *Function) IsEmpty(t FunctionType) bool {
	return C.eolian_function_is_empty(p.obj, C.Eolian_Function_Type(t)) == ETrue
}

func (p *Function) IsLegacyOnly(t FunctionType) bool {
	return C.eolian_function_is_legacy_only(p.obj, C.Eolian_Function_Type(t)) == ETrue
}

func (p *Function) IsClass() bool {
	return C.eolian_function_is_class(p.obj) == ETrue
}

func (p *Function) IsCOnly() bool {
	return C.eolian_function_is_c_only(p.obj) == ETrue
}

func (p *Function) IsBeta() bool {
	return C.eolian_function_is_beta(p.obj) == ETrue
}

func (p *Function) IsConstructor(c Class) bool {
	return C.eolian_function_is_constructor(p.obj, c.obj) == ETrue
}

func (p *Function) IsConst() bool {
	return C.eolian_function_object_is_const(p.obj) == ETrue
}

func (p *Function) Class() *Class {
	return NewClass(C.eolian_function_class_get(p.obj))
}

func (p *Function) Parameters() []*Parameter {
	return NewIterator(C.eolian_function_parameters_get(p.obj)).ParameterSlice()
}

func (p *Function) Keys(t FunctionType) []*Parameter {
	return NewIterator(C.eolian_property_keys_get(p.obj, C.Eolian_Function_Type(t))).ParameterSlice()
}

func (p *Function) Values(t FunctionType) []*Parameter {
	return NewIterator(C.eolian_property_values_get(p.obj, C.Eolian_Function_Type(t))).ParameterSlice()
}

type functionReturn struct {
	obj *C.Eolian_Function
}

func (p *Function) Return() *functionReturn {
	return &functionReturn{
		obj: p.obj,
	}
}

func (p *functionReturn) Type(t FunctionType) *Type {
	return NewType(C.eolian_function_return_type_get(p.obj, C.Eolian_Function_Type(t)))
}

func (p *functionReturn) DefaultValue(t FunctionType) *Expression {
	return NewExpression(C.eolian_function_return_default_value_get(p.obj, C.Eolian_Function_Type(t)))
}

func (p *functionReturn) Documentation(t FunctionType) *Documentation {
	return NewDocumentation(C.eolian_function_return_documentation_get(p.obj, C.Eolian_Function_Type(t)))
}

func (p *functionReturn) IsWarnUnused(t FunctionType) bool {
	return C.eolian_function_return_is_warn_unused(p.obj, C.Eolian_Function_Type(t)) == ETrue
}

func (p *Function) IsImplemented(t FunctionType, c *Class) bool {
	return C.eolian_function_is_implemented(p.obj, C.Eolian_Function_Type(t), c.obj) == ETrue
}
