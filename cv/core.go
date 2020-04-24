package cv

const (
	CV_8U  = 0
	CV_8S  = 1
	CV_16U = 2
	CV_16S = 3
	CV_32S = 4
	CV_32F = 5
	CV_64F = 6
	CV_16F = 7
)
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

func (m *Mat) Release() {
	cvReleaseMat(m)
}

func (m *Mat) Shape() (int32, int32, int32) {
	return cvMatShape(m)
}

func (m *Mat) Size() (width int32, height int32) {
	return cvMatSize(m)
}
