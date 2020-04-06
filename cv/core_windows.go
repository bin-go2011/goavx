package cv

import (
	"fmt"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

var (
	cvVersionProc,
	cvNewMatProc *windows.Proc
)

func CvVersion() string {
	if cvVersionProc == nil {
		cvVersionProc = goavx.LoadedDLL.MustFindProc("cv_version")
	}
	r1, _, _ := cvVersionProc.Call()

	version := int32(r1)
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}

func CvNewMat() Mat {
	if cvNewMatProc == nil {
		cvNewMatProc = goavx.LoadedDLL.MustFindProc("cv_new_mat")
	}

	r1, _, _ := cvNewMatProc.Call()

	return Mat{
		handle: r1,
	}
}