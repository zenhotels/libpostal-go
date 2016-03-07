// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Mon, 07 Mar 2016 22:25:38 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package postal

/*
#cgo pkg-config: libpostal
#include "libpostal.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocNormalizeOptionsMemory allocates memory for type C.normalize_options_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNormalizeOptionsMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNormalizeOptionsValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNormalizeOptionsValue = unsafe.Sizeof([1]C.normalize_options_t{})

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// allocPCharMemory allocates memory for type *C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPCharValue = unsafe.Sizeof([1]*C.char{})

// unpackSSByte transforms a sliced Go data structure into plain C format.
func unpackSSByte(x [][]byte) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.char)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}

// packSSByte reads sliced Go data structure out from plain C format.
func packSSByte(v [][]byte, ptr0 **C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m]*C.char)(unsafe.Pointer(ptr0)))[i0]
		hxfc4425b := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxfc4425b.Data = uintptr(unsafe.Pointer(ptr1))
		hxfc4425b.Cap = 0x7fffffff
		// hxfc4425b.Len = ?
	}
}

// Ref returns a reference.
func (x *NormalizeOptions) Ref() *C.normalize_options_t {
	if x == nil {
		return nil
	}
	return x.ref23955150
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *NormalizeOptions) Free() {
	if x != nil && x.allocs23955150 != nil {
		x.allocs23955150.(*cgoAllocMap).Free()
		x.ref23955150 = nil
	}
}

// NewNormalizeOptionsRef initialises a new struct holding the reference to the originaitng C struct.
func NewNormalizeOptionsRef(ref *C.normalize_options_t) *NormalizeOptions {
	if ref == nil {
		return nil
	}
	obj := new(NormalizeOptions)
	obj.ref23955150 = ref
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *NormalizeOptions) PassRef() (*C.normalize_options_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref23955150 != nil {
		return x.ref23955150, nil
	}
	mem23955150 := allocNormalizeOptionsMemory(1)
	ref23955150 := (*C.normalize_options_t)(mem23955150)
	allocs23955150 := new(cgoAllocMap)
	var clanguages_allocs *cgoAllocMap
	ref23955150.languages, clanguages_allocs = unpackSSByte(x.Languages)
	allocs23955150.Borrow(clanguages_allocs)

	var cnum_languages_allocs *cgoAllocMap
	ref23955150.num_languages, cnum_languages_allocs = (C.int)(x.NumLanguages), cgoAllocsUnknown
	allocs23955150.Borrow(cnum_languages_allocs)

	var caddress_components_allocs *cgoAllocMap
	ref23955150.address_components, caddress_components_allocs = (C.uint16_t)(x.AddressComponents), cgoAllocsUnknown
	allocs23955150.Borrow(caddress_components_allocs)

	var clatin_ascii_allocs *cgoAllocMap
	ref23955150.latin_ascii, clatin_ascii_allocs = (C._Bool)(x.LatinAscii), cgoAllocsUnknown
	allocs23955150.Borrow(clatin_ascii_allocs)

	var ctransliterate_allocs *cgoAllocMap
	ref23955150.transliterate, ctransliterate_allocs = (C._Bool)(x.Transliterate), cgoAllocsUnknown
	allocs23955150.Borrow(ctransliterate_allocs)

	var cstrip_accents_allocs *cgoAllocMap
	ref23955150.strip_accents, cstrip_accents_allocs = (C._Bool)(x.StripAccents), cgoAllocsUnknown
	allocs23955150.Borrow(cstrip_accents_allocs)

	var cdecompose_allocs *cgoAllocMap
	ref23955150.decompose, cdecompose_allocs = (C._Bool)(x.Decompose), cgoAllocsUnknown
	allocs23955150.Borrow(cdecompose_allocs)

	var clowercase_allocs *cgoAllocMap
	ref23955150.lowercase, clowercase_allocs = (C._Bool)(x.Lowercase), cgoAllocsUnknown
	allocs23955150.Borrow(clowercase_allocs)

	var ctrim_string_allocs *cgoAllocMap
	ref23955150.trim_string, ctrim_string_allocs = (C._Bool)(x.TrimString), cgoAllocsUnknown
	allocs23955150.Borrow(ctrim_string_allocs)

	var cdrop_parentheticals_allocs *cgoAllocMap
	ref23955150.drop_parentheticals, cdrop_parentheticals_allocs = (C._Bool)(x.DropParentheticals), cgoAllocsUnknown
	allocs23955150.Borrow(cdrop_parentheticals_allocs)

	var creplace_numeric_hyphens_allocs *cgoAllocMap
	ref23955150.replace_numeric_hyphens, creplace_numeric_hyphens_allocs = (C._Bool)(x.ReplaceNumericHyphens), cgoAllocsUnknown
	allocs23955150.Borrow(creplace_numeric_hyphens_allocs)

	var cdelete_numeric_hyphens_allocs *cgoAllocMap
	ref23955150.delete_numeric_hyphens, cdelete_numeric_hyphens_allocs = (C._Bool)(x.DeleteNumericHyphens), cgoAllocsUnknown
	allocs23955150.Borrow(cdelete_numeric_hyphens_allocs)

	var csplit_alpha_from_numeric_allocs *cgoAllocMap
	ref23955150.split_alpha_from_numeric, csplit_alpha_from_numeric_allocs = (C._Bool)(x.SplitAlphaFromNumeric), cgoAllocsUnknown
	allocs23955150.Borrow(csplit_alpha_from_numeric_allocs)

	var creplace_word_hyphens_allocs *cgoAllocMap
	ref23955150.replace_word_hyphens, creplace_word_hyphens_allocs = (C._Bool)(x.ReplaceWordHyphens), cgoAllocsUnknown
	allocs23955150.Borrow(creplace_word_hyphens_allocs)

	var cdelete_word_hyphens_allocs *cgoAllocMap
	ref23955150.delete_word_hyphens, cdelete_word_hyphens_allocs = (C._Bool)(x.DeleteWordHyphens), cgoAllocsUnknown
	allocs23955150.Borrow(cdelete_word_hyphens_allocs)

	var cdelete_final_periods_allocs *cgoAllocMap
	ref23955150.delete_final_periods, cdelete_final_periods_allocs = (C._Bool)(x.DeleteFinalPeriods), cgoAllocsUnknown
	allocs23955150.Borrow(cdelete_final_periods_allocs)

	var cdelete_acronym_periods_allocs *cgoAllocMap
	ref23955150.delete_acronym_periods, cdelete_acronym_periods_allocs = (C._Bool)(x.DeleteAcronymPeriods), cgoAllocsUnknown
	allocs23955150.Borrow(cdelete_acronym_periods_allocs)

	var cdrop_english_possessives_allocs *cgoAllocMap
	ref23955150.drop_english_possessives, cdrop_english_possessives_allocs = (C._Bool)(x.DropEnglishPossessives), cgoAllocsUnknown
	allocs23955150.Borrow(cdrop_english_possessives_allocs)

	var cdelete_apostrophes_allocs *cgoAllocMap
	ref23955150.delete_apostrophes, cdelete_apostrophes_allocs = (C._Bool)(x.DeleteApostrophes), cgoAllocsUnknown
	allocs23955150.Borrow(cdelete_apostrophes_allocs)

	var cexpand_numex_allocs *cgoAllocMap
	ref23955150.expand_numex, cexpand_numex_allocs = (C._Bool)(x.ExpandNumex), cgoAllocsUnknown
	allocs23955150.Borrow(cexpand_numex_allocs)

	var croman_numerals_allocs *cgoAllocMap
	ref23955150.roman_numerals, croman_numerals_allocs = (C._Bool)(x.RomanNumerals), cgoAllocsUnknown
	allocs23955150.Borrow(croman_numerals_allocs)

	x.ref23955150 = ref23955150
	x.allocs23955150 = allocs23955150
	return ref23955150, allocs23955150

}

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x *NormalizeOptions) PassValue() (C.normalize_options_t, *cgoAllocMap) {
	if x == nil {
		x = NewNormalizeOptionsRef(nil)
	} else if x.ref23955150 != nil {
		return *x.ref23955150, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *NormalizeOptions) Deref() {
	if x.ref23955150 == nil {
		return
	}
	packSSByte(x.Languages, x.ref23955150.languages)
	x.NumLanguages = (int32)(x.ref23955150.num_languages)
	x.AddressComponents = (uint16)(x.ref23955150.address_components)
	x.LatinAscii = (bool)(x.ref23955150.latin_ascii)
	x.Transliterate = (bool)(x.ref23955150.transliterate)
	x.StripAccents = (bool)(x.ref23955150.strip_accents)
	x.Decompose = (bool)(x.ref23955150.decompose)
	x.Lowercase = (bool)(x.ref23955150.lowercase)
	x.TrimString = (bool)(x.ref23955150.trim_string)
	x.DropParentheticals = (bool)(x.ref23955150.drop_parentheticals)
	x.ReplaceNumericHyphens = (bool)(x.ref23955150.replace_numeric_hyphens)
	x.DeleteNumericHyphens = (bool)(x.ref23955150.delete_numeric_hyphens)
	x.SplitAlphaFromNumeric = (bool)(x.ref23955150.split_alpha_from_numeric)
	x.ReplaceWordHyphens = (bool)(x.ref23955150.replace_word_hyphens)
	x.DeleteWordHyphens = (bool)(x.ref23955150.delete_word_hyphens)
	x.DeleteFinalPeriods = (bool)(x.ref23955150.delete_final_periods)
	x.DeleteAcronymPeriods = (bool)(x.ref23955150.delete_acronym_periods)
	x.DropEnglishPossessives = (bool)(x.ref23955150.drop_english_possessives)
	x.DeleteApostrophes = (bool)(x.ref23955150.delete_apostrophes)
	x.ExpandNumex = (bool)(x.ref23955150.expand_numex)
	x.RomanNumerals = (bool)(x.ref23955150.roman_numerals)
}

// allocAddressParserResponseMemory allocates memory for type C.address_parser_response_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAddressParserResponseMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAddressParserResponseValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAddressParserResponseValue = unsafe.Sizeof([1]C.address_parser_response_t{})

// Ref returns a reference.
func (x *AddressParserResponse) Ref() *C.address_parser_response_t {
	if x == nil {
		return nil
	}
	return x.ref50c6200b
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *AddressParserResponse) Free() {
	if x != nil && x.allocs50c6200b != nil {
		x.allocs50c6200b.(*cgoAllocMap).Free()
		x.ref50c6200b = nil
	}
}

// NewAddressParserResponseRef initialises a new struct holding the reference to the originaitng C struct.
func NewAddressParserResponseRef(ref *C.address_parser_response_t) *AddressParserResponse {
	if ref == nil {
		return nil
	}
	obj := new(AddressParserResponse)
	obj.ref50c6200b = ref
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *AddressParserResponse) PassRef() (*C.address_parser_response_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref50c6200b != nil {
		return x.ref50c6200b, nil
	}
	mem50c6200b := allocAddressParserResponseMemory(1)
	ref50c6200b := (*C.address_parser_response_t)(mem50c6200b)
	allocs50c6200b := new(cgoAllocMap)
	var cnum_components_allocs *cgoAllocMap
	ref50c6200b.num_components, cnum_components_allocs = (C.size_t)(x.NumComponents), cgoAllocsUnknown
	allocs50c6200b.Borrow(cnum_components_allocs)

	var ccomponents_allocs *cgoAllocMap
	ref50c6200b.components, ccomponents_allocs = unpackSSByte(x.Components)
	allocs50c6200b.Borrow(ccomponents_allocs)

	var clabels_allocs *cgoAllocMap
	ref50c6200b.labels, clabels_allocs = unpackSSByte(x.Labels)
	allocs50c6200b.Borrow(clabels_allocs)

	x.ref50c6200b = ref50c6200b
	x.allocs50c6200b = allocs50c6200b
	return ref50c6200b, allocs50c6200b

}

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x *AddressParserResponse) PassValue() (C.address_parser_response_t, *cgoAllocMap) {
	if x == nil {
		x = NewAddressParserResponseRef(nil)
	} else if x.ref50c6200b != nil {
		return *x.ref50c6200b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *AddressParserResponse) Deref() {
	if x.ref50c6200b == nil {
		return
	}
	x.NumComponents = (uint)(x.ref50c6200b.num_components)
	packSSByte(x.Components, x.ref50c6200b.components)
	packSSByte(x.Labels, x.ref50c6200b.labels)
}

// allocAddressParserOptionsMemory allocates memory for type C.address_parser_options_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAddressParserOptionsMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAddressParserOptionsValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAddressParserOptionsValue = unsafe.Sizeof([1]C.address_parser_options_t{})

// Ref returns a reference.
func (x *AddressParserOptions) Ref() *C.address_parser_options_t {
	if x == nil {
		return nil
	}
	return x.refcd8497af
}

// Free cleanups the memory using the free stdlib function on C side.
// Does nothing if object has no pointer.
func (x *AddressParserOptions) Free() {
	if x != nil && x.allocscd8497af != nil {
		x.allocscd8497af.(*cgoAllocMap).Free()
		x.refcd8497af = nil
	}
}

// NewAddressParserOptionsRef initialises a new struct holding the reference to the originaitng C struct.
func NewAddressParserOptionsRef(ref *C.address_parser_options_t) *AddressParserOptions {
	if ref == nil {
		return nil
	}
	obj := new(AddressParserOptions)
	obj.refcd8497af = ref
	return obj
}

// PassRef returns a reference and creates new C object if no refernce yet.
func (x *AddressParserOptions) PassRef() (*C.address_parser_options_t, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refcd8497af != nil {
		return x.refcd8497af, nil
	}
	memcd8497af := allocAddressParserOptionsMemory(1)
	refcd8497af := (*C.address_parser_options_t)(memcd8497af)
	allocscd8497af := new(cgoAllocMap)
	var clanguage_allocs *cgoAllocMap
	refcd8497af.language, clanguage_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Language)).Data)), cgoAllocsUnknown
	allocscd8497af.Borrow(clanguage_allocs)

	var ccountry_allocs *cgoAllocMap
	refcd8497af.country, ccountry_allocs = (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&x.Country)).Data)), cgoAllocsUnknown
	allocscd8497af.Borrow(ccountry_allocs)

	x.refcd8497af = refcd8497af
	x.allocscd8497af = allocscd8497af
	return refcd8497af, allocscd8497af

}

// PassValue creates a new C object if no refernce yet and returns the dereferenced value.
func (x *AddressParserOptions) PassValue() (C.address_parser_options_t, *cgoAllocMap) {
	if x == nil {
		x = NewAddressParserOptionsRef(nil)
	} else if x.refcd8497af != nil {
		return *x.refcd8497af, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref reads the internal fields of struct from its C pointer.
func (x *AddressParserOptions) Deref() {
	if x.refcd8497af == nil {
		return
	}
	hxff73280 := (*sliceHeader)(unsafe.Pointer(&x.Language))
	hxff73280.Data = uintptr(unsafe.Pointer(x.refcd8497af.language))
	hxff73280.Cap = 0x7fffffff
	// hxff73280.Len = ?

	hxfa9955c := (*sliceHeader)(unsafe.Pointer(&x.Country))
	hxfa9955c.Data = uintptr(unsafe.Pointer(x.refcd8497af.country))
	hxfa9955c.Cap = 0x7fffffff
	// hxfa9955c.Len = ?

}

// unpackArgSSByte transforms a sliced Go data structure into plain C format.
func unpackArgSSByte(x [][]byte) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.char)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}
