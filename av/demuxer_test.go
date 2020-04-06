package av

import (
	"testing"
)

func TestAudioDecoding(t *testing.T) {
	demuxer, _ := NewDemuxer("big_buck_bunny.mp4")
	defer demuxer.Release()

	adec, err := demuxer.OpenAudioDecoder()
	if err != nil {
		t.Error(err)
	}
	defer adec.Close()

	var got_frame int
	pkt, _ := NewAVPacket()
	defer pkt.Release()

	frame, _ := NewAVFrame()
	defer frame.Release()

	demuxer.ReadFrame(pkt)
	if pkt.stream_index == demuxer.AudioStreamIndex {
		err := adec.Decode(frame, &got_frame, pkt)
		if err != nil {
			t.Error(err)
		}
	}
}
