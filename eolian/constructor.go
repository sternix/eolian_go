package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type Constructor struct {
	obj *C.Eolian_Constructor
}

func NewConstructor(obj *C.Eolian_Constructor) *Constructor {
	return &Constructor{
		obj: obj,
	}
}

func (p *Constructor) IsValid() bool {
	return p.obj != nil
}

func (p *Constructor) FullName() string {
	return GoString(C.eolian_constructor_full_name_get(p.obj))
}

func (p *Constructor) Class() *Class {
	return NewClass(C.eolian_constructor_class_get(p.obj))
}

func (p *Constructor) Function() *Function {
	return NewFunction(C.eolian_constructor_function_get(p.obj))
}

func (p *Constructor) IsOptional() bool {
	return C.eolian_constructor_is_optional(p.obj) == ETrue
}
