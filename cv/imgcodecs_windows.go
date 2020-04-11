package cv

import (
	"syscall"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

const (
	IMREAD_UNCHANGED = -1 //!< If set, return the loaded image as is (with alpha channel, otherwise it gets cropped).
	IMREAD_GRAYSCALE = 0  //!< If set, always convert image to the single channel grayscale image (codec internal conversion).
	IMREAD_COLOR     = 1  //!< If set, always convert image to the 3 channel BGR color image.
	IMREAD_ANYDEPTH  = 2  //!< If set, return 16-bit/32-bit image when the input has the corresponding depth, otherwise convert it to 8-bit.
	IMREAD_ANYCOLOR  = 4  //!< If set, the image is read in any possible color format.
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
