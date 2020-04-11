package cv

func OpenVideoCapture(device int) (*VideoCapture, error) {
	cap, err := CvNewVideoCapture()
	if err != nil {
		return nil, err
	}

	err = CvVideoCaptureOpenDevice(cap, device)
	if err != nil {
		return nil, err
	}

	return cap, nil
}

func (cap *VideoCapture) Release() {
	CvReleaseVideoCapture(cap)
}

func (cap *VideoCapture) Read(mat Mat) error {
	return CvVideoCaptureRead(cap, mat)
}
