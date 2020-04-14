package cv

import (
	"math"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

var (
	cvGaussianBlurProc,
	cvPyrDownProc,
	cvCannyProc,
	cvCvtColorProc *windows.Proc
)

func cvGaussianBlur(src *Mat, dst *Mat, ksizeX int, ksizeY int, sigmaX float64, sigmaY float64, borderType int) {
	if cvGaussianBlurProc == nil {
		cvGaussianBlurProc = goavx.LoadedDLL.MustFindProc("_cv_gaussian_blur")
	}

	cvGaussianBlurProc.Call(uintptr(src.handle), uintptr(dst.handle), uintptr(ksizeX), uintptr(ksizeY), uintptr(math.Float64bits(sigmaX)), uintptr(math.Float64bits(sigmaY)), uintptr(borderType))
}

func cvPyrDown(src *Mat, dst *Mat) {
	if cvPyrDownProc == nil {
		cvPyrDownProc = goavx.LoadedDLL.MustFindProc("_cv_pyrdown")
	}
	cvPyrDownProc.Call(uintptr(src.handle), uintptr(dst.handle))
}

func cvCanny(img *Mat, edges *Mat, threshold1 float64, threshold2 float64, apertureSize int, L2gradient bool) {
	if cvCannyProc == nil {
		cvCannyProc = goavx.LoadedDLL.MustFindProc("_cv_canny")
	}

	var grad uint8
	if L2gradient {
		grad = 1
	} else {
		grad = 0
	}
	cvCannyProc.Call(uintptr(img.handle), uintptr(edges.handle), uintptr(math.Float64bits(threshold1)), uintptr(math.Float64bits(threshold2)), uintptr(apertureSize), uintptr(grad))
}

func cvCvtColor(src *Mat, dst *Mat, code int) {
	if cvCvtColorProc == nil {
		cvCvtColorProc = goavx.LoadedDLL.MustFindProc("_cv_cvt_color")
	}
	cvCvtColorProc.Call(uintptr(src.handle), uintptr(dst.handle), uintptr(code))
}
