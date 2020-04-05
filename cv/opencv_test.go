package cv

import (
	"testing"
)

func TestVersion(t *testing.T) {
	CvVersion()
}

// func TestNewWindow(t *testing.T) {
// 	w := NewWindow("Hello", 0)
// 	w.WaitKey(5)
// }

func TestOpenVideoDevice(t *testing.T) {
	VideoCaptureDevice(1)
}
