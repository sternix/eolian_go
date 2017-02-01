package eolian

/*
#define EFL_BETA_API_SUPPORT 1
#include <Eolian.h>
*/
import "C"

type Documentation struct {
	obj *C.Eolian_Documentation
}

func NewDocumentation(obj *C.Eolian_Documentation) *Documentation {
	return &Documentation{
		obj: obj,
	}
}

func (p *Documentation) IsValid() bool {
	return p.obj != nil
}

func (p *Documentation) Summary() string {
	return GoString(C.eolian_documentation_summary_get(p.obj))
}

func (p *Documentation) Description() string {
	return GoString(C.eolian_documentation_description_get(p.obj))
}

func (p *Documentation) Since() string {
	return GoString(C.eolian_documentation_since_get(p.obj))
}
