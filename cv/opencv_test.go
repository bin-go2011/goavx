package cv

import (
	"fmt"
	"testing"
)

func TestVersion(t *testing.T) {
	fmt.Println(CvVersion())
}

func TestDisplayPicture(t *testing.T) {
	mat := NewMat()
	err := CvImread("../data/lena.jpg", IMREAD_GRAYSCALE, mat)
	if err != nil {
		panic(err)
	}
	w := NewWindow("pic", WINDOW_AUTOSIZE)
	w.ShowImage(mat)
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

	mat := NewMat()
	for {
		err := cap.Read(mat)
		if err != nil {
			panic(err)
		}
		w.ShowImage(mat)
		w.WaitKey(33)
	}
	w.Destory()

}
