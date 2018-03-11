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

// TileSetAutotileBindings is an enum for AutotileBindings values.
type TileSetAutotileBindings int

const (
	TileSetBindBottom      TileSetAutotileBindings = 128
	TileSetBindBottomleft  TileSetAutotileBindings = 64
	TileSetBindBottomright TileSetAutotileBindings = 256
	TileSetBindLeft        TileSetAutotileBindings = 8
	TileSetBindRight       TileSetAutotileBindings = 32
	TileSetBindTop         TileSetAutotileBindings = 2
	TileSetBindTopleft     TileSetAutotileBindings = 1
	TileSetBindTopright    TileSetAutotileBindings = 4
)

// TileSetBitmaskMode is an enum for BitmaskMode values.
type TileSetBitmaskMode int

const (
	TileSetBitmask2X2 TileSetBitmaskMode = 0
	TileSetBitmask3X3 TileSetBitmaskMode = 1
)

//func NewTileSetFromPointer(ptr gdnative.Pointer) TileSet {
func newTileSetFromPointer(ptr gdnative.Pointer) TileSet {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := TileSet{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A TileSet is a library of tiles for a [TileMap]. It contains a list of tiles, each consisting of a sprite and optional collision shapes. Tiles are referenced by a unique integer ID.
*/
type TileSet struct {
	Resource
	owner gdnative.Object
}

func (o *TileSet) BaseClass() string {
	return "TileSet"
}

/*

	Args: [{ false autotile_id int} { false bitmask int} { false tilemap Object} { false tile_location Vector2}], Returns: Vector2
*/
func (o *TileSet) X_ForwardSubtileSelection(autotileId gdnative.Int, bitmask gdnative.Int, tilemap Object, tileLocation gdnative.Vector2) gdnative.Vector2 {
	//log.Println("Calling TileSet.X_ForwardSubtileSelection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(autotileId)
	ptrArguments[1] = gdnative.NewPointerFromInt(bitmask)
	ptrArguments[2] = gdnative.NewPointerFromObject(tilemap.GetBaseObject())
	ptrArguments[3] = gdnative.NewPointerFromVector2(tileLocation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "_forward_subtile_selection")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false drawn_id int} { false neighbor_id int}], Returns: bool
*/
func (o *TileSet) X_IsTileBound(drawnId gdnative.Int, neighborId gdnative.Int) gdnative.Bool {
	//log.Println("Calling TileSet.X_IsTileBound()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(drawnId)
	ptrArguments[1] = gdnative.NewPointerFromInt(neighborId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "_is_tile_bound")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false id int}], Returns: enum.TileSet::BitmaskMode
*/
func (o *TileSet) AutotileGetBitmaskMode(id gdnative.Int) TileSetBitmaskMode {
	//log.Println("Calling TileSet.AutotileGetBitmaskMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "autotile_get_bitmask_mode")

	// Call the parent method.
	// enum.TileSet::BitmaskMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return TileSetBitmaskMode(ret)
}

/*

	Args: [{ false id int} { false mode int}], Returns: void
*/
func (o *TileSet) AutotileSetBitmaskMode(id gdnative.Int, mode gdnative.Int) {
	//log.Println("Calling TileSet.AutotileSetBitmaskMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "autotile_set_bitmask_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clear all tiles.
	Args: [], Returns: void
*/
func (o *TileSet) Clear() {
	//log.Println("Calling TileSet.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Create a new tile which will be referenced by the given ID.
	Args: [{ false id int}], Returns: void
*/
func (o *TileSet) CreateTile(id gdnative.Int) {
	//log.Println("Calling TileSet.CreateTile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "create_tile")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Find the first tile matching the given name.
	Args: [{ false name String}], Returns: int
*/
func (o *TileSet) FindTileByName(name gdnative.String) gdnative.Int {
	//log.Println("Calling TileSet.FindTileByName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "find_tile_by_name")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Return the ID following the last currently used ID, useful when creating a new tile.
	Args: [], Returns: int
*/
func (o *TileSet) GetLastUnusedTileId() gdnative.Int {
	//log.Println("Calling TileSet.GetLastUnusedTileId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "get_last_unused_tile_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Return an array of all currently used tile IDs.
	Args: [], Returns: Array
*/
func (o *TileSet) GetTilesIds() gdnative.Array {
	//log.Println("Calling TileSet.GetTilesIds()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "get_tiles_ids")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Remove the tile referenced by the given ID.
	Args: [{ false id int}], Returns: void
*/
func (o *TileSet) RemoveTile(id gdnative.Int) {
	//log.Println("Calling TileSet.RemoveTile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "remove_tile")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false id int} { false shape Shape2D} { false shape_transform Transform2D} {False true one_way bool} {(0, 0) true autotile_coord Vector2}], Returns: void
*/
func (o *TileSet) TileAddShape(id gdnative.Int, shape Shape2D, shapeTransform gdnative.Transform2D, oneWay gdnative.Bool, autotileCoord gdnative.Vector2) {
	//log.Println("Calling TileSet.TileAddShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromObject(shape.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromTransform2D(shapeTransform)
	ptrArguments[3] = gdnative.NewPointerFromBool(oneWay)
	ptrArguments[4] = gdnative.NewPointerFromVector2(autotileCoord)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_add_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Return the light occluder of the tile.
	Args: [{ false id int}], Returns: OccluderPolygon2D
*/
func (o *TileSet) TileGetLightOccluder(id gdnative.Int) OccluderPolygon2DImplementer {
	//log.Println("Calling TileSet.TileGetLightOccluder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_light_occluder")

	// Call the parent method.
	// OccluderPolygon2D
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newOccluderPolygon2DFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(OccluderPolygon2DImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "OccluderPolygon2D" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(OccluderPolygon2DImplementer)
	}

	return &ret
}

/*
        Return the material of the tile.
	Args: [{ false id int}], Returns: ShaderMaterial
*/
func (o *TileSet) TileGetMaterial(id gdnative.Int) ShaderMaterialImplementer {
	//log.Println("Calling TileSet.TileGetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_material")

	// Call the parent method.
	// ShaderMaterial
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newShaderMaterialFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ShaderMaterialImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "ShaderMaterial" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ShaderMaterialImplementer)
	}

	return &ret
}

/*
        Return the name of the tile.
	Args: [{ false id int}], Returns: String
*/
func (o *TileSet) TileGetName(id gdnative.Int) gdnative.String {
	//log.Println("Calling TileSet.TileGetName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Return the navigation polygon of the tile.
	Args: [{ false id int}], Returns: NavigationPolygon
*/
func (o *TileSet) TileGetNavigationPolygon(id gdnative.Int) NavigationPolygonImplementer {
	//log.Println("Calling TileSet.TileGetNavigationPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_navigation_polygon")

	// Call the parent method.
	// NavigationPolygon
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newNavigationPolygonFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(NavigationPolygonImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "NavigationPolygon" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(NavigationPolygonImplementer)
	}

	return &ret
}

/*
        Return the offset of the tile's navigation polygon.
	Args: [{ false id int}], Returns: Vector2
*/
func (o *TileSet) TileGetNavigationPolygonOffset(id gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling TileSet.TileGetNavigationPolygonOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_navigation_polygon_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false id int}], Returns: Texture
*/
func (o *TileSet) TileGetNormalMap(id gdnative.Int) TextureImplementer {
	//log.Println("Calling TileSet.TileGetNormalMap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_normal_map")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Return the offset of the tile's light occluder.
	Args: [{ false id int}], Returns: Vector2
*/
func (o *TileSet) TileGetOccluderOffset(id gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling TileSet.TileGetOccluderOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_occluder_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Return the tile sub-region in the texture.
	Args: [{ false id int}], Returns: Rect2
*/
func (o *TileSet) TileGetRegion(id gdnative.Int) gdnative.Rect2 {
	//log.Println("Calling TileSet.TileGetRegion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_region")

	// Call the parent method.
	// Rect2
	retPtr := gdnative.NewEmptyRect2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRect2FromPointer(retPtr)
	return ret
}

/*

	Args: [{ false id int} { false shape_id int}], Returns: Shape2D
*/
func (o *TileSet) TileGetShape(id gdnative.Int, shapeId gdnative.Int) Shape2DImplementer {
	//log.Println("Calling TileSet.TileGetShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(shapeId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_shape")

	// Call the parent method.
	// Shape2D
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newShape2DFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(Shape2DImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Shape2D" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(Shape2DImplementer)
	}

	return &ret
}

/*

	Args: [{ false id int}], Returns: int
*/
func (o *TileSet) TileGetShapeCount(id gdnative.Int) gdnative.Int {
	//log.Println("Calling TileSet.TileGetShapeCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_shape_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false id int} { false shape_id int}], Returns: bool
*/
func (o *TileSet) TileGetShapeOneWay(id gdnative.Int, shapeId gdnative.Int) gdnative.Bool {
	//log.Println("Calling TileSet.TileGetShapeOneWay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(shapeId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_shape_one_way")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false id int} { false shape_id int}], Returns: Transform2D
*/
func (o *TileSet) TileGetShapeTransform(id gdnative.Int, shapeId gdnative.Int) gdnative.Transform2D {
	//log.Println("Calling TileSet.TileGetShapeTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(shapeId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_shape_transform")

	// Call the parent method.
	// Transform2D
	retPtr := gdnative.NewEmptyTransform2D()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransform2DFromPointer(retPtr)
	return ret
}

/*
        Return the array of shapes of the tile.
	Args: [{ false id int}], Returns: Array
*/
func (o *TileSet) TileGetShapes(id gdnative.Int) gdnative.Array {
	//log.Println("Calling TileSet.TileGetShapes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_shapes")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Return the texture of the tile.
	Args: [{ false id int}], Returns: Texture
*/
func (o *TileSet) TileGetTexture(id gdnative.Int) TextureImplementer {
	//log.Println("Calling TileSet.TileGetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Return the texture offset of the tile.
	Args: [{ false id int}], Returns: Vector2
*/
func (o *TileSet) TileGetTextureOffset(id gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling TileSet.TileGetTextureOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_get_texture_offset")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Set a light occluder for the tile.
	Args: [{ false id int} { false light_occluder OccluderPolygon2D}], Returns: void
*/
func (o *TileSet) TileSetLightOccluder(id gdnative.Int, lightOccluder OccluderPolygon2D) {
	//log.Println("Calling TileSet.TileSetLightOccluder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromObject(lightOccluder.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_light_occluder")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the material of the tile.
	Args: [{ false id int} { false material ShaderMaterial}], Returns: void
*/
func (o *TileSet) TileSetMaterial(id gdnative.Int, material ShaderMaterial) {
	//log.Println("Calling TileSet.TileSetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the name of the tile, for descriptive purposes.
	Args: [{ false id int} { false name String}], Returns: void
*/
func (o *TileSet) TileSetName(id gdnative.Int, name gdnative.String) {
	//log.Println("Calling TileSet.TileSetName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set a navigation polygon for the tile.
	Args: [{ false id int} { false navigation_polygon NavigationPolygon}], Returns: void
*/
func (o *TileSet) TileSetNavigationPolygon(id gdnative.Int, navigationPolygon NavigationPolygon) {
	//log.Println("Calling TileSet.TileSetNavigationPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromObject(navigationPolygon.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_navigation_polygon")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set an offset for the tile's navigation polygon.
	Args: [{ false id int} { false navigation_polygon_offset Vector2}], Returns: void
*/
func (o *TileSet) TileSetNavigationPolygonOffset(id gdnative.Int, navigationPolygonOffset gdnative.Vector2) {
	//log.Println("Calling TileSet.TileSetNavigationPolygonOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromVector2(navigationPolygonOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_navigation_polygon_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false id int} { false normal_map Texture}], Returns: void
*/
func (o *TileSet) TileSetNormalMap(id gdnative.Int, normalMap Texture) {
	//log.Println("Calling TileSet.TileSetNormalMap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromObject(normalMap.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_normal_map")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set an offset for the tile's light occluder.
	Args: [{ false id int} { false occluder_offset Vector2}], Returns: void
*/
func (o *TileSet) TileSetOccluderOffset(id gdnative.Int, occluderOffset gdnative.Vector2) {
	//log.Println("Calling TileSet.TileSetOccluderOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromVector2(occluderOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_occluder_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the tile sub-region in the texture. This is common in texture atlases.
	Args: [{ false id int} { false region Rect2}], Returns: void
*/
func (o *TileSet) TileSetRegion(id gdnative.Int, region gdnative.Rect2) {
	//log.Println("Calling TileSet.TileSetRegion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromRect2(region)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_region")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false id int} { false shape_id int} { false shape Shape2D}], Returns: void
*/
func (o *TileSet) TileSetShape(id gdnative.Int, shapeId gdnative.Int, shape Shape2D) {
	//log.Println("Calling TileSet.TileSetShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(shapeId)
	ptrArguments[2] = gdnative.NewPointerFromObject(shape.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false id int} { false shape_id int} { false one_way bool}], Returns: void
*/
func (o *TileSet) TileSetShapeOneWay(id gdnative.Int, shapeId gdnative.Int, oneWay gdnative.Bool) {
	//log.Println("Calling TileSet.TileSetShapeOneWay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(shapeId)
	ptrArguments[2] = gdnative.NewPointerFromBool(oneWay)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_shape_one_way")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false id int} { false shape_id int} { false shape_transform Transform2D}], Returns: void
*/
func (o *TileSet) TileSetShapeTransform(id gdnative.Int, shapeId gdnative.Int, shapeTransform gdnative.Transform2D) {
	//log.Println("Calling TileSet.TileSetShapeTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromInt(shapeId)
	ptrArguments[2] = gdnative.NewPointerFromTransform2D(shapeTransform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_shape_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set an array of shapes for the tile, enabling physics to collide with it.
	Args: [{ false id int} { false shapes Array}], Returns: void
*/
func (o *TileSet) TileSetShapes(id gdnative.Int, shapes gdnative.Array) {
	//log.Println("Calling TileSet.TileSetShapes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromArray(shapes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_shapes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the texture of the tile.
	Args: [{ false id int} { false texture Texture}], Returns: void
*/
func (o *TileSet) TileSetTexture(id gdnative.Int, texture Texture) {
	//log.Println("Calling TileSet.TileSetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the texture offset of the tile.
	Args: [{ false id int} { false texture_offset Vector2}], Returns: void
*/
func (o *TileSet) TileSetTextureOffset(id gdnative.Int, textureOffset gdnative.Vector2) {
	//log.Println("Calling TileSet.TileSetTextureOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromVector2(textureOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TileSet", "tile_set_texture_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// TileSetImplementer is an interface that implements the methods
// of the TileSet class.
type TileSetImplementer interface {
	ResourceImplementer
	X_ForwardSubtileSelection(autotileId gdnative.Int, bitmask gdnative.Int, tilemap Object, tileLocation gdnative.Vector2) gdnative.Vector2
	X_IsTileBound(drawnId gdnative.Int, neighborId gdnative.Int) gdnative.Bool
	AutotileSetBitmaskMode(id gdnative.Int, mode gdnative.Int)
	Clear()
	CreateTile(id gdnative.Int)
	FindTileByName(name gdnative.String) gdnative.Int
	GetLastUnusedTileId() gdnative.Int
	GetTilesIds() gdnative.Array
	RemoveTile(id gdnative.Int)
	TileAddShape(id gdnative.Int, shape Shape2D, shapeTransform gdnative.Transform2D, oneWay gdnative.Bool, autotileCoord gdnative.Vector2)
	TileGetLightOccluder(id gdnative.Int) OccluderPolygon2DImplementer
	TileGetMaterial(id gdnative.Int) ShaderMaterialImplementer
	TileGetName(id gdnative.Int) gdnative.String
	TileGetNavigationPolygon(id gdnative.Int) NavigationPolygonImplementer
	TileGetNavigationPolygonOffset(id gdnative.Int) gdnative.Vector2
	TileGetNormalMap(id gdnative.Int) TextureImplementer
	TileGetOccluderOffset(id gdnative.Int) gdnative.Vector2
	TileGetRegion(id gdnative.Int) gdnative.Rect2
	TileGetShape(id gdnative.Int, shapeId gdnative.Int) Shape2DImplementer
	TileGetShapeCount(id gdnative.Int) gdnative.Int
	TileGetShapeOneWay(id gdnative.Int, shapeId gdnative.Int) gdnative.Bool
	TileGetShapeTransform(id gdnative.Int, shapeId gdnative.Int) gdnative.Transform2D
	TileGetShapes(id gdnative.Int) gdnative.Array
	TileGetTexture(id gdnative.Int) TextureImplementer
	TileGetTextureOffset(id gdnative.Int) gdnative.Vector2
	TileSetLightOccluder(id gdnative.Int, lightOccluder OccluderPolygon2D)
	TileSetMaterial(id gdnative.Int, material ShaderMaterial)
	TileSetName(id gdnative.Int, name gdnative.String)
	TileSetNavigationPolygon(id gdnative.Int, navigationPolygon NavigationPolygon)
	TileSetNavigationPolygonOffset(id gdnative.Int, navigationPolygonOffset gdnative.Vector2)
	TileSetNormalMap(id gdnative.Int, normalMap Texture)
	TileSetOccluderOffset(id gdnative.Int, occluderOffset gdnative.Vector2)
	TileSetRegion(id gdnative.Int, region gdnative.Rect2)
	TileSetShape(id gdnative.Int, shapeId gdnative.Int, shape Shape2D)
	TileSetShapeOneWay(id gdnative.Int, shapeId gdnative.Int, oneWay gdnative.Bool)
	TileSetShapeTransform(id gdnative.Int, shapeId gdnative.Int, shapeTransform gdnative.Transform2D)
	TileSetShapes(id gdnative.Int, shapes gdnative.Array)
	TileSetTexture(id gdnative.Int, texture Texture)
	TileSetTextureOffset(id gdnative.Int, textureOffset gdnative.Vector2)
}
