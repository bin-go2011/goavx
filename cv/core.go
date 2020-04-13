package cv

func NewMat() (*Mat, error) {
	return cvNewMat()
}

func (m *Mat) Release() {
	cvReleaseMat(m)
}
