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
	cvImshowProc,
	cvDestroyWindowProc *windows.Proc
)

func cvNamedWindow(name string, flags int) {
	if cvNamedWindowPro == nil {
		cvNamedWindowPro = goavx.LoadedDLL.MustFindProc("_cv_named_window")
	}

	winname, _ := syscall.BytePtrFromString(name)

	cvNamedWindowPro.Call(uintptr(unsafe.Pointer(winname)), uintptr(flags))
}

func cvWaitKey(delay int) int8 {
	if cvWaitKeyProc == nil {
		cvWaitKeyProc = goavx.LoadedDLL.MustFindProc("_cv_wait_key")
	}

	r1, _, _ := cvWaitKeyProc.Call(uintptr(delay))
	return int8(r1)
}

func cvImshow(name string, img *Mat) {
	if cvImshowProc == nil {
		cvImshowProc = goavx.LoadedDLL.MustFindProc("_cv_imshow")
	}

	winname, _ := syscall.BytePtrFromString(name)

	cvImshowProc.Call(uintptr(unsafe.Pointer(winname)), img.handle)
}

func cvDestroyWindow(name string) {
	if cvDestroyWindowProc == nil {
		cvDestroyWindowProc = goavx.LoadedDLL.MustFindProc("_cv_destroy_window")
	}

	winname, _ := syscall.BytePtrFromString(name)
	cvDestroyWindowProc.Call(uintptr(unsafe.Pointer(winname)))
}
