package cv

const (
	BORDER_CONSTANT    = 0 //!< `iiiiii|abcdefgh|iiiiiii`  with some specified `i`
	BORDER_REPLICATE   = 1 //!< `aaaaaa|abcdefgh|hhhhhhh`
	BORDER_REFLECT     = 2 //!< `fedcba|abcdefgh|hgfedcb`
	BORDER_WRAP        = 3 //!< `cdefgh|abcdefgh|abcdefg`
	BORDER_REFLECT_101 = 4 //!< `gfedcb|abcdefgh|gfedcba`
	BORDER_TRANSPARENT = 5 //!< `uvwxyz|abcdefgh|ijklmno`

	BORDER_REFLECT101 = BORDER_REFLECT_101 //!< same as BORDER_REFLECT_101
	BORDER_DEFAULT    = BORDER_REFLECT_101 //!< same as BORDER_REFLECT_101
	BORDER_ISOLATED   = 16                 //!< do not look outside of ROI

)

const (
	THRESH_BINARY     = 0 //!< \f[\texttt{dst} (x,y) =  \fork{\texttt{maxval}}{if \(\texttt{src}(x,y) > \texttt{thresh}\)}{0}{otherwise}\f]
	THRESH_BINARY_INV = 1 //!< \f[\texttt{dst} (x,y) =  \fork{0}{if \(\texttt{src}(x,y) > \texttt{thresh}\)}{\texttt{maxval}}{otherwise}\f]
	THRESH_TRUNC      = 2 //!< \f[\texttt{dst} (x,y) =  \fork{\texttt{threshold}}{if \(\texttt{src}(x,y) > \texttt{thresh}\)}{\texttt{src}(x,y)}{otherwise}\f]
	THRESH_TOZERO     = 3 //!< \f[\texttt{dst} (x,y) =  \fork{\texttt{src}(x,y)}{if \(\texttt{src}(x,y) > \texttt{thresh}\)}{0}{otherwise}\f]
	THRESH_TOZERO_INV = 4 //!< \f[\texttt{dst} (x,y) =  \fork{0}{if \(\texttt{src}(x,y) > \texttt{thresh}\)}{\texttt{src}(x,y)}{otherwise}\f]
	THRESH_MASK       = 7
	THRESH_OTSU       = 8  //!< flag, use Otsu algorithm to choose the optimal threshold value
	THRESH_TRIANGLE   = 16 //!< flag, use Triangle algorithm to choose the optimal threshold value
)

const (
	COLOR_BGR2GRAY = 6 //!< convert between RGB/BGR and grayscale, @ref color_convert_rgb_gray "color conversions"
)

const (
	INTER_NEAREST = 0
	/** bilinear interpolation */
	INTER_LINEAR = 1
	/** bicubic interpolation */
	INTER_CUBIC = 2
	/** resampling using pixel area relation. It may be a preferred method for image decimation, as
	  it gives moire'-free results. But when the image is zoomed, it is similar to the INTER_NEAREST
	  method. */
	INTER_AREA = 3
	/** Lanczos interpolation over 8x8 neighborhood */
	INTER_LANCZOS4 = 4
	/** Bit exact bilinear interpolation */
	INTER_LINEAR_EXACT = 5
	/** mask for interpolation codes */
	INTER_MAX = 7
	/** flag, fills all of the destination image pixels. If some of them correspond to outliers in the
	  source image, they are set to zero */
	WARP_FILL_OUTLIERS = 8
)

func GaussianBlur(src *Mat, ksizeX int, ksizeY int, sigmaX float64, sigmaY float64) *Mat {
	dst, err := NewMat()
	if err != nil {
		return nil
	}
	cvGaussianBlur(src, dst, ksizeX, ksizeY, sigmaX, sigmaY, BORDER_DEFAULT)
	return dst
}

func PyrDown(src *Mat) *Mat {
	dst, err := NewMat()
	if err != nil {
		return nil
	}
	cvPyrDown(src, dst)
	return dst
}

func Canny(img *Mat, threshold1 float64, threshold2 float64, apertureSize int, L2gradient bool) *Mat {
	edges, err := NewMat()
	if err != nil {
		return nil
	}
	cvCanny(img, edges, threshold1, threshold2, apertureSize, L2gradient)
	return edges
}

func CvtColor(src *Mat, code int) *Mat {
	dst, err := NewMat()
	if err != nil {
		return nil
	}
	cvCvtColor(src, dst, code)
	return dst
}

func MedianBlur(src *Mat, ksize int) *Mat {
	dst, err := NewMat()
	if err != nil {
		return nil
	}
	cvMedianBlur(src, dst, ksize)
	return dst
}

func Laplacian(src *Mat, ddepth int, ksize int) *Mat {
	dst, err := NewMat()
	if err != nil {
		return nil
	}
	cvLaplacian(src, dst, ddepth, ksize)
	return dst
}

func Threshold(src *Mat, thresh float64, maxval float64, thresh_type int) *Mat {
	dst, err := NewMat()
	if err != nil {
		return nil
	}
	cvThreshold(src, dst, thresh, maxval, thresh_type)
	return dst
}

func Resize(src *Mat, dst *Mat, size Size, fx float64, fy float64, interpolation int) {
	cvResize(src, dst, size, fx, fy, interpolation)
}

func BilateralFilter(src *Mat, dst *Mat, d int, sigmaColor float64, sigmaSpace float64, borderType int) {
	cvBilateralFilter(src, dst, d, sigmaColor, sigmaSpace, borderType)
}
