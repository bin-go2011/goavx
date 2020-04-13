package cv

import (
	"fmt"
	"path/filepath"
	"testing"
)

var (
	SAMPLE_VIDEO string
	SAMPLE_FILE  string
)

func init() {
	SAMPLE_VIDEO, _ = filepath.Abs("../data/big_buck_bunny.mp4")
	SAMPLE_FILE, _ = filepath.Abs("../data/lena.jpg")
}
func TestVersion(t *testing.T) {
	fmt.Println(CvVersion())
}

func TestDisplayPicture(t *testing.T) {
	mat, _ := NewMat()
	defer mat.Release()

	err := CvImread(SAMPLE_FILE, IMREAD_GRAYSCALE, mat)
	if err != nil {
		panic(err)
	}
	w := NewWindow("Example 2-2", WINDOW_AUTOSIZE)
	defer w.Destory()

	w.ShowImage(mat)
	w.WaitKey(0)
}

func TestOpenVideoDevice(t *testing.T) {
	w := NewWindow("Example 2-3", WINDOW_AUTOSIZE)
	defer w.Destory()

	cap, err := OpenVideoDevice(0)
	if err != nil {
		panic(err)
	}
	defer cap.Release()

	mat, _ := NewMat()
	defer mat.Release()

	for {
		err := cap.Read(mat)
		if err != nil {
			panic(err)
		}
		w.ShowImage(mat)
		if key := w.WaitKey(33); key >= 0 {
			break
		}
	}
}

func TestOpenVideoFile(t *testing.T) {
	w := NewWindow("Example 2-4", WINDOW_AUTOSIZE)
	defer w.Destory()

	cap, err := OpenVideoFile(SAMPLE_VIDEO)
	if err != nil {
		panic(err)
	}
	defer cap.Release()

	mat, _ := NewMat()
	defer mat.Release()

	for {
		err := cap.Read(mat)
		if err != nil {
			break
		}
		w.ShowImage(mat)
		w.WaitKey(33)
	}
}

func TestGaussianBlur(t *testing.T) {
	w_in := NewWindow("Example 2-5-in", WINDOW_AUTOSIZE)
	defer w_in.Destory()

	w_out := NewWindow("Example 2-5-out", WINDOW_AUTOSIZE)
	defer w_out.Destory()

	img, _ := NewMat()
	defer img.Release()

	err := CvImread(SAMPLE_FILE, IMREAD_UNCHANGED, img)
	if err != nil {
		panic(err)
	}
	w_in.ShowImage(img)

	out := img.GaussianBlur(5, 5, 3, 3)
	defer out.Release()

	out1 := out.GaussianBlur(5, 5, 3, 3)
	defer out1.Release()

	w_out.ShowImage(out1)

	w_in.WaitKey(0)
}

func TestPyrDown(t *testing.T) {
	w_in := NewWindow("Example 2-6-in", WINDOW_AUTOSIZE)
	defer w_in.Destory()

	w_out := NewWindow("Example 2-6-out", WINDOW_AUTOSIZE)
	defer w_out.Destory()

	img1, _ := NewMat()
	defer img1.Release()

	img2, _ := NewMat()
	defer img2.Release()

	err := CvImread(SAMPLE_FILE, IMREAD_UNCHANGED, img1)
	if err != nil {
		panic(err)
	}
	w_in.ShowImage(img1)

	CvPyrDown(img1, img2)
	w_out.ShowImage(img2)

	w_in.WaitKey(0)
}
