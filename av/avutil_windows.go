package av

import (
	"github.com/bin-go2011/goavx"
	"golang.org/x/sys/windows"
)

const AV_NUM_DATA_POINTERS = 8

type CAVFrame struct {
	data                   [AV_NUM_DATA_POINTERS]uintptr
	linesize               [AV_NUM_DATA_POINTERS]int32
	extended_data          uintptr
	width                  int32
	height                 int32
	nb_samples             int32
	format                 int32
	key_frame              int32
	pict_type              int32
	sample_aspect_ratio    int64
	pts                    int64
	pkt_pts                int64
	pkt_dts                int64
	coded_picture_number   int32
	display_picture_number int32
	quality                int32
	opaque                 uintptr
	errornumber            [AV_NUM_DATA_POINTERS]uint64
	repeat_pict            int32
	interlaced_frame       int32
	top_field_first        int32
	palette_has_changed    int32
	reordered_opaque       int64
	sample_rate            int32
	channel_layout         int64
	buf                    [AV_NUM_DATA_POINTERS]uintptr
	extended_buf           uintptr
	nb_extended_buf        int32
	side_data              uintptr
	nb_side_data           int32
	flags                  int32
	color_range            int32
	color_primaries        int32
	color_trc              int32
	colorspace             int32
	chroma_location        int32
	best_effort_timestamp  int64
	pkt_pos                int64
	pkt_duration           int64
	metadata               uintptr
	decode_error_flags     int32
	channels               int32
	pkt_size               int32
	hw_frames_ctx          uintptr
	opaque_ref             uintptr

	crop_top    uintptr
	crop_bottom uintptr
	crop_left   uintptr
	crop_right  uintptr

	private_ref uintptr
}

type AVFrame struct {
	CAVFrame
	handle uintptr
	Data   []byte
}

const (
	AVMEDIA_TYPE_UNKNOWN = -1 ///< Usually treated as AVMEDIA_TYPE_DATA
	AVMEDIA_TYPE_VIDEO   = 0
	AVMEDIA_TYPE_AUDIO   = 1
)

var (
	avAllocFrameProc,
	avFreeFrameProc,
	avGetBytesPerSampleProc *windows.Proc
)

func AvAllocFrame() (*AVFrame, error) {
	if avAllocFrameProc == nil {
		avAllocFrameProc = goavx.LoadedDLL.MustFindProc("_av_frame_alloc")
	}

	r1, _, err := avAllocFrameProc.Call()

	if r1 == 0 {
		return nil, err
	}

	frame := &AVFrame{
		handle: r1,
	}

	return frame, nil
}

func AvFreeFrame(frame *AVFrame) {
	if avFreeFrameProc == nil {
		avFreeFrameProc = goavx.LoadedDLL.MustFindProc("_av_frame_free")
	}
	avFreeFrameProc.Call(frame.handle)
}

func AvGetBytesPerSample(samplefmt int32) int32 {
	if avGetBytesPerSampleProc == nil {
		avGetBytesPerSampleProc = goavx.LoadedDLL.MustFindProc("_av_get_bytes_per_sample")
	}

	r1, _, _ := avGetBytesPerSampleProc.Call(uintptr(samplefmt))

	return int32(r1)
}
