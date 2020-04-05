package cv

import (
	"syscall"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

const (
	WINDOW_NORMAL   = 0x00000000 //!< the user can resize the window (no constraint) / also use to switch a fullscreen window to a normal size.
	WINDOW_AUTOSIZE = 0x00000001 //!< the user cannot resize the window, the size is constrainted by the image displayed.
	WINDOW_OPENGL   = 0x00001000 //!< window with opengl support.
)

var (
	cvNamedWindowPro,
	cvWaitKeyProc,
	cvImshowProc,
	cvDestroyWindowProc *windows.Proc
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

func CvImshow(name string, img Mat) {
	if cvImshowProc == nil {
		cvImshowProc = goavx.LoadedDLL.MustFindProc("cv_imshow")
	}

	winname, _ := syscall.BytePtrFromString(name)

	cvImshowProc.Call(uintptr(unsafe.Pointer(winname)), img.handle)
}

func CvDestroyWindow(name string) {
	if cvDestroyWindowProc == nil {
		cvDestroyWindowProc = goavx.LoadedDLL.MustFindProc("cv_destroy_window")
	}

	winname, _ := syscall.BytePtrFromString(name)
	cvDestroyWindowProc.Call(uintptr(unsafe.Pointer(winname)))
}
