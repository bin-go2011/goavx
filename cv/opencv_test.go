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
	err := CvImread("../data/lena.jpg", IMREAD_GRAYSCALE, img)
	if err != nil {
		panic(err)
	}
	w := NewWindow("pic", WINDOW_AUTOSIZE)
	w.Show(img)
	w.WaitKey(0)
	w.Destory()
}

func TestPlayVideo(t *testing.T) {
	w := NewWindow("video", WINDOW_AUTOSIZE)
	cap, err := OpenVideoCapture(0)
	if err != nil {
		panic(err)
	}
	defer cap.Release()

	img := NewMat()
	for {
		err := cap.Read(img)
		if err != nil {
			panic(err)
		}
		w.Show(img)
		w.WaitKey(33)
	}
	w.Destory()

}
