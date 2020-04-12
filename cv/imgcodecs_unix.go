package cv

/*
#include "goavx/cv/imgcodecs.h"
*/
import "C"
import "fmt"

func CvImread(file string, flags int, mat *Mat) error {
	ret := C._cv_imread(C.CString(file), (C.int)(flags), (C.MatPtr)(mat.handle))
	if int32(ret) < 0 {
		err := fmt.Errorf("failed to read %s", file)
		return err
	}
	return nil
}
