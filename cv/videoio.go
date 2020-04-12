package cv

func OpenVideoDevice(device int) (*VideoCapture, error) {
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

func OpenVideoFile(file string) (*VideoCapture, error) {
	cap, err := CvNewVideoCapture()
	if err != nil {
		return nil, err
	}

	err = CvVideoCaptureOpenFile(cap, file)
	if err != nil {
		return nil, err
	}

	return cap, nil
}

func (cap *VideoCapture) Release() {
	CvReleaseVideoCapture(cap)
}

func (cap *VideoCapture) Read(mat *Mat) error {
	return CvVideoCaptureRead(cap, mat)
}
