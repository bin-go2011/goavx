package cv

import (
	"syscall"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

var (
	cvNamedWindowPro,
	cvWaitKeyProc,
	cvImshowProc *windows.Proc
)

func CvNamedWindow(name string, flags int) {
	if cvNamedWindowPro == nil {
		cvNamedWindowPro = goavx.LoadedDLL.MustFindProc("cv_named_window")
	}

	winname, _ := syscall.BytePtrFromString(name)

	cvNamedWindowPro.Call(uintptr(unsafe.Pointer(winname)), uintptr(flags))
}

func CvWaitKey(delay int) {
	if cvWaitKeyProc == nil {
		cvWaitKeyProc = goavx.LoadedDLL.MustFindProc("cv_wait_key")
	}

	cvWaitKeyProc.Call(uintptr(delay))

}

func CvImshow(name string, mat Mat) {
	if cvImshowProc == nil {
		cvImshowProc = goavx.LoadedDLL.MustFindProc("cv_imshow")
	}

	winname, _ := syscall.BytePtrFromString(name)

	cvImshowProc.Call(uintptr(unsafe.Pointer(winname)), mat.handle)
}
