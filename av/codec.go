package av

type AudioDecoder struct {
	avctx *AVCodecContext
}

func (adec *AudioDecoder) Decode(frame *AVFrame, got_frame_ptr *int, avpkt *AVPacket) error {
	return adec.avctx.DecodeAudio(frame, got_frame_ptr, avpkt)
}

func (adec *AudioDecoder) Close() {
	adec.avctx.Close()
}
