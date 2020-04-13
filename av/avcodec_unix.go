package av

/*
#include "goavx/av/codec.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type (
	AVPacket       C.struct_AVPacket
	AVCodecContext C.struct_AVCodecContext
)

func avcodecVersion() string {
	version := int32(C._av_codec_version())

	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}

func avAllocPacket() (*AVPacket, error) {
	pkt := C._av_packet_alloc()
	if pkt == nil {
		err := fmt.Errorf("packet allocation failed")
		return nil, err
	}
	return (*AVPacket)(pkt), nil
}

func avFreePacket(pkt *AVPacket) {
	C._av_free_packet((*C.struct_AVPacket)(pkt))
}

func avcodecOpenContext(fmtctx *AVFormatContext, id int32) (*AVCodecContext, error) {
	avctx := C._av_codec_open_context((*C.struct_AVFormatContext)(fmtctx), C.int(id))
	if avctx == nil {
		err := fmt.Errorf("open avcodec failed")
		return nil, err
	}
	return (*AVCodecContext)(avctx), nil
}

func avcodecFreeContext(avctx *AVCodecContext) {
	C._av_codec_free_context((*C.struct_AVCodecContext)(avctx))
}

func avcodecDecodeAudio4(avctx *AVCodecContext, frame *AVFrame, got_frame_ptr *int, avpkt *AVPacket) error {
	C._av_codec_decode_audio4((*C.struct_AVCodecContext)(avctx), (*C.struct_AVFrame)(frame), (*C.int)(unsafe.Pointer(got_frame_ptr)), (*C.struct_AVPacket)(avpkt))
	return nil
}
