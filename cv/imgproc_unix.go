package cv

/*
#include "goavx/cv/imgproc.h"
*/
import "C"

func cvGaussianBlur(src *Mat, dst *Mat, ksizeX int, ksizeY int, sigmaX float64, sigmaY float64, borderType int) {
	C._cv_gaussian_blur((C.MatPtr)(src.handle), (C.MatPtr)(dst.handle), (C.int)(ksizeX), (C.int)(ksizeY), (C.double)(sigmaX), (C.double)(sigmaY), (C.int)(borderType))
}

func cvPyrDown(src *Mat, dst *Mat) {
	C._cv_pyrdown((C.MatPtr)(src.handle), (C.MatPtr)(dst.handle))
}

func cvCanny(img *Mat, edges *Mat, threshold1 float64, threshold2 float64, apertureSize int, L2gradient int) {
	C._cv_canny((C.MatPtr)(img.handle), (C.MatPtr)(edges.handle), (C.double)(threshold1), (C.double)(threshold2), (C.int)(apertureSize), (C.int)(L2gradient))
}
