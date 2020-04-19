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
	cvMatColsProc,
	cvMatChannelsProc *windows.Proc
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

func cvMatShape(mat *Mat) (int, int, int) {
	if cvMatColsProc == nil {
		cvMatColsProc = goavx.LoadedDLL.MustFindProc("_cv_mat_cols")
	}

	if cvMatRowsProc == nil {
		cvMatRowsProc = goavx.LoadedDLL.MustFindProc("_cv_mat_rows")
	}

	if cvMatChannelsProc == nil {
		cvMatChannelsProc = goavx.LoadedDLL.MustFindProc("_cv_mat_channels")
	}

	rows, _, _ := cvMatRowsProc.Call(uintptr(mat.handle))
	cols, _, _ := cvMatColsProc.Call(uintptr(mat.handle))
	chs, _, _ := cvMatChannelsProc.Call(uintptr(mat.handle))

	return int(rows), int(cols), int(chs)
}
