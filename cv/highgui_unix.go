package cv

/*
#include "goavx/cv/highgui.h"
*/
import "C"

func CvNamedWindow(name string, flags int) {
	C.cv_named_window(C.CString(name), (C.int)(flags))
}

func CvWaitKey(delay int) {
	C.cv_wait_key((C.int)(delay))
}

func CvImshow(name string, img Mat) {
	C.cv_imshow(C.CString(name), (C.MatPtr)(img))
}

func CvDestroyWindow(name string) {
	C.cv_destroy_window(C.CString(name))
}
