package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"
import "unsafe"

const (
	ETrue  = C.Eina_Bool(1)
	EFalse = C.Eina_Bool(0)
)

func EBool(b bool) C.Eina_Bool {
	if b {
		return ETrue
	}
	return EFalse
}

func GoBool(b C.Eina_Bool) bool {
	if b == ETrue {
		return true
	}
	return false
}

func GoString(str *C.Eina_Stringshare) string {
	//defer C.eina_stringshare_del(str)
	return C.GoString((*C.char)(str))
}

type Iterator struct {
	itr  *C.Eina_Iterator
	data unsafe.Pointer
}

func NewIterator(itr *C.Eina_Iterator) *Iterator {
	return &Iterator{
		itr: itr,
	}
}

func (p *Iterator) Next() bool {
	if p.itr != nil {
		if C.eina_iterator_next(p.itr, &p.data) == ETrue {
			return true
		} else {
			C.eina_iterator_free(p.itr)
			p.itr = nil
			return false
		}
	}
	return false
}

func (p *Iterator) StringSlice() (ret []string) {
	for p.Next() {
		//ret = append(ret, GoString((*C.Eina_Stringshare)(p.data)))
		ret = append(ret, C.GoString((*C.char)(p.data)))
	}
	return
}

func (p *Iterator) ClassSlice() (ret []*Class) {
	for p.Next() {
		ret = append(ret, NewClass((*C.Eolian_Class)(p.data)))
	}
	return
}

func (p *Iterator) FunctionSlice() (ret []*Function) {
	for p.Next() {
		ret = append(ret, NewFunction((*C.Eolian_Function)(p.data)))
	}
	return
}

func (p *Iterator) TypeSlice() (ret []*Type) {
	for p.Next() {
		ret = append(ret, NewType((*C.Eolian_Type)(p.data)))
	}
	return
}

func (p *Iterator) TypedeclSlice() (ret []*Typedecl) {
	for p.Next() {
		ret = append(ret, NewTypedecl((*C.Eolian_Typedecl)(p.data)))
	}
	return
}

func (p *Iterator) ParameterSlice() (ret []*Parameter) {
	for p.Next() {
		ret = append(ret, NewParameter((*C.Eolian_Function_Parameter)(p.data)))
	}
	return
}

func (p *Iterator) ImplementSlice() (ret []*Implement) {
	for p.Next() {
		ret = append(ret, NewImplement((*C.Eolian_Implement)(p.data)))
	}
	return
}

func (p *Iterator) ConstructorSlice() (ret []*Constructor) {
	for p.Next() {
		ret = append(ret, NewConstructor((*C.Eolian_Constructor)(p.data)))
	}
	return
}

func (p *Iterator) EventSlice() (ret []*Event) {
	for p.Next() {
		ret = append(ret, NewEvent((*C.Eolian_Event)(p.data)))
	}
	return
}

func (p *Iterator) ExpressionSlice() (ret []*Expression) {
	for p.Next() {
		ret = append(ret, NewExpression((*C.Eolian_Expression)(p.data)))
	}
	return
}

func (p *Iterator) VariableSlice() (ret []*Variable) {
	for p.Next() {
		ret = append(ret, NewVariable((*C.Eolian_Variable)(p.data)))
	}
	return
}

func (p *Iterator) StructFieldSlice() (ret []*StructField) {
	for p.Next() {
		ret = append(ret, NewStructField((*C.Eolian_Struct_Type_Field)(p.data)))
	}
	return
}

func (p *Iterator) EnumFieldSlice() (ret []*EnumField) {
	for p.Next() {
		ret = append(ret, NewEnumField((*C.Eolian_Enum_Type_Field)(p.data)))
	}
	return
}

func (p *Iterator) DeclarationSlice() (ret []*Declaration) {
	for p.Next() {
		ret = append(ret, NewDeclaration((*C.Eolian_Declaration)(p.data)))
	}
	return
}

func (p *Iterator) DocumentationSlice() (ret []*Documentation) {
	for p.Next() {
		ret = append(ret, NewDocumentation((*C.Eolian_Documentation)(p.data)))
	}
	return
}

/*
func (p *Iterator) ValueSlice() (ret []*Value) {
	for p.Next() {
		ret = append(ret, NewValue((C.Eolian_Value)(p.data)))
	}
	return
}
*/
