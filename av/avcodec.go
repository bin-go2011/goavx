package av

func NewAVPacket() (*AVPacket, error) {
	pkt, err := AvAllocPacket()
	if err != nil {
		return nil, err
	}
	return pkt, nil
}

func (pkt *AVPacket) Release() {
	AvFreePacket(pkt)
}

func (pkt *AVPacket) StreamIndex() int32 {
	return int32(pkt.stream_index)
}

func (avctx *AVCodecContext) DecodeAudio(frame *AVFrame, got_frame_ptr *int, avpkt *AVPacket) error {
	return AvcodecDecodeAudio4(avctx, frame, got_frame_ptr, avpkt)
}

func (avctx *AVCodecContext) Close() {
	AvcodecFreeContext(avctx)
}
