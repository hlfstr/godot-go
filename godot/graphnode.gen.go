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

// GraphNodeOverlay is an enum for Overlay values.
type GraphNodeOverlay int

const (
	GraphNodeOverlayBreakpoint GraphNodeOverlay = 1
	GraphNodeOverlayDisabled   GraphNodeOverlay = 0
	GraphNodeOverlayPosition   GraphNodeOverlay = 2
)

//func NewGraphNodeFromPointer(ptr gdnative.Pointer) GraphNode {
func newGraphNodeFromPointer(ptr gdnative.Pointer) GraphNode {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := GraphNode{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A GraphNode is a container defined by a title. It can have 1 or more input and output slots, which can be enabled (shown) or disabled (not shown) and have different (incompatible) types. Colors can also be assigned to slots. A tuple of input and output slots is defined for each GUI element included in the GraphNode. Input and output connections are left and right slots, but only enabled slots are counted as connections.
*/
type GraphNode struct {
	Container
	owner gdnative.Object
}

func (o *GraphNode) BaseClass() string {
	return "GraphNode"
}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *GraphNode) X_GuiInput(arg0 InputEvent) {
	//log.Println("Calling GraphNode.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Disable all input and output slots of the GraphNode.
	Args: [], Returns: void
*/
func (o *GraphNode) ClearAllSlots() {
	//log.Println("Calling GraphNode.ClearAllSlots()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "clear_all_slots")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Disable input and output slot whose index is 'idx'.
	Args: [{ false idx int}], Returns: void
*/
func (o *GraphNode) ClearSlot(idx gdnative.Int) {
	//log.Println("Calling GraphNode.ClearSlot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "clear_slot")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Return the color of the input connection 'idx'.
	Args: [{ false idx int}], Returns: Color
*/
func (o *GraphNode) GetConnectionInputColor(idx gdnative.Int) gdnative.Color {
	//log.Println("Calling GraphNode.GetConnectionInputColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_connection_input_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Return the number of enabled input slots (connections) to the GraphNode.
	Args: [], Returns: int
*/
func (o *GraphNode) GetConnectionInputCount() gdnative.Int {
	//log.Println("Calling GraphNode.GetConnectionInputCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_connection_input_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Return the position of the input connection 'idx'.
	Args: [{ false idx int}], Returns: Vector2
*/
func (o *GraphNode) GetConnectionInputPosition(idx gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling GraphNode.GetConnectionInputPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_connection_input_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Return the type of the input connection 'idx'.
	Args: [{ false idx int}], Returns: int
*/
func (o *GraphNode) GetConnectionInputType(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling GraphNode.GetConnectionInputType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_connection_input_type")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Return the color of the output connection 'idx'.
	Args: [{ false idx int}], Returns: Color
*/
func (o *GraphNode) GetConnectionOutputColor(idx gdnative.Int) gdnative.Color {
	//log.Println("Calling GraphNode.GetConnectionOutputColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_connection_output_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Return the number of enabled output slots (connections) of the GraphNode.
	Args: [], Returns: int
*/
func (o *GraphNode) GetConnectionOutputCount() gdnative.Int {
	//log.Println("Calling GraphNode.GetConnectionOutputCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_connection_output_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Return the position of the output connection 'idx'.
	Args: [{ false idx int}], Returns: Vector2
*/
func (o *GraphNode) GetConnectionOutputPosition(idx gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling GraphNode.GetConnectionOutputPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_connection_output_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Return the type of the output connection 'idx'.
	Args: [{ false idx int}], Returns: int
*/
func (o *GraphNode) GetConnectionOutputType(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling GraphNode.GetConnectionOutputType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_connection_output_type")

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
	Args: [], Returns: Vector2
*/
func (o *GraphNode) GetOffset() gdnative.Vector2 {
	//log.Println("Calling GraphNode.GetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.GraphNode::Overlay
*/
func (o *GraphNode) GetOverlay() GraphNodeOverlay {
	//log.Println("Calling GraphNode.GetOverlay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_overlay")

	// Call the parent method.
	// enum.GraphNode::Overlay
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return GraphNodeOverlay(ret)
}

/*
        Return the color set to 'idx' left (input) slot.
	Args: [{ false idx int}], Returns: Color
*/
func (o *GraphNode) GetSlotColorLeft(idx gdnative.Int) gdnative.Color {
	//log.Println("Calling GraphNode.GetSlotColorLeft()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_slot_color_left")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Return the color set to 'idx' right (output) slot.
	Args: [{ false idx int}], Returns: Color
*/
func (o *GraphNode) GetSlotColorRight(idx gdnative.Int) gdnative.Color {
	//log.Println("Calling GraphNode.GetSlotColorRight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_slot_color_right")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Return the (integer) type of left (input) 'idx' slot.
	Args: [{ false idx int}], Returns: int
*/
func (o *GraphNode) GetSlotTypeLeft(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling GraphNode.GetSlotTypeLeft()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_slot_type_left")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Return the (integer) type of right (output) 'idx' slot.
	Args: [{ false idx int}], Returns: int
*/
func (o *GraphNode) GetSlotTypeRight(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling GraphNode.GetSlotTypeRight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_slot_type_right")

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
	Args: [], Returns: String
*/
func (o *GraphNode) GetTitle() gdnative.String {
	//log.Println("Calling GraphNode.GetTitle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "get_title")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *GraphNode) IsCloseButtonVisible() gdnative.Bool {
	//log.Println("Calling GraphNode.IsCloseButtonVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "is_close_button_visible")

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
func (o *GraphNode) IsComment() gdnative.Bool {
	//log.Println("Calling GraphNode.IsComment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "is_comment")

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
func (o *GraphNode) IsResizable() gdnative.Bool {
	//log.Println("Calling GraphNode.IsResizable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "is_resizable")

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
func (o *GraphNode) IsSelected() gdnative.Bool {
	//log.Println("Calling GraphNode.IsSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "is_selected")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Return true if left (input) slot 'idx' is enabled. False otherwise.
	Args: [{ false idx int}], Returns: bool
*/
func (o *GraphNode) IsSlotEnabledLeft(idx gdnative.Int) gdnative.Bool {
	//log.Println("Calling GraphNode.IsSlotEnabledLeft()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "is_slot_enabled_left")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Return true if right (output) slot 'idx' is enabled. False otherwise.
	Args: [{ false idx int}], Returns: bool
*/
func (o *GraphNode) IsSlotEnabledRight(idx gdnative.Int) gdnative.Bool {
	//log.Println("Calling GraphNode.IsSlotEnabledRight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "is_slot_enabled_right")

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
	Args: [{ false comment bool}], Returns: void
*/
func (o *GraphNode) SetComment(comment gdnative.Bool) {
	//log.Println("Calling GraphNode.SetComment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(comment)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "set_comment")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false offset Vector2}], Returns: void
*/
func (o *GraphNode) SetOffset(offset gdnative.Vector2) {
	//log.Println("Calling GraphNode.SetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "set_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false overlay int}], Returns: void
*/
func (o *GraphNode) SetOverlay(overlay gdnative.Int) {
	//log.Println("Calling GraphNode.SetOverlay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(overlay)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "set_overlay")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false resizable bool}], Returns: void
*/
func (o *GraphNode) SetResizable(resizable gdnative.Bool) {
	//log.Println("Calling GraphNode.SetResizable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(resizable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "set_resizable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false selected bool}], Returns: void
*/
func (o *GraphNode) SetSelected(selected gdnative.Bool) {
	//log.Println("Calling GraphNode.SetSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(selected)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "set_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false show bool}], Returns: void
*/
func (o *GraphNode) SetShowCloseButton(show gdnative.Bool) {
	//log.Println("Calling GraphNode.SetShowCloseButton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(show)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "set_show_close_button")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false idx int} { false enable_left bool} { false type_left int} { false color_left Color} { false enable_right bool} { false type_right int} { false color_right Color} {[Object:null] true custom_left Texture} {[Object:null] true custom_right Texture}], Returns: void
*/
func (o *GraphNode) SetSlot(idx gdnative.Int, enableLeft gdnative.Bool, typeLeft gdnative.Int, colorLeft gdnative.Color, enableRight gdnative.Bool, typeRight gdnative.Int, colorRight gdnative.Color, customLeft Texture, customRight Texture) {
	//log.Println("Calling GraphNode.SetSlot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 9, 9)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromBool(enableLeft)
	ptrArguments[2] = gdnative.NewPointerFromInt(typeLeft)
	ptrArguments[3] = gdnative.NewPointerFromColor(colorLeft)
	ptrArguments[4] = gdnative.NewPointerFromBool(enableRight)
	ptrArguments[5] = gdnative.NewPointerFromInt(typeRight)
	ptrArguments[6] = gdnative.NewPointerFromColor(colorRight)
	ptrArguments[7] = gdnative.NewPointerFromObject(customLeft.GetBaseObject())
	ptrArguments[8] = gdnative.NewPointerFromObject(customRight.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "set_slot")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false title String}], Returns: void
*/
func (o *GraphNode) SetTitle(title gdnative.String) {
	//log.Println("Calling GraphNode.SetTitle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(title)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GraphNode", "set_title")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// GraphNodeImplementer is an interface that implements the methods
// of the GraphNode class.
type GraphNodeImplementer interface {
	ContainerImplementer
	ClearAllSlots()
	ClearSlot(idx gdnative.Int)
	GetConnectionInputColor(idx gdnative.Int) gdnative.Color
	GetConnectionInputCount() gdnative.Int
	GetConnectionInputPosition(idx gdnative.Int) gdnative.Vector2
	GetConnectionInputType(idx gdnative.Int) gdnative.Int
	GetConnectionOutputColor(idx gdnative.Int) gdnative.Color
	GetConnectionOutputCount() gdnative.Int
	GetConnectionOutputPosition(idx gdnative.Int) gdnative.Vector2
	GetConnectionOutputType(idx gdnative.Int) gdnative.Int
	GetOffset() gdnative.Vector2
	GetSlotColorLeft(idx gdnative.Int) gdnative.Color
	GetSlotColorRight(idx gdnative.Int) gdnative.Color
	GetSlotTypeLeft(idx gdnative.Int) gdnative.Int
	GetSlotTypeRight(idx gdnative.Int) gdnative.Int
	GetTitle() gdnative.String
	IsCloseButtonVisible() gdnative.Bool
	IsComment() gdnative.Bool
	IsResizable() gdnative.Bool
	IsSelected() gdnative.Bool
	IsSlotEnabledLeft(idx gdnative.Int) gdnative.Bool
	IsSlotEnabledRight(idx gdnative.Int) gdnative.Bool
	SetComment(comment gdnative.Bool)
	SetOffset(offset gdnative.Vector2)
	SetOverlay(overlay gdnative.Int)
	SetResizable(resizable gdnative.Bool)
	SetSelected(selected gdnative.Bool)
	SetShowCloseButton(show gdnative.Bool)
	SetSlot(idx gdnative.Int, enableLeft gdnative.Bool, typeLeft gdnative.Int, colorLeft gdnative.Color, enableRight gdnative.Bool, typeRight gdnative.Int, colorRight gdnative.Color, customLeft Texture, customRight Texture)
	SetTitle(title gdnative.String)
}
