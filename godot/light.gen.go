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

// LightBakeMode is an enum for BakeMode values.
type LightBakeMode int

const (
	LightBakeAll      LightBakeMode = 2
	LightBakeDisabled LightBakeMode = 0
	LightBakeIndirect LightBakeMode = 1
)

// LightParam is an enum for Param values.
type LightParam int

const (
	LightParamAttenuation          LightParam = 4
	LightParamContactShadowSize    LightParam = 7
	LightParamEnergy               LightParam = 0
	LightParamIndirectEnergy       LightParam = 1
	LightParamMax                  LightParam = 15
	LightParamRange                LightParam = 3
	LightParamShadowBias           LightParam = 13
	LightParamShadowBiasSplitScale LightParam = 14
	LightParamShadowMaxDistance    LightParam = 8
	LightParamShadowNormalBias     LightParam = 12
	LightParamShadowSplit1Offset   LightParam = 9
	LightParamShadowSplit2Offset   LightParam = 10
	LightParamShadowSplit3Offset   LightParam = 11
	LightParamSpecular             LightParam = 2
	LightParamSpotAngle            LightParam = 5
	LightParamSpotAttenuation      LightParam = 6
)

//func NewLightFromPointer(ptr gdnative.Pointer) Light {
func newLightFromPointer(ptr gdnative.Pointer) Light {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Light{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Light is the abstract base class for light nodes, so it shouldn't be used directly (It can't be instanced). Other types of light nodes inherit from it. Light contains the common variables and parameters used for lighting.
*/
type Light struct {
	VisualInstance
	owner gdnative.Object
}

func (o *Light) BaseClass() string {
	return "Light"
}

/*
        Undocumented
	Args: [], Returns: enum.Light::BakeMode
*/
func (o *Light) GetBakeMode() LightBakeMode {
	//log.Println("Calling Light.GetBakeMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "get_bake_mode")

	// Call the parent method.
	// enum.Light::BakeMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return LightBakeMode(ret)
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *Light) GetColor() gdnative.Color {
	//log.Println("Calling Light.GetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "get_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Light) GetCullMask() gdnative.Int {
	//log.Println("Calling Light.GetCullMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "get_cull_mask")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false param int}], Returns: float
*/
func (o *Light) GetParam(param gdnative.Int) gdnative.Float {
	//log.Println("Calling Light.GetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "get_param")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *Light) GetShadowColor() gdnative.Color {
	//log.Println("Calling Light.GetShadowColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "get_shadow_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Light) GetShadowReverseCullFace() gdnative.Bool {
	//log.Println("Calling Light.GetShadowReverseCullFace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "get_shadow_reverse_cull_face")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Light) HasShadow() gdnative.Bool {
	//log.Println("Calling Light.HasShadow()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "has_shadow")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Light) IsEditorOnly() gdnative.Bool {
	//log.Println("Calling Light.IsEditorOnly()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "is_editor_only")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Light) IsNegative() gdnative.Bool {
	//log.Println("Calling Light.IsNegative()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "is_negative")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false bake_mode int}], Returns: void
*/
func (o *Light) SetBakeMode(bakeMode gdnative.Int) {
	//log.Println("Calling Light.SetBakeMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bakeMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "set_bake_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *Light) SetColor(color gdnative.Color) {
	//log.Println("Calling Light.SetColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "set_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false cull_mask int}], Returns: void
*/
func (o *Light) SetCullMask(cullMask gdnative.Int) {
	//log.Println("Calling Light.SetCullMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(cullMask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "set_cull_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false editor_only bool}], Returns: void
*/
func (o *Light) SetEditorOnly(editorOnly gdnative.Bool) {
	//log.Println("Calling Light.SetEditorOnly()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(editorOnly)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "set_editor_only")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Light) SetNegative(enabled gdnative.Bool) {
	//log.Println("Calling Light.SetNegative()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "set_negative")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false param int} { false value float}], Returns: void
*/
func (o *Light) SetParam(param gdnative.Int, value gdnative.Float) {
	//log.Println("Calling Light.SetParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromFloat(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "set_param")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Light) SetShadow(enabled gdnative.Bool) {
	//log.Println("Calling Light.SetShadow()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "set_shadow")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false shadow_color Color}], Returns: void
*/
func (o *Light) SetShadowColor(shadowColor gdnative.Color) {
	//log.Println("Calling Light.SetShadowColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(shadowColor)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "set_shadow_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *Light) SetShadowReverseCullFace(enable gdnative.Bool) {
	//log.Println("Calling Light.SetShadowReverseCullFace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Light", "set_shadow_reverse_cull_face")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// LightImplementer is an interface that implements the methods
// of the Light class.
type LightImplementer interface {
	VisualInstanceImplementer
	GetColor() gdnative.Color
	GetCullMask() gdnative.Int
	GetParam(param gdnative.Int) gdnative.Float
	GetShadowColor() gdnative.Color
	GetShadowReverseCullFace() gdnative.Bool
	HasShadow() gdnative.Bool
	IsEditorOnly() gdnative.Bool
	IsNegative() gdnative.Bool
	SetBakeMode(bakeMode gdnative.Int)
	SetColor(color gdnative.Color)
	SetCullMask(cullMask gdnative.Int)
	SetEditorOnly(editorOnly gdnative.Bool)
	SetNegative(enabled gdnative.Bool)
	SetParam(param gdnative.Int, value gdnative.Float)
	SetShadow(enabled gdnative.Bool)
	SetShadowColor(shadowColor gdnative.Color)
	SetShadowReverseCullFace(enable gdnative.Bool)
}
