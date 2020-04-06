package av

import (
	"fmt"
	"reflect"
	"syscall"
	"unsafe"

	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

type AVFormatContext struct {
	handle uintptr
}

var (
	avformatOpenInputProc,
	avformatCloseInputProc,
	avDumpFormatProc,
	avFindBestAudioStreamProc,
	avFindBestVideoStreamProc,
	avReadFrameProc,
	avformatVersionProc *windows.Proc
)

func AvformatOpenInput(file string) (*AVFormatContext, error) {
	if avformatOpenInputProc == nil {
		avformatOpenInputProc = goavx.LoadedDLL.MustFindProc("_av_format_open_input")
	}

	f, _ := syscall.BytePtrFromString(file)

	r1, _, err := avformatOpenInputProc.Call(uintptr(unsafe.Pointer(f)))

	if r1 == 0 {
		return nil, err
	}

	fmtctx := &AVFormatContext{
		handle: r1,
	}
	return fmtctx, nil
}

func AvformatCloseInput(fmtctx *AVFormatContext) {
	if avformatCloseInputProc == nil {
		avformatCloseInputProc = goavx.LoadedDLL.MustFindProc("_av_format_close_input")
	}

	avformatCloseInputProc.Call(fmtctx.handle)
}

func AvDumpFormat(fmtctx *AVFormatContext, file string) {
	if avDumpFormatProc == nil {
		avDumpFormatProc = goavx.LoadedDLL.MustFindProc("_av_dump_format")
	}

	f, _ := syscall.BytePtrFromString(file)

	avDumpFormatProc.Call(fmtctx.handle, uintptr(unsafe.Pointer(f)))
}

func AvFindBestStream(fmtctx *AVFormatContext, mediaType int) (int32, error) {
	if avFindBestAudioStreamProc == nil {
		avFindBestAudioStreamProc = goavx.LoadedDLL.MustFindProc("_av_find_best_stream")
	}

	r1, _, err := avFindBestAudioStreamProc.Call(fmtctx.handle, uintptr(mediaType))
	return int32(r1), err
}

func AvReadFrame(fmtctx *AVFormatContext, pkt *AVPacket) int32 {
	if avReadFrameProc == nil {
		avReadFrameProc = goavx.LoadedDLL.MustFindProc("_av_read_frame")
	}

	r1, _, _ := avReadFrameProc.Call(fmtctx.handle, pkt.handle)

	pkt.size = *(*int32)(unsafe.Pointer(pkt.handle + unsafe.Offsetof(pkt.size)))
	pkt.stream_index = *(*int32)(unsafe.Pointer(pkt.handle + unsafe.Offsetof(pkt.stream_index)))
	pkt.data = pkt.handle + unsafe.Offsetof(pkt.data)
	pkt.flags = *(*int32)(unsafe.Pointer(pkt.handle + unsafe.Offsetof(pkt.flags)))
	pkt.duration = *(*int64)(unsafe.Pointer(pkt.handle + unsafe.Offsetof(pkt.duration)))
	pkt.pos = *(*int64)(unsafe.Pointer(pkt.handle + unsafe.Offsetof(pkt.pos)))

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&(pkt.Data)))
	sliceHeader.Cap = int(pkt.size)
	sliceHeader.Len = int(pkt.size)
	sliceHeader.Data = uintptr(*(*uintptr)(unsafe.Pointer(pkt.data)))

	return int32(r1)
}

func AvformatVersion() string {
	if avformatVersionProc == nil {
		avformatVersionProc = goavx.LoadedDLL.MustFindProc("_av_format_version")
	}

	r1, _, _ := avformatVersionProc.Call()

	version := int32(r1)
	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)
}
