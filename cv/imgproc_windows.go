package cv

import (
	"math"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

var (
	cvGaussianBlurProc *windows.Proc
)

func CvGaussianBlur(src *Mat, dst *Mat, ksizeX int, ksizeY int, sigmaX float64, sigmaY float64, borderType int) {
	if cvGaussianBlurProc == nil {
		cvGaussianBlurProc = goavx.LoadedDLL.MustFindProc("_cv_gaussian_blur")
	}

	cvGaussianBlurProc.Call(uintptr(src.handle), uintptr(dst.handle), uintptr(ksizeX), uintptr(ksizeY), uintptr(math.Float64bits(sigmaX)), uintptr(math.Float64bits(sigmaY)), uintptr(borderType))
}
