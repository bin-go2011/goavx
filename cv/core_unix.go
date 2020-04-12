package cv

/*
#include "goavx/cv/core.h"
*/
import "C"
import (
	"fmt"
)

type Mat struct {
	handle C.MatPtr
}

func CvVersion() string {
	version := C._cv_version()
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}

func CvNewMat() (*Mat, error) {
	mat := C._cv_new_mat()
	if mat == nil {
		err := fmt.Errorf("failed to new Mat object")
		return nil, err
	}
	return &Mat{
		handle: mat,
	}, nil
}

func CvReleaseMat(mat *Mat) {
	C._cv_release_mat((C.MatPtr)(mat.handle))
}
