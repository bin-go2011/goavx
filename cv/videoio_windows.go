package cv

import (
	"syscall"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

type VideoCapture struct {
	handle uintptr
}

var (
	cvNewVideoCaptureProc,
	cvVideoCaptureOpenDeviceProc,
	cvReleaseVideoCaptureProc,
	cvVideoCaptureReadProc,
	cvVideoCaptureOpenFileProc *windows.Proc
)

func cvNewVideoCapture() (*VideoCapture, error) {
	if cvNewVideoCaptureProc == nil {
		cvNewVideoCaptureProc = goavx.LoadedDLL.MustFindProc("_cv_new_videocapture")
	}

	r1, _, err := cvNewVideoCaptureProc.Call()

	if r1 == 0 {
		return nil, err
	}

	return &VideoCapture{
		handle: r1,
	}, nil
}

func cvVideoCaptureOpenDevice(cap *VideoCapture, device int) error {
	if cvVideoCaptureOpenDeviceProc == nil {
		cvVideoCaptureOpenDeviceProc = goavx.LoadedDLL.MustFindProc("_cv_videocapture_open_device")
	}

	r1, _, err := cvVideoCaptureOpenDeviceProc.Call(uintptr(cap.handle), uintptr(device))

	if int(r1) < 0 {
		return err
	}

	return nil
}

func cvReleaseVideoCapture(cap *VideoCapture) {
	if cvReleaseVideoCaptureProc == nil {
		cvReleaseVideoCaptureProc = goavx.LoadedDLL.MustFindProc("_cv_release_videocapture")
	}

	cvReleaseVideoCaptureProc.Call(uintptr(cap.handle))
}

func cvVideoCaptureRead(cap *VideoCapture, mat *Mat) error {
	if cvVideoCaptureReadProc == nil {
		cvVideoCaptureReadProc = goavx.LoadedDLL.MustFindProc("_cv_videocapture_read")
	}

	r1, _, err := cvVideoCaptureReadProc.Call(uintptr(cap.handle), uintptr(mat.handle))
	if int(r1) == 0 {
		return err
	}

	return nil
}

func cvVideoCaptureOpenFile(cap *VideoCapture, filename string) error {
	if cvVideoCaptureOpenFileProc == nil {
		cvVideoCaptureOpenFileProc = goavx.LoadedDLL.MustFindProc("_cv_videocapture_open_file")
	}

	f, _ := syscall.BytePtrFromString(filename)
	r1, _, err := cvVideoCaptureOpenFileProc.Call(uintptr(cap.handle), uintptr(unsafe.Pointer(f)))
	if int(r1) < 0 {
		return err
	}

	return nil

}
