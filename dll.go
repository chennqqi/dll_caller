package godll

import (
	"syscall"
)

type Dll interface {
	Call(funcName string, funcParams ...interface{}) (result FuncCallResult, err error)
	FreeLibrary() error
	InitalFunctions(funcNames ...string) error
	IsDllLoaded() bool
	LoadLibrary(fileName string) error
}

type FuncCallResult struct {
	Ret1  uintptr
	Ret2  uintptr
	Errno syscall.Errno
}
