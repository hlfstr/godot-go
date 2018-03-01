package toolbutton

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
This is a helper class to generate a flat [Button] (see [method Button.set_flat]), creating a ToolButton is equivalent to: [codeblock] var btn = Button.new() btn.set_flat(true) [/codeblock]
*/
type ToolButton struct {
	Button
}

func (o *ToolButton) BaseClass() string {
	return "ToolButton"
}

/*
   ToolButtonImplementer is an interface for ToolButton objects.
*/
type ToolButtonImplementer interface {
	Class
}