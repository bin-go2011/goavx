package cv

type Mat struct {
	handle uintptr
}

func NewMat() Mat {
	return CvNewMat()
}
