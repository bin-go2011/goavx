package av

import (
	"fmt"
	"testing"
)

const SAMPLE_FILE = "../data/big_buck_bunny.mp4"

func TestNewDemuxer(t *testing.T) {
	dmx, _ := NewDemuxer(SAMPLE_FILE)
	dmx.DumpFormat()
}

func TestVersion(t *testing.T) {
	fmt.Println("libavformat: " + AvformatVersion())
	fmt.Println("libavcodec: " + AvcodecVersion())
}
