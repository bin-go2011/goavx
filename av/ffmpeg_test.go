package av

import (
	"fmt"
	"path/filepath"
	"testing"
)

var SAMPLE_FILE string

func init() {
	SAMPLE_FILE, _ = filepath.Abs("../data/big_buck_bunny.mp4")

}

func TestNewDemuxer(t *testing.T) {
	dmx, _ := NewDemuxer(SAMPLE_FILE)
	dmx.DumpFormat()
}

func TestVersion(t *testing.T) {
	fmt.Println("libavformat: " + avformatVersion())
	fmt.Println("libavcodec: " + avcodecVersion())
}
