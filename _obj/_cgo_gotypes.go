// Created by cgo - DO NOT EDIT

package main

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})

//go:cgo_import_static _cgo_4dd2ab524fb7_Cfunc_SayHello
//go:linkname __cgofn__cgo_4dd2ab524fb7_Cfunc_SayHello _cgo_4dd2ab524fb7_Cfunc_SayHello
var __cgofn__cgo_4dd2ab524fb7_Cfunc_SayHello byte
var _cgo_4dd2ab524fb7_Cfunc_SayHello = unsafe.Pointer(&__cgofn__cgo_4dd2ab524fb7_Cfunc_SayHello)

//go:cgo_unsafe_args
func _Cfunc_SayHello(p0 string) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_4dd2ab524fb7_Cfunc_SayHello, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_export_dynamic SayHello
//go:linkname _cgoexp_4dd2ab524fb7_SayHello _cgoexp_4dd2ab524fb7_SayHello
//go:cgo_export_static _cgoexp_4dd2ab524fb7_SayHello
//go:nosplit
//go:norace
func _cgoexp_4dd2ab524fb7_SayHello(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_4dd2ab524fb7_SayHello
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_4dd2ab524fb7_SayHello(p0 string) {
	SayHello(p0)
}
