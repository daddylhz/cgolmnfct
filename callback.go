package cgolmnfct

import "C"

import (
	_ "fmt"
	"github.com/mattn/go-pointer"
	"unsafe"
)

type ConntrackInstance struct {
	Instancensn string
	Pod         string
}
type Callback struct {
	Func func(int, unsafe.Pointer, *map[string][]ConntrackInstance)
	Data *map[string][]ConntrackInstance
}

//export goCallbackFunc
func goCallbackFunc(msg int, ct unsafe.Pointer, data unsafe.Pointer) {
	cb := pointer.Restore(data).(*Callback)
	if cb == nil {
		return
	}
	cb.Func(msg, ct, cb.Data)
}
