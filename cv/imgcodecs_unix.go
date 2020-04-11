package cv

/*
#include "goavx/cv/imgcodecs.h"
*/
import "C"

const (
	IMREAD_UNCHANGED = -1 //!< If set, return the loaded image as is (with alpha channel, otherwise it gets cropped).
	IMREAD_GRAYSCALE = 0  //!< If set, always convert image to the single channel grayscale image (codec internal conversion).
	IMREAD_COLOR     = 1  //!< If set, always convert image to the 3 channel BGR color image.
	IMREAD_ANYDEPTH  = 2  //!< If set, return 16-bit/32-bit image when the input has the corresponding depth, otherwise convert it to 8-bit.
	IMREAD_ANYCOLOR  = 4  //!< If set, the image is read in any possible color format.
)

func CvImread(file string, flags int, mat Mat) {
	C._cv_imread(C.CString(file), (C.int)(flags), (C.MatPtr)(mat))
}
