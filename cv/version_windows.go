package cv

import (
	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

var (
	cvVersionProc *windows.Proc
)

func CvVersion() {
	if cvVersionProc == nil {
		cvVersionProc = goavx.LoadedDLL.MustFindProc("cv_version")
	}
	cvVersionProc.Call()
}
