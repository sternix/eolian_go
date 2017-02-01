package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type Implement struct {
	obj *C.Eolian_Implement
}

func NewImplement(obj *C.Eolian_Implement) *Implement {
	return &Implement{
		obj: obj,
	}
}

func (p *Implement) IsValid() bool {
	return p.obj != nil
}

func (p *Implement) FullName() string {
	return GoString(C.eolian_implement_full_name_get(p.obj))
}

func (p *Implement) Class() *Class {
	return NewClass(C.eolian_implement_class_get(p.obj))
}

func (p *Implement) Function() (*Function, FunctionType) {
	var t C.Eolian_Function_Type
	fobj := C.eolian_implement_function_get(p.obj, &t)
	return NewFunction(fobj), FunctionType(t)
}

func (p *Implement) IsAuto() bool {
	return C.eolian_implement_is_auto(p.obj) == ETrue
}

func (p *Implement) IsEmpty() bool {
	return C.eolian_implement_is_empty(p.obj) == ETrue
}

func (p *Implement) IsVirtual() bool {
	return C.eolian_implement_is_virtual(p.obj) == ETrue
}

func (p *Implement) IsGetProperty() bool {
	return C.eolian_implement_is_prop_get(p.obj) == ETrue
}

func (p *Implement) IsSetProperty() bool {
	return C.eolian_implement_is_prop_set(p.obj) == ETrue
}
