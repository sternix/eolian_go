package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type Event struct {
	obj *C.Eolian_Event
}

func NewEvent(obj *C.Eolian_Event) *Event {
	return &Event{
		obj: obj,
	}
}

func (p *Event) IsValid() bool {
	return p.obj != nil
}

func (p *Event) Name() string {
	return GoString(C.eolian_event_name_get(p.obj))
}

func (p *Event) Type() *Type {
	return NewType(C.eolian_event_type_get(p.obj))
}

func (p *Event) Documentation() *Documentation {
	return NewDocumentation(C.eolian_event_documentation_get(p.obj))
}

func (p *Event) Scope() ObjectScope {
	return ObjectScope(C.eolian_event_scope_get(p.obj))
}

func (p *Event) IsBeta() bool {
	return C.eolian_event_is_beta(p.obj) == ETrue
}

func (p *Event) IsHot() bool {
	return C.eolian_event_is_hot(p.obj) == ETrue
}

func (p *Event) IsRestart() bool {
	return C.eolian_event_is_restart(p.obj) == ETrue
}

func (p *Event) CName() string {
	return GoStringFromShared(C.eolian_event_c_name_get(p.obj))
}
