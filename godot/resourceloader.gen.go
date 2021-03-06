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

//func NewresourceLoaderFromPointer(ptr gdnative.Pointer) resourceLoader {
func new_ResourceLoaderFromPointer(ptr gdnative.Pointer) resourceLoader {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := resourceLoader{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonResourceLoader() *resourceLoader {
	return &resourceLoader{}
}

/*
   Resource Loader. This is a static object accessible as [code]ResourceLoader[/code]. GDScript has a simplified load() function, though.
*/
var ResourceLoader = newSingletonResourceLoader()

/*
Resource Loader. This is a static object accessible as [code]ResourceLoader[/code]. GDScript has a simplified load() function, though.
*/
type resourceLoader struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *resourceLoader) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("_ResourceLoader")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *resourceLoader) BaseClass() string {
	return "_ResourceLoader"
}

/*
        Undocumented
	Args: [{ false path String}], Returns: PoolStringArray
*/
func (o *resourceLoader) GetDependencies(path gdnative.String) gdnative.PoolStringArray {
	o.ensureSingleton()
	//log.Println("Calling _ResourceLoader.GetDependencies()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_ResourceLoader", "get_dependencies")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false type String}], Returns: PoolStringArray
*/
func (o *resourceLoader) GetRecognizedExtensionsForType(aType gdnative.String) gdnative.PoolStringArray {
	o.ensureSingleton()
	//log.Println("Calling _ResourceLoader.GetRecognizedExtensionsForType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_ResourceLoader", "get_recognized_extensions_for_type")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false path String}], Returns: bool
*/
func (o *resourceLoader) Has(path gdnative.String) gdnative.Bool {
	o.ensureSingleton()
	//log.Println("Calling _ResourceLoader.Has()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_ResourceLoader", "has")

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
	Args: [{ false path String} { true type_hint String} {False true p_no_cache bool}], Returns: Resource
*/
func (o *resourceLoader) Load(path gdnative.String, typeHint gdnative.String, pNoCache gdnative.Bool) ResourceImplementer {
	o.ensureSingleton()
	//log.Println("Calling _ResourceLoader.Load()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromString(typeHint)
	ptrArguments[2] = gdnative.NewPointerFromBool(pNoCache)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_ResourceLoader", "load")

	// Call the parent method.
	// Resource
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newResourceFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ResourceImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Resource" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ResourceImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false path String} { true type_hint String}], Returns: ResourceInteractiveLoader
*/
func (o *resourceLoader) LoadInteractive(path gdnative.String, typeHint gdnative.String) ResourceInteractiveLoaderImplementer {
	o.ensureSingleton()
	//log.Println("Calling _ResourceLoader.LoadInteractive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromString(typeHint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_ResourceLoader", "load_interactive")

	// Call the parent method.
	// ResourceInteractiveLoader
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newResourceInteractiveLoaderFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ResourceInteractiveLoaderImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "ResourceInteractiveLoader" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ResourceInteractiveLoaderImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false abort bool}], Returns: void
*/
func (o *resourceLoader) SetAbortOnMissingResources(abort gdnative.Bool) {
	o.ensureSingleton()
	//log.Println("Calling _ResourceLoader.SetAbortOnMissingResources()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(abort)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_ResourceLoader", "set_abort_on_missing_resources")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ResourceLoaderImplementer is an interface that implements the methods
// of the ResourceLoader class.
type ResourceLoaderImplementer interface {
	ObjectImplementer
	GetDependencies(path gdnative.String) gdnative.PoolStringArray
	GetRecognizedExtensionsForType(aType gdnative.String) gdnative.PoolStringArray
	Has(path gdnative.String) gdnative.Bool
	Load(path gdnative.String, typeHint gdnative.String, pNoCache gdnative.Bool) ResourceImplementer
	LoadInteractive(path gdnative.String, typeHint gdnative.String) ResourceInteractiveLoaderImplementer
	SetAbortOnMissingResources(abort gdnative.Bool)
}
