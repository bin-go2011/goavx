package av

import (
	"fmt"
	"os"
	"testing"
)

func TestDemuxingAndDecoding(t *testing.T) {
	demuxer, _ := NewDemuxer("big_buck_bunny.mp4")
	defer demuxer.Release()

	f, _ := os.Create("big_buck_bunny.pcm")
	defer f.Close()

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

	for {
		n := demuxer.ReadFrame(pkt)
		if n < 0 {
			break
		}
		// fmt.Printf("%#v\n", pkt)
		if pkt.stream_index == demuxer.AudioStreamIndex {
			err := adec.Decode(frame, &got_frame, pkt)
			if err != nil {
				t.Error(err)
			}
			// fmt.Printf("%#v\n", frame.CAVFrame)

			if got_frame > 0 {
				f.Write(frame.Data)
			}
		}
	}

	fmt.Println("Done")
}
