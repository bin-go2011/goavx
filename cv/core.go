package cv

func NewMat() (*Mat, error) {
	return CvNewMat()
}

func (m *Mat) Release() {

}
