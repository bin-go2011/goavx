package av

import "fmt"

/*
#include <stdio.h>
#include "goavx/av/format.h"
*/
import "C"

type (
	AVFormatContext C.struct_AVFormatContext
)

func avformatVersion() string {
	version := int32(C._av_format_version())

	subminor := version & 0xff
	minor := version >> 8 & 0xff
	major := version >> 16 & 0xff
	return fmt.Sprintf("%d.%d.%d", major, minor, subminor)

}

func avDumpFormat(fmtctx *AVFormatContext, file string) {
	C._av_dump_format((*C.struct_AVFormatContext)(fmtctx), C.CString(file))
}

func avformatOpenInput(file string) (*AVFormatContext, error) {
	fmtctx := (*AVFormatContext)(C._av_format_open_input(C.CString(file)))
	if fmtctx == nil {
		err := fmt.Errorf("%s", C.GoString(C.strerror(C.ENOENT)))
		return nil, err
	}
	return fmtctx, nil
}

func avformatCloseInput(fmtctx *AVFormatContext) {
	C._av_format_close_input((*C.struct_AVFormatContext)(fmtctx))
}

func avFindBestStream(fmtctx *AVFormatContext, mediaType int) (int32, error) {
	idx := C._av_find_best_stream((*C.struct_AVFormatContext)(fmtctx), C.enum_AVMediaType(mediaType))
	if idx < 0 {
		err := fmt.Errorf("No medium type %d found ", mediaType)
		return -1, err
	}
	return int32(idx), nil
}

func avReadFrame(fmtctx *AVFormatContext, pkt *AVPacket) int32 {
	return (int32)(C._av_read_frame((*C.struct_AVFormatContext)(fmtctx), (*C.struct_AVPacket)(pkt)))
}
