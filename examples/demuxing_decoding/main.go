package main

import (
	"fmt"
	"os"

	"github.com/bin-go2011/goavx/av"
)

func main() {
	demuxer, _ := av.NewDemuxer("../../big_buck_bunny.mp4")
	defer demuxer.Release()

	f, _ := os.Create("big_buck_bunny.pcm")
	defer f.Close()

	adec, err := demuxer.OpenAudioDecoder()
	if err != nil {
		panic(err)
	}
	defer adec.Close()

	var got_frame int
	pkt, _ := av.NewAVPacket()
	defer pkt.Release()

	frame, _ := av.NewAVFrame()
	defer frame.Release()

	for {
		n := demuxer.ReadFrame(pkt)
		if n < 0 {
			break
		}
		// fmt.Printf("%#v\n", pkt)
		if pkt.StreamIndex() == demuxer.AudioStreamIndex {
			err := adec.Decode(frame, &got_frame, pkt)
			if err != nil {
				panic(err)
			}
			// fmt.Printf("%#v\n", frame.CAVFrame)

			if got_frame > 0 {
				f.Write(frame.Data)
			}
		}
	}

	fmt.Println("Done")
}
