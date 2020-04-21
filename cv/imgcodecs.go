package cv

const (
	IMREAD_UNCHANGED = -1 //!< If set, return the loaded image as is (with alpha channel, otherwise it gets cropped).
	IMREAD_GRAYSCALE = 0  //!< If set, always convert image to the single channel grayscale image (codec internal conversion).
	IMREAD_COLOR     = 1  //!< If set, always convert image to the 3 channel BGR color image.
	IMREAD_ANYDEPTH  = 2  //!< If set, return 16-bit/32-bit image when the input has the corresponding depth, otherwise convert it to 8-bit.
	IMREAD_ANYCOLOR  = 4  //!< If set, the image is read in any possible color format.
)

func ReadImage(file string, flags int, mat *Mat) error {
	return cvImread(file, flags, mat)
}

func Imread(file string) (*Mat, error) {
	mat, err := NewMat()
	if err != nil {
		return nil, err
	}
	err = cvImread(file, IMREAD_UNCHANGED, mat)
	if err != nil {
		return nil, err
	}
	return mat, nil
}
