package cv

import (
	"math"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

var (
	cvGaussianBlurProc,
	cvPyrDownProc,
	cvCannyProc,
	cvCvtColorProc,
	cvMedianBlurProc,
	cvLaplacianProc,
	cvThresholdProc,
	cvResizeProc *windows.Proc
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

func cvMedianBlur(src *Mat, dst *Mat, ksize int) {
	if cvMedianBlurProc == nil {
		cvMedianBlurProc = goavx.LoadedDLL.MustFindProc("_cv_median_blur")
	}
	cvMedianBlurProc.Call(uintptr(src.handle), uintptr(dst.handle), uintptr(ksize))
}

func cvLaplacian(src *Mat, dst *Mat, ddepth int, ksize int) {
	if cvLaplacianProc == nil {
		cvLaplacianProc = goavx.LoadedDLL.MustFindProc("_cv_laplacian")
	}

	cvLaplacianProc.Call(uintptr(src.handle), uintptr(dst.handle), uintptr(ddepth), uintptr(ksize),
		uintptr(math.Float64bits(1.0)), uintptr(math.Float64bits(1.0)), uintptr(BORDER_DEFAULT))
}

func cvThreshold(src *Mat, dst *Mat, thresh float64, maxval float64, threshold_type int) {
	if cvThresholdProc == nil {
		cvThresholdProc = goavx.LoadedDLL.MustFindProc("_cv_threshold")
	}

	cvThresholdProc.Call(uintptr(src.handle), uintptr(dst.handle),
		uintptr(math.Float64bits(thresh)), uintptr(math.Float64bits(maxval)), uintptr(threshold_type))
}

func cvResize(src *Mat, dst *Mat, size Size, fx float64, fy float64, interpolation int) {
	if cvResizeProc == nil {
		cvResizeProc = goavx.LoadedDLL.MustFindProc("_cv_resize")
	}

	cvResizeProc.Call(uintptr(src.handle), uintptr(dst.handle),
		uintptr(unsafe.Pointer(&size)), uintptr(math.Float64bits(fx)), uintptr(math.Float64bits(fy)), INTER_LINEAR)
}
