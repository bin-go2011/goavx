package av

func (avctx *AVCodecContext) DecodeAudio(frame *AVFrame, got_frame_ptr *int, avpkt *AVPacket) error {
	return AvcodecDecodeAudio4(avctx, frame, got_frame_ptr, avpkt)
}

func (avctx *AVCodecContext) Close() {
	AvcodecFreeContext(avctx)
}
