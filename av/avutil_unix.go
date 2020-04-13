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

func avAllocFrame() (*AVFrame, error) {
	frame := C._av_frame_alloc()
	if frame == nil {
		err := fmt.Errorf("AVFrame allocation failed")
		return nil, err
	}
	return (*AVFrame)(frame), nil
}

func avFreeFrame(frame *AVFrame) {
	C._av_frame_free((*C.struct_AVFrame)(frame))
}

func avGetBytesPerSample(samplefmt int32) int32 {
	return (int32)(C._av_get_bytes_per_sample((C.enum_AVSampleFormat)(samplefmt)))
}
