package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"
import "unsafe"

type Variable struct {
	obj        *C.Eolian_Variable
	namespaces []string
}

var (
	allConstants []*Variable
	allVariables []*Variable
)

func NewVariable(obj *C.Eolian_Variable) *Variable {
	return &Variable{
		obj: obj,
	}
}

func GlobalVariableByName(name string) *Variable {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewVariable(C.eolian_variable_global_get_by_name(cname))
}

func ConstantByName(name string) *Variable {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewVariable(C.eolian_variable_constant_get_by_name(cname))
}

func GlobalVariablesByFile(fname string) []*Variable {
	cfname := C.CString(fname)
	defer C.free(unsafe.Pointer(cfname))
	return NewIterator(C.eolian_variable_globals_get_by_file(cfname)).VariableSlice()
}

func ConstantsByFile(fname string) []*Variable {
	cfname := C.CString(fname)
	defer C.free(unsafe.Pointer(cfname))
	return NewIterator(C.eolian_variable_constants_get_by_file(cfname)).VariableSlice()
}

func AllConstants() []*Variable {
	if allConstants != nil {
		return allConstants
	}
	allConstants = NewIterator(C.eolian_variable_all_constants_get()).VariableSlice()
	return allConstants
}

func AllGlobalVariables() []*Variable {
	if allVariables != nil {
		return allVariables
	}
	allVariables = NewIterator(C.eolian_variable_all_globals_get()).VariableSlice()
	return allVariables
}

func (p *Variable) IsValid() bool {
	return p.obj != nil
}

func (p *Variable) Type() VariableType {
	return VariableType(C.eolian_variable_type_get(p.obj))
}

func (p *Variable) Documentation() *Documentation {
	return NewDocumentation(C.eolian_variable_documentation_get(p.obj))
}

func (p *Variable) File() string {
	return GoString(C.eolian_variable_file_get(p.obj))
}

func (p *Variable) BaseType() *Type {
	return NewType(C.eolian_variable_base_type_get(p.obj))
}

func (p *Variable) Value() *Expression {
	return NewExpression(C.eolian_variable_value_get(p.obj))
}

func (p *Variable) Name() string {
	return GoString(C.eolian_variable_name_get(p.obj))
}

func (p *Variable) FullName() string {
	return GoString(C.eolian_variable_full_name_get(p.obj))
}

func (p *Variable) Namespaces() []string {
	if p.namespaces != nil {
		return p.namespaces
	}
	p.namespaces = NewIterator(C.eolian_variable_namespaces_get(p.obj)).StringSlice()
	return p.namespaces
}

func (p *Variable) IsExtern() bool {
	return C.eolian_variable_is_extern(p.obj) == ETrue
}
