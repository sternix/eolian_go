package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"
import "unsafe"

type Typedecl struct {
	obj *C.Eolian_Typedecl
}

func NewTypedecl(obj *C.Eolian_Typedecl) *Typedecl {
	return &Typedecl{
		obj: obj,
	}
}

func NewTypedeclAliasByName(name string) *Typedecl {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewTypedecl(C.eolian_typedecl_alias_get_by_name(cname))
}

func NewTypedeclStructByName(name string) *Typedecl {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewTypedecl(C.eolian_typedecl_struct_get_by_name(cname))
}

func NewTypedeclEnumByName(name string) *Typedecl {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewTypedecl(C.eolian_typedecl_enum_get_by_name(cname))
}

func TypedeclAliasesByFile(name string) []*Typedecl {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewIterator(C.eolian_typedecl_aliases_get_by_file(cname)).TypedeclSlice()
}

func TypedeclStructsByFile(name string) []*Typedecl {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewIterator(C.eolian_typedecl_structs_get_by_file(cname)).TypedeclSlice()
}

func TypedeclEnumsByFile(name string) []*Typedecl {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewIterator(C.eolian_typedecl_enums_get_by_file(cname)).TypedeclSlice()
}

func TypedeclAllAliases() []*Typedecl {
	return NewIterator(C.eolian_typedecl_all_aliases_get()).TypedeclSlice()
}

func TypedeclAllStructs() []*Typedecl {
	return NewIterator(C.eolian_typedecl_all_structs_get()).TypedeclSlice()
}

func TypedeclAllEnums() []*Typedecl {
	return NewIterator(C.eolian_typedecl_all_enums_get()).TypedeclSlice()
}

func (p *Typedecl) IsValid() bool {
	return p.obj != nil
}

func (p *Typedecl) Type() TypedeclType {
	return TypedeclType(C.eolian_typedecl_type_get(p.obj))
}

func (p *Typedecl) StructFields() []*StructField {
	return NewIterator(C.eolian_typedecl_struct_fields_get(p.obj)).StructFieldSlice()
}

func (p *Typedecl) StructFieldByName(name string) *StructField {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewStructField(C.eolian_typedecl_struct_field_get(p.obj, cname))
}

func (p *Typedecl) EnumFields() []*EnumField {
	return NewIterator(C.eolian_typedecl_enum_fields_get(p.obj)).EnumFieldSlice()
}

func (p *Typedecl) EnumFieldByName(name string) *EnumField {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewEnumField(C.eolian_typedecl_enum_field_get(p.obj, cname))
}

func (p *Typedecl) EnumLegacyPrefix() string {
	return GoString(C.eolian_typedecl_enum_legacy_prefix_get(p.obj))
}

func (p *Typedecl) Documentation() *Documentation {
	return NewDocumentation(C.eolian_typedecl_documentation_get(p.obj))
}

func (p *Typedecl) File() string {
	return GoString(C.eolian_typedecl_file_get(p.obj))
}

func (p *Typedecl) BaseType() *Type {
	return NewType(C.eolian_typedecl_base_type_get(p.obj))
}

func (p *Typedecl) AliasedBaseType() *Type {
	return NewType(C.eolian_typedecl_aliased_base_get(p.obj))
}

func (p *Typedecl) IsExtern() bool {
	return C.eolian_typedecl_is_extern(p.obj) == ETrue
}

func (p *Typedecl) CType() string {
	return GoString(C.eolian_typedecl_c_type_get(p.obj))
}

func (p *Typedecl) Name() string {
	return GoString(C.eolian_typedecl_name_get(p.obj))
}

func (p *Typedecl) FullName() string {
	return GoString(C.eolian_typedecl_full_name_get(p.obj))
}

func (p *Typedecl) Namespaces() []string {
	return NewIterator(C.eolian_typedecl_namespaces_get(p.obj)).StringSlice()
}

func (p *Typedecl) FreeFunc() string {
	return GoString(C.eolian_typedecl_free_func_get(p.obj))
}
