package av

/*
#include "goavx/av/util.h"
*/
import "C"
import (
	"fmt"
)

type (
	AVFrame C.struct_AVFrame
)

func AvAllocFrame() (*AVFrame, error) {
	frame := C._av_frame_alloc()
	if frame == nil {
		err := fmt.Errorf("AVFrame allocation failed")
		return nil, err
	}
	return (*AVFrame)(frame), nil
}

func AvFreeFrame(frame *AVFrame) {
	C._av_frame_free((*C.struct_AVFrame)(frame))
}

func AvGetBytesPerSample(samplefmt int32) int32 {
	return (int32)(C._av_get_bytes_per_sample((C.enum_AVSampleFormat)(samplefmt)))
}
