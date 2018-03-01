package path

import (
	"log"
	"reflect"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
This class is a container/Node-ification of a [Curve3D], so it can have [Spatial] properties and [Node] info.
*/
type Path struct {
	Spatial
}

func (o *Path) BaseClass() string {
	return "Path"
}

/*
   Undocumented
*/
func (o *Path) X_CurveChanged() {
	log.Println("Calling Path.X_CurveChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_curve_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *Path) GetCurve() *Curve3D {
	log.Println("Calling Path.GetCurve()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_curve", goArguments, "*Curve3D")

	returnValue := goRet.Interface().(*Curve3D)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Path) SetCurve(curve *Curve3D) {
	log.Println("Calling Path.SetCurve()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(curve)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_curve", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   PathImplementer is an interface for Path objects.
*/
type PathImplementer interface {
	Class
}