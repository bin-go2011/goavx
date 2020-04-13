package cv

import (
	"fmt"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

type Mat struct {
	handle uintptr
}

var (
	cvVersionProc,
	cvNewMatProc,
	cvReleaseMatProc *windows.Proc
)

func CvVersion() string {
	if cvVersionProc == nil {
		cvVersionProc = goavx.LoadedDLL.MustFindProc("_cv_version")
	}
	r1, _, _ := cvVersionProc.Call()

	version := int32(r1)
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}

func CvNewMat() (*Mat, error) {
	if cvNewMatProc == nil {
		cvNewMatProc = goavx.LoadedDLL.MustFindProc("_cv_new_mat")
	}

	r1, _, _ := cvNewMatProc.Call()

	if r1 == 0 {
		err := fmt.Errorf("failed to new Mat object")
		return nil, err
	}
	return &Mat{
		handle: r1,
	}, nil

}

func CvReleaseMat(mat *Mat) {
	if cvReleaseMatProc == nil {
		cvReleaseMatProc = goavx.LoadedDLL.MustFindProc("_cv_release_mat")
	}

	cvReleaseMatProc.Call(uintptr(mat.handle))
}
