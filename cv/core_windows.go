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
	cvReleaseMatProc,
	cvMatRowsProc,
	cvMatColsProc *windows.Proc
)

func cvVersion() string {
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

func cvNewMat() (*Mat, error) {
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

func cvReleaseMat(mat *Mat) {
	if cvReleaseMatProc == nil {
		cvReleaseMatProc = goavx.LoadedDLL.MustFindProc("_cv_release_mat")
	}

	cvReleaseMatProc.Call(uintptr(mat.handle))
}

func cvMatShape(mat *Mat) (rows int, cols int) {
	if cvMatColsProc == nil {
		cvMatColsProc = goavx.LoadedDLL.MustFindProc("_cv_mat_cols")
	}

	if cvMatRowsProc == nil {
		cvMatRowsProc = goavx.LoadedDLL.MustFindProc("_cv_mat_rows")
	}
	r1, _, _ := cvMatRowsProc.Call(uintptr(mat.handle))
	c1, _, _ := cvMatColsProc.Call(uintptr(mat.handle))

	return int(r1), int(c1)
}
