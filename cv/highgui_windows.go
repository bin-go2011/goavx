package cv

import (
	"syscall"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

var (
	cvNamedWindowPro,
	cvWaitKeyProc *windows.Proc
)

func CvNamedWindow(name string, flags int) {
	if cvNamedWindowPro == nil {
		cvNamedWindowPro = goavx.LoadedDLL.MustFindProc("cv_named_window")
	}

	n, _ := syscall.BytePtrFromString(name)

	cvNamedWindowPro.Call(uintptr(unsafe.Pointer(n)), uintptr(flags))
}

func CvWaitKey(delay int) {
	if cvWaitKeyProc == nil {
		cvWaitKeyProc = goavx.LoadedDLL.MustFindProc("cv_wait_key")
	}

	cvWaitKeyProc.Call(uintptr(delay))

}
