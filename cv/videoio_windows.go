package cv

import (
	"golang.org/x/sys/windows"
)

type VideoCapture struct {
	handle uintptr
}

var (
	cvNewVideoCaptureProc,
	cvVideoCaptureOpenDeviceProc *windows.Proc
)

// func CvNewVideoCapture() (*VideoCapture, error) {
// 	if cvNewVideoCaptureProc == nil {
// 		cvNewVideoCaptureProc = goavx.LoadedDLL.MustFindProc("_cv_new_videocapture")
// 	}

// 	r1, _, err := cvNewVideoCaptureProc.Call()

// 	if r1 == 0 {
// 		return nil, err
// 	}

// 	cap := &VideoCapture{
// 		handle: r1,
// 	}

// 	return cap, nil
// }

// func CvVideoCaptureOpenDevice(vc *VideoCapture, device int) {
// 	if cvVideoCaptureOpenDeviceProc == nil {
// 		cvVideoCaptureOpenDeviceProc = goavx.LoadedDLL.MustFindProc("_cv_videocapture_open_device")
// 	}

// 	r1, _, _ := cvVideoCaptureOpenDeviceProc.Call(uintptr(vc.handle), uintptr(device))

// 	fmt.Println(r1)
// }
