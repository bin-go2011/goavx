package cv

func OpenVideoDevice(device int) (*VideoCapture, error) {
	cap, err := cvNewVideoCapture()
	if err != nil {
		return nil, err
	}

	err = cvVideoCaptureOpenDevice(cap, device)
	if err != nil {
		return nil, err
	}

	return cap, nil
}

func OpenVideoFile(file string) (*VideoCapture, error) {
	cap, err := cvNewVideoCapture()
	if err != nil {
		return nil, err
	}

	err = cvVideoCaptureOpenFile(cap, file)
	if err != nil {
		return nil, err
	}

	return cap, nil
}

func (cap *VideoCapture) Release() {
	cvReleaseVideoCapture(cap)
}

func (cap *VideoCapture) Read(mat *Mat) error {
	return cvVideoCaptureRead(cap, mat)
}

func (cap *VideoCapture) IsOpened() bool {
	return cvVideoCaptureIsOpened(cap)
}

func (cap *VideoCapture) Get(propId int) float64 {
	return cvVideoCaptureGet(cap, propId)
}
