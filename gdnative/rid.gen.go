package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#include <gdnative/rid.h>
*/
import "C"

type Rid struct {
	base *C.godot_rid
}

func (t *Rid) getBase() *C.godot_rid {
	return t.base
}

// NewRid godot_rid_new [[godot_rid * r_dest]]

//func NewRid(dest Rid, ) *Rid {
//	return &Rid{}
//}

// GetId godot_rid_get_id [[const godot_rid * p_self]]

// NewRidWithResource godot_rid_new_with_resource [[godot_rid * r_dest] [const godot_object * p_from]]

//func NewRidWithResource(dest Rid, from ConstObject, ) *Rid {
//	return &Rid{}
//}

// OperatorEqual godot_rid_operator_equal [[const godot_rid * p_self] [const godot_rid * p_b]]

// OperatorLess godot_rid_operator_less [[const godot_rid * p_self] [const godot_rid * p_b]]
