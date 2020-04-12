package cv

/*
#include "goavx/cv/imgproc.h"
*/
import "C"

func CvGaussianBlur(src *Mat, dst *Mat, ksizeX int, ksizeY int, sigmaX float64, sigmaY float64, borderType int) {
	C._cv_gaussian_blur((C.MatPtr)(src.handle), (C.MatPtr)(dst.handle), (C.int)(ksizeX), (C.int)(ksizeY), (C.double)(sigmaX), (C.double)(sigmaY), (C.int)(borderType))
}
