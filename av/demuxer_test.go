package av

import (
	"fmt"
	"testing"
)

func TestAudioDecoding(t *testing.T) {
	demuxer, _ := NewDemuxer(SAMPLE_FILE)
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

func TestVideoDecoding(t *testing.T) {
	demuxer, _ := NewDemuxer(SAMPLE_FILE)
	defer demuxer.Release()

	demuxer.FindBestVideoStream()

	pkt, _ := NewAVPacket()
	defer pkt.Release()

	for {
		demuxer.ReadFrame(pkt)
		if pkt.stream_index == demuxer.VideoStreamIndex {
			fmt.Printf("%#v\n", pkt.CAVPacket)

		}
	}
}
