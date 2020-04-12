package cv

func NewMat() (*Mat, error) {
	return CvNewMat()
}

func (m *Mat) Release() {
	CvReleaseMat(m)
}

func (m *Mat) GaussianBlur(ksizeX int, ksizeY int, sigmaX float64, sigmaY float64) *Mat {
	dst, _ := NewMat()
	CvGaussianBlur(m, dst, ksizeX, ksizeY, sigmaX, sigmaY, BORDER_DEFAULT)
	return dst
}
