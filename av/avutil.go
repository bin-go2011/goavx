package av

const (
	AVMEDIA_TYPE_UNKNOWN = -1 ///< Usually treated as AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_VIDEO   = 0
	AVMEDIA_TYPE_AUDIO   = 1
)

func NewAVFrame() (*AVFrame, error) {
	frame, err := avAllocFrame()
	if err != nil {
		return nil, err
	}
	return frame, nil
}

func (frame *AVFrame) Release() {
	avFreeFrame(frame)
}
