package av

func NewAVFrame() (*AVFrame, error) {
	frame, err := AvAllocFrame()
	if err != nil {
		return nil, err
	}
	return frame, nil
}

func (frame *AVFrame) Release() {
	AvFreeFrame(frame)
}
