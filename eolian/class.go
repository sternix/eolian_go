package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"
import "unsafe"

type Class struct {
	obj *C.Eolian_Class
}

func NewClass(obj *C.Eolian_Class) *Class {
	return &Class{
		obj: obj,
	}
}

func ClassByName(name string) *Class {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewClass(C.eolian_class_get_by_name(cname))
}

func ClassByFile(file string) *Class {
	cfile := C.CString(file)
	defer C.free(unsafe.Pointer(cfile))
	return NewClass(C.eolian_class_get_by_file(cfile))
}

func AllClasses() []*Class {
	return NewIterator(C.eolian_all_classes_get()).ClassSlice()
}

func (p *Class) IsValid() bool {
	return p.obj != nil
}

func (p *Class) File() string {
	return GoString(C.eolian_class_file_get(p.obj))
}

func (p *Class) FullName() string {
	return GoString(C.eolian_class_full_name_get(p.obj))
}

func (p *Class) Name() string {
	return GoString(C.eolian_class_name_get(p.obj))
}

func (p *Class) Namespaces() []string {
	ns := C.eolian_class_namespaces_get(p.obj)
	return NewIterator(ns).StringSlice()
}

func (p *Class) Type() ClassType {
	return ClassType(C.eolian_class_type_get(p.obj))
}

func (p *Class) LegacyPrefix() string {
	return GoString(C.eolian_class_legacy_prefix_get(p.obj))
}

func (p *Class) EoPrefix() string {
	return GoString(C.eolian_class_eo_prefix_get(p.obj))
}

func (p *Class) EventPrefix() string {
	return GoString(C.eolian_class_event_prefix_get(p.obj))
}

func (p *Class) DataType() string {
	return GoString(C.eolian_class_data_type_get(p.obj))
}

func (p *Class) Inherits() []string {
	return NewIterator(C.eolian_class_inherits_get(p.obj)).StringSlice()
}

func (p *Class) Functions(t FunctionType) []*Function {
	return NewIterator(C.eolian_class_functions_get(p.obj, C.Eolian_Function_Type(t))).FunctionSlice()
}

func (p *Class) FunctionByName(name string, t FunctionType) *Function {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewFunction(C.eolian_class_function_get_by_name(p.obj, cname, C.Eolian_Function_Type(t)))
}

func (p *Class) Implements() []*Implement {
	return NewIterator(C.eolian_class_implements_get(p.obj)).ImplementSlice()
}

func (p *Class) Constructors() []*Constructor {
	return NewIterator(C.eolian_class_constructors_get(p.obj)).ConstructorSlice()
}

func (p *Class) Events() []*Event {
	return NewIterator(C.eolian_class_events_get(p.obj)).EventSlice()
}

func (p *Class) EventByName(name string) *Event {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewEvent(C.eolian_class_event_get_by_name(p.obj, cname))
}

func (p *Class) ConstructorEnabled() bool {
	return C.eolian_class_ctor_enable_get(p.obj) == ETrue
}

func (p *Class) DestructorEnabled() bool {
	return C.eolian_class_dtor_enable_get(p.obj) == ETrue
}

func (p *Class) CGetFunctionName() string {
	return GoStringFromShared(C.eolian_class_c_get_function_name_get(p.obj))
}

func (p *Class) Documentation() *Documentation {
	return NewDocumentation(C.eolian_class_documentation_get(p.obj))
}

/* not in 1.18.4
func (p *Class) CDataType() string {
	return GoStringFromShared(C.eolian_class_c_data_type_get(p.obj))
}
*/
