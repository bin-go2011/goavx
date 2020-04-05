package av

import (
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

func AvAllocPacket() (*AVPacket, error) {
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

func AvFreePacket(pkt *AVPacket) {
	if avFreePacketProc == nil {
		avFreePacketProc = goavx.LoadedDLL.MustFindProc("_av_free_packet")
	}

	avFreePacketProc.Call(pkt.handle)
}
