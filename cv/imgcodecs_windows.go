package cv

import (
	"syscall"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

var (
	cvImreadProc *windows.Proc
)

func cvImread(file string, flags int, mat *Mat) error {
	if cvImreadProc == nil {
		cvImreadProc = goavx.LoadedDLL.MustFindProc("_cv_imread")
	}

	f, _ := syscall.BytePtrFromString(file)
	r1, _, err := cvImreadProc.Call(uintptr(unsafe.Pointer(f)), uintptr(flags), mat.handle)

	if int32(r1) < 0 {
		return err
	}
	return nil
}
