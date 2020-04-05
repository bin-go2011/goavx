package av

import (
	"testing"
)

const SAMPLE_FILE = "../big_buck_bunny.mp4"

var f *AVFormatContext

func init() {
	f, _ = AvformatOpenInput(SAMPLE_FILE)
}

func TestNewDemuxer(t *testing.T) {
	dmx, _ := NewDemuxer(SAMPLE_FILE)
	dmx.DumpFormat()
}

func TestAvFindBestAudioStream(t *testing.T) {
	i := AvFindBestStream(f, 1)
	if i != 0 {
		t.Errorf("wrong audio index %d, expected %d", i, 0)
	}
}

func TestAvFindBestVideoStream(t *testing.T) {
	i := AvFindBestStream(f, 0)
	if i != 1 {
		t.Errorf("wrong video index %d, expected %d", i, 1)
	}
}
