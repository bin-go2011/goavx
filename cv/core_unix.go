package cv

/*
#include "goavx/cv/core.h"
*/
import "C"
import (
	"fmt"
)

type (
	Mat C.MatPtr
)

func CvVersion() string {
	version := C.cv_version()
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}

func CvNewMat() Mat {
	return Mat(C.cv_new_mat())
}
