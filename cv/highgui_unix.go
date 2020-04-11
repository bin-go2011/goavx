package cv

/*
#include "goavx/cv/highgui.h"
*/
import "C"

func CvNamedWindow(name string, flags int) {
	C._cv_named_window(C.CString(name), (C.int)(flags))
}

func CvWaitKey(delay int) {
	C._cv_wait_key((C.int)(delay))
}

func CvImshow(name string, img Mat) {
	C._cv_imshow(C.CString(name), (C.MatPtr)(img))
}

func CvDestroyWindow(name string) {
	C._cv_destroy_window(C.CString(name))
}
