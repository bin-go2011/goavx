package cv

/*
#include "goavx/cv/core.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Mat struct {
	handle C.MatPtr
}

func cvVersion() string {
	version := C._cv_version()
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}

func cvNewMat() (*Mat, error) {
	mat := C._cv_new_mat()
	if mat == nil {
		err := fmt.Errorf("failed to new Mat object")
		return nil, err
	}
	return &Mat{
		handle: mat,
	}, nil
}

func cvReleaseMat(mat *Mat) {
	C._cv_release_mat(C.MatPtr(mat.handle))
}

func cvMatShape(mat *Mat) (rows int32, cols int32, channels int32) {
	rows = *(*int32)(unsafe.Pointer(uintptr(mat.handle) + offsetOfMatRows))
	cols = *(*int32)(unsafe.Pointer(uintptr(mat.handle) + offsetOfMatCols))
	channels = int32(C._cv_mat_channels(C.MatPtr(mat.handle)))

	return rows, cols, channels
}
