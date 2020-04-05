package av

import (
	"reflect"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

type AVCodecContext struct {
	handle uintptr
}

var (
	avcodecOpenContextProc,
	avcodecFreeContextProc,
	avcodecDecodeAudio4Proc *windows.Proc
)

func AvcodecOpenContext(fmtctx *AVFormatContext, id int32) (*AVCodecContext, error) {
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

func AvcodecFreeContext(avctx *AVCodecContext) {
	if avcodecFreeContextProc == nil {
		avcodecFreeContextProc = goavx.LoadedDLL.MustFindProc("_av_codec_free_context")
	}

	avcodecFreeContextProc.Call(avctx.handle)
}

func AvcodecDecodeAudio4(avctx *AVCodecContext, frame *AVFrame, got_frame_ptr *int, avpkt *AVPacket) error {
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
