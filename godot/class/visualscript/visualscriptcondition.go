package visualscript

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
Undocumented
*/
type VisualScriptCondition struct {
	VisualScriptNode
}

func (o *VisualScriptCondition) BaseClass() string {
	return "VisualScriptCondition"
}

/*
   VisualScriptConditionImplementer is an interface for VisualScriptCondition objects.
*/
type VisualScriptConditionImplementer interface {
	Class
}