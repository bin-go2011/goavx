package cv

func NewMat() (*Mat, error) {
	return cvNewMat()
}

func (m *Mat) Release() {
	cvReleaseMat(m)
}

func (m *Mat) Shape() (int32, int32, int32) {
	return cvMatShape(m)
}
