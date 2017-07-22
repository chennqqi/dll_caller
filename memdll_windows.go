package godll

/*
#include "MemoryModule/MemoryModule.c"

#ifndef UNICODE
const CHAR* ToChar(const char* p)
{
	return (const CHAR*)p;
}
#endif

*/
import "C"

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

type MemDll struct {
	mem        []byte
	dllHandler C.HMEMORYMODULE
	funcProcs  map[string]uintptr
}

var ErrNotFound = errors.New("Not Found dll resource")

func NewMemDll(dllMem []byte) (dll Dll, err error) {
	newDll := new(MemDll)
	if newDll.funcProcs == nil {
		newDll.funcProcs = make(map[string]uintptr)
	}

	//do deep copy to void memory changed
	newDll.mem = make([]byte, len(dllMem))
	copy(newDll.mem, dllMem)
	if err = newDll.LoadLibrary(""); err != nil {
		return
	}

	return newDll, nil
}

func (p *MemDll) LoadLibrary(fileName string) error {
	if handler, e := C.MemoryLoadLibrary(unsafe.Pointer(&p.mem[0]), C.size_t(len(p.mem))); e != nil {
		return e
	} else {
		p.dllHandler = handler
	}
	return nil
}

func (p *MemDll) FreeLibrary() error {
	if p.IsDllLoaded() {
		defer func() {
			p.dllHandler = nil
		}()
		C.MemoryFreeLibrary(p.dllHandler)
		return nil
	}
	return nil
}

func (p *MemDll) IsDllLoaded() bool {
	if uintptr(p.dllHandler) == 0 {
		return false
	}
	return true
}

func (p *MemDll) InitalFunctions(funcNames ...string) error {
	if funcNames == nil {
		return nil
	}

	if !p.IsDllLoaded() {
		return errors.New("dll should loaded befor inital functions")
	}

	if p.funcProcs == nil {
		p.funcProcs = make(map[string]uintptr)
	}

	for _, funcName := range funcNames {
		funcName = strings.TrimSpace(funcName)
		if funcName == "" {
			return errors.New("function name could not be empty")
		}
		pCFuncName := C.CString(funcName)
		defer C.free(unsafe.Pointer(pCFuncName))
		if proc, e := C.MemoryGetProcAddress(p.dllHandler, C.ToChar(pCFuncName)); e != nil {
			return e
		} else {
			p.funcProcs[funcName] = uintptr(unsafe.Pointer(proc))
		}
	}
	return nil
}

func (p *MemDll) Call(funcName string, funcParams ...interface{}) (result FuncCallResult, err error) {
	var lenParam uintptr = uintptr(len(funcParams))

	if p.funcProcs == nil {
		err = errors.New("function address not initaled")
		return
	}

	var funcAddress uintptr
	if addr, exist := p.funcProcs[funcName]; !exist {
		err = errors.New("function address not exist")
		return
	} else {
		funcAddress = addr
	}

	var r1, r2 uintptr
	var errno syscall.Errno

	var a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15 uintptr

	for parmIndex, param := range funcParams {
		var vPtr uintptr = 0

		switch v := param.(type) {
		case string:
			vPtr = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(v)))
		case *string:
			vPtr = uintptr(unsafe.Pointer(syscall.StringBytePtr(*v)))
		case bool:
			vPtr = uintptr(unsafe.Pointer(&v))
		case int:
			vPtr = uintptr(v)
		case int8:
			vPtr = uintptr(v)
		case uint8:
			vPtr = uintptr(v)
		case *uint8:
			vPtr = uintptr(unsafe.Pointer(v))
		case int16:
			vPtr = uintptr(v)
		case uint16:
			vPtr = uintptr(v)
		case *uint16:
			vPtr = uintptr(unsafe.Pointer(v))
		case int32:
			vPtr = uintptr(v)
		case *int32:
			vPtr = uintptr(unsafe.Pointer(v))
		case uint32:
			vPtr = uintptr(v)
		case *uint32:
			vPtr = uintptr(unsafe.Pointer(v))
		case int64:
			vPtr = uintptr(v)
		case *int64:
			vPtr = uintptr(unsafe.Pointer(v))
		case uint64:
			vPtr = uintptr(v)
		case *uint64:
			vPtr = uintptr(unsafe.Pointer(v))
		case *int:
			vPtr = uintptr(unsafe.Pointer(v))
		case float32:
			vPtr = uintptr(v)
		case float64:
			vPtr = uintptr(v)
		case []byte:
			vPtr = uintptr(unsafe.Pointer(&v[0]))
		case uintptr:
			ptr, _ := param.(uintptr)
			vPtr = ptr
		default:
			err = fmt.Errorf("unsupport convert type %v to uintptr", reflect.TypeOf(param))
			return
		}

		switch parmIndex + 1 {
		case 1:
			a1 = vPtr
		case 2:
			a2 = vPtr
		case 3:
			a3 = vPtr
		case 4:
			a4 = vPtr
		case 5:
			a5 = vPtr
		case 6:
			a6 = vPtr
		case 7:
			a7 = vPtr
		case 8:
			a8 = vPtr
		case 9:
			a9 = vPtr
		case 10:
			a10 = vPtr
		case 11:
			a11 = vPtr
		case 12:
			a12 = vPtr
		case 13:
			a13 = vPtr
		case 14:
			a14 = vPtr
		case 15:
			a15 = vPtr
		}
	}

	r1, r2, errno = syscall.Syscall15(funcAddress, lenParam, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15)
	result.Ret1 = r1
	result.Ret2 = r2
	result.Errno = errno

	return
}
