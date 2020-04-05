package cv

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	fmt.Println(CvVersion())
}

func TestDisplayPicture(t *testing.T) {
	img := NewMat()
	CvImread("../data/lena.jpg", IMREAD_GRAYSCALE, img)
	w := NewWindow("Hello", WINDOW_AUTOSIZE)
	w.Show(img)
	w.WaitKey(0)
}

func TestOpenVideoDevice(t *testing.T) {
	VideoCaptureDevice(1)
}
