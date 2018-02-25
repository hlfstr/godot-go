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
#include <gdnative/array.h>
*/
import "C"

type Array struct {
	base *C.godot_array
}

func (t Array) getBase() *C.godot_array {
	return t.base
}

// Set godot_array_set [[godot_array * p_self] [const godot_int p_idx] [const godot_variant * p_value]] void
func (t *Array) Set(idx Int, value Variant) {
	arg0 := t.getBase()
	arg1 := idx.getBase()
	arg2 := value.getBase()

	C.go_godot_array_set(GDNative.api, arg0, arg1, arg2)

}

// Get godot_array_get [[const godot_array * p_self] [const godot_int p_idx]] godot_variant
func (t *Array) Get(idx Int) Variant {
	arg0 := t.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_array_get(GDNative.api, arg0, arg1)

	return Variant{base: ret}

}

// OperatorIndex godot_array_operator_index [[godot_array * p_self] [const godot_int p_idx]] godot_variant *
func (t *Array) OperatorIndex(idx Int) Variant {
	arg0 := t.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_array_operator_index(GDNative.api, arg0, arg1)

	return Variant{base: ret}

}

// OperatorIndexConst godot_array_operator_index_const [[const godot_array * p_self] [const godot_int p_idx]] const godot_variant *
func (t *Array) OperatorIndexConst(idx Int) Variant {
	arg0 := t.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_array_operator_index_const(GDNative.api, arg0, arg1)

	return Variant{base: ret}

}

// Append godot_array_append [[godot_array * p_self] [const godot_variant * p_value]] void
func (t *Array) Append(value Variant) {
	arg0 := t.getBase()
	arg1 := value.getBase()

	C.go_godot_array_append(GDNative.api, arg0, arg1)

}

// Clear godot_array_clear [[godot_array * p_self]] void
func (t *Array) Clear() {
	arg0 := t.getBase()

	C.go_godot_array_clear(GDNative.api, arg0)

}

// Count godot_array_count [[const godot_array * p_self] [const godot_variant * p_value]] godot_int
func (t *Array) Count(value Variant) Int {
	arg0 := t.getBase()
	arg1 := value.getBase()

	ret := C.go_godot_array_count(GDNative.api, arg0, arg1)

	return Int{base: ret}

}

// Empty godot_array_empty [[const godot_array * p_self]] godot_bool
func (t *Array) Empty() Bool {
	arg0 := t.getBase()

	ret := C.go_godot_array_empty(GDNative.api, arg0)

	return Bool{base: ret}

}

// Erase godot_array_erase [[godot_array * p_self] [const godot_variant * p_value]] void
func (t *Array) Erase(value Variant) {
	arg0 := t.getBase()
	arg1 := value.getBase()

	C.go_godot_array_erase(GDNative.api, arg0, arg1)

}

// Front godot_array_front [[const godot_array * p_self]] godot_variant
func (t *Array) Front() Variant {
	arg0 := t.getBase()

	ret := C.go_godot_array_front(GDNative.api, arg0)

	return Variant{base: ret}

}

// Back godot_array_back [[const godot_array * p_self]] godot_variant
func (t *Array) Back() Variant {
	arg0 := t.getBase()

	ret := C.go_godot_array_back(GDNative.api, arg0)

	return Variant{base: ret}

}

// Find godot_array_find [[const godot_array * p_self] [const godot_variant * p_what] [const godot_int p_from]] godot_int
func (t *Array) Find(what Variant, from Int) Int {
	arg0 := t.getBase()
	arg1 := what.getBase()
	arg2 := from.getBase()

	ret := C.go_godot_array_find(GDNative.api, arg0, arg1, arg2)

	return Int{base: ret}

}

// FindLast godot_array_find_last [[const godot_array * p_self] [const godot_variant * p_what]] godot_int
func (t *Array) FindLast(what Variant) Int {
	arg0 := t.getBase()
	arg1 := what.getBase()

	ret := C.go_godot_array_find_last(GDNative.api, arg0, arg1)

	return Int{base: ret}

}

// Has godot_array_has [[const godot_array * p_self] [const godot_variant * p_value]] godot_bool
func (t *Array) Has(value Variant) Bool {
	arg0 := t.getBase()
	arg1 := value.getBase()

	ret := C.go_godot_array_has(GDNative.api, arg0, arg1)

	return Bool{base: ret}

}

// Hash godot_array_hash [[const godot_array * p_self]] godot_int
func (t *Array) Hash() Int {
	arg0 := t.getBase()

	ret := C.go_godot_array_hash(GDNative.api, arg0)

	return Int{base: ret}

}

// Insert godot_array_insert [[godot_array * p_self] [const godot_int p_pos] [const godot_variant * p_value]] void
func (t *Array) Insert(pos Int, value Variant) {
	arg0 := t.getBase()
	arg1 := pos.getBase()
	arg2 := value.getBase()

	C.go_godot_array_insert(GDNative.api, arg0, arg1, arg2)

}

// Invert godot_array_invert [[godot_array * p_self]] void
func (t *Array) Invert() {
	arg0 := t.getBase()

	C.go_godot_array_invert(GDNative.api, arg0)

}

// PopBack godot_array_pop_back [[godot_array * p_self]] godot_variant
func (t *Array) PopBack() Variant {
	arg0 := t.getBase()

	ret := C.go_godot_array_pop_back(GDNative.api, arg0)

	return Variant{base: ret}

}

// PopFront godot_array_pop_front [[godot_array * p_self]] godot_variant
func (t *Array) PopFront() Variant {
	arg0 := t.getBase()

	ret := C.go_godot_array_pop_front(GDNative.api, arg0)

	return Variant{base: ret}

}

// PushBack godot_array_push_back [[godot_array * p_self] [const godot_variant * p_value]] void
func (t *Array) PushBack(value Variant) {
	arg0 := t.getBase()
	arg1 := value.getBase()

	C.go_godot_array_push_back(GDNative.api, arg0, arg1)

}

// PushFront godot_array_push_front [[godot_array * p_self] [const godot_variant * p_value]] void
func (t *Array) PushFront(value Variant) {
	arg0 := t.getBase()
	arg1 := value.getBase()

	C.go_godot_array_push_front(GDNative.api, arg0, arg1)

}

// Remove godot_array_remove [[godot_array * p_self] [const godot_int p_idx]] void
func (t *Array) Remove(idx Int) {
	arg0 := t.getBase()
	arg1 := idx.getBase()

	C.go_godot_array_remove(GDNative.api, arg0, arg1)

}

// Resize godot_array_resize [[godot_array * p_self] [const godot_int p_size]] void
func (t *Array) Resize(size Int) {
	arg0 := t.getBase()
	arg1 := size.getBase()

	C.go_godot_array_resize(GDNative.api, arg0, arg1)

}

// Rfind godot_array_rfind [[const godot_array * p_self] [const godot_variant * p_what] [const godot_int p_from]] godot_int
func (t *Array) Rfind(what Variant, from Int) Int {
	arg0 := t.getBase()
	arg1 := what.getBase()
	arg2 := from.getBase()

	ret := C.go_godot_array_rfind(GDNative.api, arg0, arg1, arg2)

	return Int{base: ret}

}

// Size godot_array_size [[const godot_array * p_self]] godot_int
func (t *Array) Size() Int {
	arg0 := t.getBase()

	ret := C.go_godot_array_size(GDNative.api, arg0)

	return Int{base: ret}

}

// Sort godot_array_sort [[godot_array * p_self]] void
func (t *Array) Sort() {
	arg0 := t.getBase()

	C.go_godot_array_sort(GDNative.api, arg0)

}

// SortCustom godot_array_sort_custom [[godot_array * p_self] [godot_object * p_obj] [const godot_string * p_func]] void
func (t *Array) SortCustom(obj Object, function String) {
	arg0 := t.getBase()
	arg1 := obj.getBase()
	arg2 := function.getBase()

	C.go_godot_array_sort_custom(GDNative.api, arg0, arg1, arg2)

}

// Bsearch godot_array_bsearch [[godot_array * p_self] [const godot_variant * p_value] [const godot_bool p_before]] godot_int
func (t *Array) Bsearch(value Variant, before Bool) Int {
	arg0 := t.getBase()
	arg1 := value.getBase()
	arg2 := before.getBase()

	ret := C.go_godot_array_bsearch(GDNative.api, arg0, arg1, arg2)

	return Int{base: ret}

}

// BsearchCustom godot_array_bsearch_custom [[godot_array * p_self] [const godot_variant * p_value] [godot_object * p_obj] [const godot_string * p_func] [const godot_bool p_before]] godot_int
func (t *Array) BsearchCustom(value Variant, obj Object, function String, before Bool) Int {
	arg0 := t.getBase()
	arg1 := value.getBase()
	arg2 := obj.getBase()
	arg3 := function.getBase()
	arg4 := before.getBase()

	ret := C.go_godot_array_bsearch_custom(GDNative.api, arg0, arg1, arg2, arg3, arg4)

	return Int{base: ret}

}

// Destroy godot_array_destroy [[godot_array * p_self]] void
func (t *Array) Destroy() {
	arg0 := t.getBase()

	C.go_godot_array_destroy(GDNative.api, arg0)

}