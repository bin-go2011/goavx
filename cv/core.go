package cv

const (
	// MatChannels1 is a single channel Mat.
	MatChannels1 = 0

	// MatChannels2 is 2 channel Mat.
	MatChannels2 = 8

	// MatChannels3 is 3 channel Mat.
	MatChannels3 = 16

	// MatChannels4 is 4 channel Mat.
	MatChannels4 = 24
)

const (
	CV_8U  = 0
	CV_8S  = 1
	CV_16U = 2
	CV_16S = 3
	CV_32S = 4
	CV_32F = 5
	CV_64F = 6
	CV_16F = 7

	CV_8UC3 = CV_8U + MatChannels3
)

type CvSize struct {
	w int32
	h int32
}

const (
	offsetOfMatRows = 8
	offsetOfMatCols = 12
)

func Version() string {
	return cvVersion()
}

func NewMat() (*Mat, error) {
	return cvNewMat()
}

func NewMatWith(width int, height int, mt int) (*Mat, error) {
	return cvNewMatWith(width, height, mt)
}

func NewMatFromSize(size CvSize, mt int) (*Mat, error) {
	return cvNewMatFromSize(size, mt)
}

func (m *Mat) Release() {
	cvReleaseMat(m)
}

func (m *Mat) Shape() (int32, int32, int32) {
	return cvMatShape(m)
}

func (m *Mat) Size() (width int32, height int32) {
	return cvMatSize(m)
}
