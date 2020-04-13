package av

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

type CAVPacket struct {
	buf             uintptr
	pts             int64
	dts             int64
	data            uintptr
	size            int32
	stream_index    int32
	flags           int32
	side_data       uintptr
	side_data_elems int32
	duration        int64
	pos             int64
}
type AVPacket struct {
	CAVPacket
	handle uintptr
	Data   []byte
}

var (
	avAllocPacketProc,
	avFreePacketProc *windows.Proc
)

func avAllocPacket() (*AVPacket, error) {
	if avAllocPacketProc == nil {
		avAllocPacketProc = goavx.LoadedDLL.MustFindProc("_av_packet_alloc")
	}

	r1, _, err := avAllocPacketProc.Call()

	if r1 == 0 {
		return nil, err
	}

	pkt := &AVPacket{
		handle: r1,
	}

	return pkt, nil
}

func avFreePacket(pkt *AVPacket) {
	if avFreePacketProc == nil {
		avFreePacketProc = goavx.LoadedDLL.MustFindProc("_av_free_packet")
	}

	avFreePacketProc.Call(pkt.handle)
}

type AVCodecContext struct {
	handle uintptr
}

var (
	avcodecOpenContextProc,
	avcodecFreeContextProc,
	avcodecDecodeAudio4Proc,
	avcodecVersionProc *windows.Proc
)

func avcodecOpenContext(fmtctx *AVFormatContext, id int32) (*AVCodecContext, error) {
	if avcodecOpenContextProc == nil {
		avcodecOpenContextProc = goavx.LoadedDLL.MustFindProc("_av_codec_open_context")
	}

	r1, _, err := avcodecOpenContextProc.Call(fmtctx.handle, uintptr(id))

	if r1 == 0 {
		return nil, err
	}

	avctx := &AVCodecContext{
		handle: r1,
	}
	return avctx, nil
}

func avcodecFreeContext(avctx *AVCodecContext) {
	if avcodecFreeContextProc == nil {
		avcodecFreeContextProc = goavx.LoadedDLL.MustFindProc("_av_codec_free_context")
	}

	avcodecFreeContextProc.Call(avctx.handle)
}

func avcodecDecodeAudio4(avctx *AVCodecContext, frame *AVFrame, got_frame_ptr *int, avpkt *AVPacket) error {
	if avcodecDecodeAudio4Proc == nil {
		avcodecDecodeAudio4Proc = goavx.LoadedDLL.MustFindProc("_av_codec_decode_audio4")
	}

	r1, _, err := avcodecDecodeAudio4Proc.Call(avctx.handle, frame.handle, uintptr(unsafe.Pointer(got_frame_ptr)), avpkt.handle)

	if int(r1) < 0 {
		return err
	}
	frame.nb_samples = *(*int32)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.nb_samples)))
	frame.format = *(*int32)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.format)))
	frame.linesize[0] = *(*int32)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.linesize)))
	frame.sample_rate = *(*int32)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.sample_rate)))
	frame.key_frame = *(*int32)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.key_frame)))
	frame.channel_layout = *(*int64)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.channel_layout)))
	frame.pkt_pos = *(*int64)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.pkt_pos)))
	frame.pkt_duration = *(*int64)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.pkt_duration)))
	frame.channels = *(*int32)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.channels)))
	frame.pkt_size = *(*int32)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.pkt_size)))

	size := frame.nb_samples * AvGetBytesPerSample(frame.format)

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&(frame.Data)))
	sliceHeader.Cap = int(size)
	sliceHeader.Len = int(size)
	sliceHeader.Data = uintptr(*(*uintptr)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(frame.handle + unsafe.Offsetof(frame.extended_data))))))

	return nil
}

func avcodecVersion() string {
	if avcodecVersionProc == nil {
		avcodecVersionProc = goavx.LoadedDLL.MustFindProc("_av_codec_version")
	}

	r1, _, _ := avcodecVersionProc.Call()

	version := int32(r1)
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}
