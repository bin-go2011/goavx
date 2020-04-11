package cv

/*
#include "goavx/cv/videoio.h"
*/
import "C"
import "fmt"

type VideoCapture struct {
	handle C.VideoCapturePtr
}

func CvNewVideoCapture() (*VideoCapture, error) {
	cap := C._cv_new_videocapture()
	if cap == nil {
		err := fmt.Errorf("failed to initialize VideoCapture object")
		return nil, err
	}
	return &VideoCapture{
		handle: cap,
	}, nil
}

func CvVideoCaptureOpenDevice(cap *VideoCapture, device int) error {
	ret := C._cv_videocapture_opendevice((C.VideoCapturePtr)(cap.handle), (C.int)(device))
	if int(ret) < 0 {
		err := fmt.Errorf("failed to open video capture device %d", device)
		return err
	}

	return nil
}

func CvReleaseVideoCapture(cap *VideoCapture) {
	C._cv_release_videocapture((C.VideoCapturePtr)(cap.handle))
}

func CvVideoCaptureRead(cap *VideoCapture, mat Mat) error {
	ret := C._cv_videocapture_read((C.VideoCapturePtr)(cap.handle), (C.MatPtr)(mat))
	if int(ret) < 0 {
		err := fmt.Errorf("failed to read from video capture device")
		return err
	}

	return nil
}
