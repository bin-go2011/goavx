package cv

import (
	"syscall"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

const (
	CAP_PROP_FRAME_WIDTH  = 3
	CAP_PROP_FRAME_HEIGHT = 4
	CAP_PROP_FPS          = 5
)

type VideoCapture struct {
	handle uintptr
}

var (
	cvNewVideoCaptureProc,
	cvVideoCaptureOpenDeviceProc,
	cvReleaseVideoCaptureProc,
	cvVideoCaptureReadProc,
	cvVideoCaptureOpenFileProc,
	cvVideoCaptureIsOpenedProc,
	cvVideoCaptureGetProc *windows.Proc
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

func cvVideoCaptureIsOpened(cap *VideoCapture) bool {
	if cvVideoCaptureIsOpenedProc == nil {
		cvVideoCaptureIsOpenedProc = goavx.LoadedDLL.MustFindProc("_cv_videocapture_is_opened")
	}

	r1, _, _ := cvVideoCaptureIsOpenedProc.Call(uintptr(cap.handle))
	if int(r1) > 0 {
		return true
	} else {
		return false
	}
}

func cvVideoCaptureGet(cap *VideoCapture, propId int) int32 {
	if cvVideoCaptureGetProc == nil {
		cvVideoCaptureGetProc = goavx.LoadedDLL.MustFindProc("_cv_videocapture_get")
	}

	r1, _, _ := cvVideoCaptureGetProc.Call(uintptr(cap.handle), uintptr(propId))
	return int32(r1)
}
