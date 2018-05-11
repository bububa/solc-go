package solc

/*
#cgo LDFLAGS: -L${SRCDIR}/lib/ -lsolc -ljsoncpp -ldevcore -lsolidity -levmasm -lboost_filesystem -lboost_regex -lstdc++
#cgo CFLAGS: -I${SRCDIR}/include/

#include <stdlib.h>
#include <libsolc.h>

*/
import "C"

import "unsafe"

func License() string {
	licenseC := C.license()
	return C.GoString(licenseC)
}

func Version() string {
	versionC := C.version()
	return C.GoString(versionC)
}

func CompileJSON(input string, optimize bool) string {
	_input := C.CString(input)
	defer C.free(unsafe.Pointer(_input))
	ret := C.compileJSON(_input, C._Bool(optimize))
	return C.GoString(ret)
}

func CompileJSONMulti(input string, optimize bool) string {
	_input := C.CString(input)
	defer C.free(unsafe.Pointer(_input))
	ret := C.compileJSONMulti(_input, C._Bool(optimize))
	return C.GoString(ret)
}

func CompileStandard(input string) string {
	_input := C.CString(input)
	defer C.free(unsafe.Pointer(_input))
	ret := C.compileStandard(_input, nil)
	return C.GoString(ret)
}
