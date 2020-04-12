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

func CvImread(file string, flags int, mat Mat) {
	if cvImreadProc == nil {
		cvImreadProc = goavx.LoadedDLL.MustFindProc("_cv_imread")
	}

	f, _ := syscall.BytePtrFromString(file)
	cvImreadProc.Call(uintptr(unsafe.Pointer(f)), uintptr(flags), mat.handle)
}
