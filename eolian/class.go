package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"
import "unsafe"

type Class struct {
	obj          *C.Eolian_Class
	namespaces   []string
	inherits     []string
	functions    []*Function
	implements   []*Implement
	constructors []*Constructor
	events       []*Event
}

var (
	classes []*Class
)

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
	if classes != nil {
		return classes
	}
	classes = NewIterator(C.eolian_all_classes_get()).ClassSlice()
	return classes
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
	if p.namespaces != nil {
		return p.namespaces
	}
	p.namespaces = NewIterator(C.eolian_class_namespaces_get(p.obj)).StringSlice()
	return p.namespaces
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
	if p.inherits != nil {
		return p.inherits
	}
	p.inherits = NewIterator(C.eolian_class_inherits_get(p.obj)).StringSlice()
	return p.inherits
}

func (p *Class) Functions(t FunctionType) []*Function {
	if p.functions != nil {
		return p.functions
	}
	p.functions = NewIterator(C.eolian_class_functions_get(p.obj, C.Eolian_Function_Type(t))).FunctionSlice()
	return p.functions
}

func (p *Class) FunctionByName(name string, t FunctionType) *Function {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return NewFunction(C.eolian_class_function_get_by_name(p.obj, cname, C.Eolian_Function_Type(t)))
}

func (p *Class) Implements() []*Implement {
	if p.implements != nil {
		return p.implements
	}
	p.implements = NewIterator(C.eolian_class_implements_get(p.obj)).ImplementSlice()
	return p.implements
}

func (p *Class) Constructors() []*Constructor {
	if p.constructors != nil {
		return p.constructors
	}
	p.constructors = NewIterator(C.eolian_class_constructors_get(p.obj)).ConstructorSlice()
	return p.constructors
}

func (p *Class) Events() []*Event {
	if p.events != nil {
		return p.events
	}
	p.events = NewIterator(C.eolian_class_events_get(p.obj)).EventSlice()
	return p.events
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
	return GoString(C.eolian_class_c_get_function_name_get(p.obj), true)
}

func (p *Class) Documentation() *Documentation {
	return NewDocumentation(C.eolian_class_documentation_get(p.obj))
}

/* not in 1.18.4
func (p *Class) CDataType() string {
	return GoString(C.eolian_class_c_data_type_get(p.obj) , true)
}
*/
