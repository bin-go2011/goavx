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
	CvImread("../data/lena.jpg", 1, img)
	w := NewWindow("Hello", 1)
	w.Show(img)
	w.WaitKey(0)
}

func TestOpenVideoDevice(t *testing.T) {
	VideoCaptureDevice(1)
}
