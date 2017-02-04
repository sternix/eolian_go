package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type StructField struct {
	obj *C.Eolian_Struct_Type_Field
}

func NewStructField(obj *C.Eolian_Struct_Type_Field) *StructField {
	return &StructField{
		obj: obj,
	}
}

func (p *StructField) IsValid() bool {
	return p.obj != nil
}

func (p *StructField) Name() string {
	return GoString(C.eolian_typedecl_struct_field_name_get(p.obj))
}

func (p *StructField) Documentation() *Documentation {
	return NewDocumentation(C.eolian_typedecl_struct_field_documentation_get(p.obj))
}

func (p *StructField) Type() *Type {
	return NewType(C.eolian_typedecl_struct_field_type_get(p.obj))
}

// EnumField

type EnumField struct {
	obj *C.Eolian_Enum_Type_Field
}

func NewEnumField(obj *C.Eolian_Enum_Type_Field) *EnumField {
	return &EnumField{
		obj: obj,
	}
}

func (p *EnumField) IsValid() bool {
	return p.obj != nil
}

func (p *EnumField) Name() string {
	return GoString(C.eolian_typedecl_enum_field_name_get(p.obj))
}

func (p *EnumField) CName() string {
	return GoString(C.eolian_typedecl_enum_field_c_name_get(p.obj), true)
}

func (p *EnumField) Documentation() *Documentation {
	return NewDocumentation(C.eolian_typedecl_enum_field_documentation_get(p.obj))
}

func (p *EnumField) Value(force bool) *Expression {
	return NewExpression(C.eolian_typedecl_enum_field_value_get(p.obj, EBool(force)))
}
