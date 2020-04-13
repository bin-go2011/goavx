package av

func NewAVPacket() (*AVPacket, error) {
	pkt, err := avAllocPacket()
	if err != nil {
		return nil, err
	}
	return pkt, nil
}

func (pkt *AVPacket) Release() {
	avFreePacket(pkt)
}

func (pkt *AVPacket) StreamIndex() int32 {
	return int32(pkt.stream_index)
}

func (avctx *AVCodecContext) DecodeAudio(frame *AVFrame, got_frame_ptr *int, avpkt *AVPacket) error {
	return avcodecDecodeAudio4(avctx, frame, got_frame_ptr, avpkt)
}

func (avctx *AVCodecContext) Close() {
	avcodecFreeContext(avctx)
}
