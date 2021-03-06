package cv

import (
	"fmt"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

type Mat struct {
	handle uintptr
}

var (
	cvVersionProc,
	cvNewMatProc,
	cvNewMatWithProc,
	cvNewMatFromSizeProc,
	cvReleaseMatProc,
	cvMatChannelsProc,
	cvMatSizeProc,
	cvMatCopyToProc,
	cvMatSetToProc *windows.Proc
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

func cvNewMatWith(width int, height int, mt int) (*Mat, error) {
	if cvNewMatWithProc == nil {
		cvNewMatWithProc = goavx.LoadedDLL.MustFindProc("_cv_new_mat_with")
	}

	r1, _, _ := cvNewMatWithProc.Call(uintptr(width), uintptr(height), uintptr(mt))

	if r1 == 0 {
		err := fmt.Errorf("failed to new Mat object")
		return nil, err
	}
	return &Mat{
		handle: r1,
	}, nil

}

func cvNewMatFromSize(size Size, mt int) (*Mat, error) {
	if cvNewMatFromSizeProc == nil {
		cvNewMatFromSizeProc = goavx.LoadedDLL.MustFindProc("_cv_new_mat_from_size")
	}

	r1, _, _ := cvNewMatFromSizeProc.Call(uintptr(unsafe.Pointer(&size)), uintptr(mt))

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

func cvMatShape(mat *Mat) (rows int32, cols int32, channels int32) {
	if cvMatChannelsProc == nil {
		cvMatChannelsProc = goavx.LoadedDLL.MustFindProc("_cv_mat_channels")
	}

	rows = *(*int32)(unsafe.Pointer(uintptr(mat.handle) + offsetOfMatRows))
	cols = *(*int32)(unsafe.Pointer(uintptr(mat.handle) + offsetOfMatCols))
	chs, _, _ := cvMatChannelsProc.Call(uintptr(mat.handle))

	return rows, cols, int32(chs)
}

func cvMatSize(mat *Mat) (width int32, height int32) {
	if cvMatSizeProc == nil {
		cvMatSizeProc = goavx.LoadedDLL.MustFindProc("_cv_mat_size")
	}

	cvMatSizeProc.Call(uintptr(mat.handle), uintptr(unsafe.Pointer(&width)), uintptr(unsafe.Pointer(&height)))
	return
}

func cvMatCopyTo(src *Mat, dst *Mat, mask *Mat) {
	if cvMatCopyToProc == nil {
		cvMatCopyToProc = goavx.LoadedDLL.MustFindProc("_cv_mat_copy_to")
	}

	cvMatCopyToProc.Call(uintptr(src.handle), uintptr(dst.handle), uintptr(mask.handle))
}

func cvMatSetTo(mat *Mat, value int) {
	if cvMatSetToProc == nil {
		cvMatSetToProc = goavx.LoadedDLL.MustFindProc("_cv_mat_set_to")
	}

	cvMatSetToProc.Call(uintptr(mat.handle), uintptr(value))
}
