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
	COLOR_BGR2GRAY = 6 //!< convert between RGB/BGR and grayscale, @ref color_convert_rgb_gray "color conversions"
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
