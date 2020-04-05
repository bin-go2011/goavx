package cv

import (
	"fmt"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

type VideoCapture struct {
	handle uintptr
}

var (
	cvNewVideoCaptureProc,
	cvVideoCaptureOpenDeviceProc *windows.Proc
)

func CvNewVideoCapture() (*VideoCapture, error) {
	if cvNewVideoCaptureProc == nil {
		cvNewVideoCaptureProc = goavx.LoadedDLL.MustFindProc("cv_new_videocapture")
	}

	r1, _, err := cvNewVideoCaptureProc.Call()

	if r1 == 0 {
		return nil, err
	}

	cap := &VideoCapture{
		handle: r1,
	}

	return cap, nil
}

func CvVideoCaptureOpenDevice(vc *VideoCapture, device int) {
	if cvVideoCaptureOpenDeviceProc == nil {
		cvVideoCaptureOpenDeviceProc = goavx.LoadedDLL.MustFindProc("cv_videocapture_opendevice")
	}

	r1, _, _ := cvVideoCaptureOpenDeviceProc.Call(uintptr(vc.handle), uintptr(device))

	fmt.Println(r1)
}
