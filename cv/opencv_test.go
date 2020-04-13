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
	fmt.Println(cvVersion())
}

func TestDisplayPicture(t *testing.T) {
	mat, _ := NewMat()
	defer mat.Release()

	err := cvImread(SAMPLE_FILE, IMREAD_GRAYSCALE, mat)
	if err != nil {
		panic(err)
	}
	w := NewWindow("Example 2-2", WINDOW_AUTOSIZE)
	defer w.Destory()

	w.ShowImage(mat)
	WaitKey(0)
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
		if key := WaitKey(33); key >= 0 {
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
		WaitKey(33)
	}
}

func TestGaussianBlur(t *testing.T) {
	w_in := NewWindow("Example 2-5-in", WINDOW_AUTOSIZE)
	defer w_in.Destory()

	w_out := NewWindow("Example 2-5-out", WINDOW_AUTOSIZE)
	defer w_out.Destory()

	img, _ := NewMat()
	defer img.Release()

	err := cvImread(SAMPLE_FILE, IMREAD_UNCHANGED, img)
	if err != nil {
		panic(err)
	}
	w_in.ShowImage(img)

	out, _ := NewMat()
	defer out.Release()

	GaussianBlur(img, out, 5, 5, 3, 3)

	out1, _ := NewMat()
	defer out1.Release()

	GaussianBlur(out, out1, 5, 5, 3, 3)

	w_out.ShowImage(out1)

	WaitKey(0)
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

	err := cvImread(SAMPLE_FILE, IMREAD_UNCHANGED, img1)
	if err != nil {
		panic(err)
	}
	w_in.ShowImage(img1)

	cvPyrDown(img1, img2)
	w_out.ShowImage(img2)

	WaitKey(0)
}

func TestCanny(t *testing.T) {
	w_in := NewWindow("Example Gray", WINDOW_AUTOSIZE)
	defer w_in.Destory()

	w_out := NewWindow("Example Canny", WINDOW_AUTOSIZE)
	defer w_out.Destory()

	img_rgb, _ := NewMat()
	defer img_rgb.Release()

	img_gry, _ := NewMat()
	defer img_gry.Release()

	img_cny, _ := NewMat()
	defer img_cny.Release()

	err := cvImread(SAMPLE_FILE, IMREAD_UNCHANGED, img_rgb)
	if err != nil {
		panic(err)
	}
	cvCvtColor(img_rgb, img_gry, COLOR_BGR2GRAY)
	w_in.ShowImage(img_gry)

	cvCanny(img_gry, img_cny, 10, 100, 3, 1)
	w_out.ShowImage(img_cny)

	WaitKey(0)

}
