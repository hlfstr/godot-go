package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

// VisualScriptMathConstantMathConstant is an enum for MathConstant values.
type VisualScriptMathConstantMathConstant int

const (
	VisualScriptMathConstantMathConstantE      VisualScriptMathConstantMathConstant = 4
	VisualScriptMathConstantMathConstantHalfPi VisualScriptMathConstantMathConstant = 2
	VisualScriptMathConstantMathConstantInf    VisualScriptMathConstantMathConstant = 6
	VisualScriptMathConstantMathConstantMax    VisualScriptMathConstantMathConstant = 8
	VisualScriptMathConstantMathConstantNan    VisualScriptMathConstantMathConstant = 7
	VisualScriptMathConstantMathConstantOne    VisualScriptMathConstantMathConstant = 0
	VisualScriptMathConstantMathConstantPi     VisualScriptMathConstantMathConstant = 1
	VisualScriptMathConstantMathConstantSqrt2  VisualScriptMathConstantMathConstant = 5
	VisualScriptMathConstantMathConstantTau    VisualScriptMathConstantMathConstant = 3
)

//func NewVisualScriptMathConstantFromPointer(ptr gdnative.Pointer) VisualScriptMathConstant {
func newVisualScriptMathConstantFromPointer(ptr gdnative.Pointer) VisualScriptMathConstant {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptMathConstant{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptMathConstant struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptMathConstant) BaseClass() string {
	return "VisualScriptMathConstant"
}

/*
        Undocumented
	Args: [], Returns: enum.VisualScriptMathConstant::MathConstant
*/
func (o *VisualScriptMathConstant) GetMathConstant() VisualScriptMathConstantMathConstant {
	//log.Println("Calling VisualScriptMathConstant.GetMathConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptMathConstant", "get_math_constant")

	// Call the parent method.
	// enum.VisualScriptMathConstant::MathConstant
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualScriptMathConstantMathConstant(ret)
}

/*
        Undocumented
	Args: [{ false which int}], Returns: void
*/
func (o *VisualScriptMathConstant) SetMathConstant(which gdnative.Int) {
	//log.Println("Calling VisualScriptMathConstant.SetMathConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(which)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptMathConstant", "set_math_constant")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptMathConstantImplementer is an interface that implements the methods
// of the VisualScriptMathConstant class.
type VisualScriptMathConstantImplementer interface {
	VisualScriptNodeImplementer
	SetMathConstant(which gdnative.Int)
}
