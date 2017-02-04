package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"
import "unsafe"

type Declaration struct {
	obj *C.Eolian_Declaration
}

var (
	declarations []*Declaration
)

func NewDeclaration(obj *C.Eolian_Declaration) *Declaration {
	return &Declaration{
		obj: obj,
	}
}

func DeclarationByName(name string) *Declaration {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewDeclaration(C.eolian_declaration_get_by_name(cname))
}

func DeclarationByFile(fname string) []*Declaration {
	cfname := C.CString(fname)
	defer C.free(unsafe.Pointer(cfname))
	return NewIterator(C.eolian_declarations_get_by_file(cfname)).DeclarationSlice()
}

func AllDeclarations() []*Declaration {
	if declarations != nil {
		return declarations
	}
	declarations = NewIterator(C.eolian_all_declarations_get()).DeclarationSlice()
	return declarations
}

func (p *Declaration) IsValid() bool {
	return p.obj != nil
}

func (p *Declaration) Type() DeclarationType {
	return DeclarationType(C.eolian_declaration_type_get(p.obj))
}

func (p *Declaration) Name() string {
	return GoString(C.eolian_declaration_name_get(p.obj))
}

func (p *Declaration) Class() *Class {
	return NewClass(C.eolian_declaration_class_get(p.obj))
}

func (p *Declaration) DataType() *Typedecl {
	return NewTypedecl(C.eolian_declaration_data_type_get(p.obj))
}

func (p *Declaration) Variable() *Variable {
	return NewVariable(C.eolian_declaration_variable_get(p.obj))
}
