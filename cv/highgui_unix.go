package cv

/*
#include "goavx/cv/highgui.h"
*/
import "C"

func cvNamedWindow(name string, flags int) {
	C._cv_named_window(C.CString(name), C.int(flags))
}

func cvWaitKey(delay int) int8 {
	return int8(C._cv_wait_key(C.int(delay)))
}

func cvImshow(name string, img *Mat) {
	C._cv_imshow(C.CString(name), C.MatPtr(img.handle))
}

func cvDestroyWindow(name string) {
	C._cv_destroy_window(C.CString(name))
}
