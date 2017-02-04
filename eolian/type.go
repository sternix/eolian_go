package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type Type struct {
	obj        *C.Eolian_Type
	namespaces []string
}

func NewType(obj *C.Eolian_Type) *Type {
	return &Type{
		obj: obj,
	}
}

func (p *Type) IsValid() bool {
	return p.obj != nil
}

func (p *Type) Type() EoType {
	return EoType(C.eolian_type_type_get(p.obj))
}

func (p *Type) File() string {
	return GoString(C.eolian_type_file_get(p.obj))
}

func (p *Type) BaseType() *Type {
	return NewType(C.eolian_type_base_type_get(p.obj))
}

func (p *Type) NextType() *Type {
	return NewType(C.eolian_type_next_type_get(p.obj))
}

func (p *Type) Typedecl() *Typedecl {
	return NewTypedecl(C.eolian_type_typedecl_get(p.obj))
}

func (p *Type) AliasedBase() *Type {
	return NewType(C.eolian_type_aliased_base_get(p.obj))
}

func (p *Type) Class() *Class {
	return NewClass(C.eolian_type_class_get(p.obj))
}

func (p *Type) ArraySize() int {
	return int(C.eolian_type_array_size_get(p.obj))
}

func (p *Type) IsOwn() bool {
	return C.eolian_type_is_own(p.obj) == ETrue
}

func (p *Type) IsConst() bool {
	return C.eolian_type_is_const(p.obj) == ETrue
}

func (p *Type) IsRef() bool {
	return C.eolian_type_is_ref(p.obj) == ETrue
}

func (p *Type) CType() string {
	return GoString(C.eolian_type_c_type_get(p.obj), true)
}

func (p *Type) Name() string {
	return GoString(C.eolian_type_name_get(p.obj))
}

func (p *Type) FullName() string {
	return GoString(C.eolian_type_full_name_get(p.obj))
}

func (p *Type) Namespaces() []string {
	if p.namespaces != nil {
		return p.namespaces
	}
	p.namespaces = NewIterator(C.eolian_type_namespaces_get(p.obj)).StringSlice()
	return p.namespaces
}

func (p *Type) FreeFunc() string {
	return GoString(C.eolian_type_free_func_get(p.obj))
}
