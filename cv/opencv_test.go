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

	err := ReadImage(SAMPLE_FILE, IMREAD_GRAYSCALE, mat)
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

	frame_width := cap.Get(CAP_PROP_FRAME_WIDTH)
	frame_height := cap.Get(CAP_PROP_FRAME_HEIGHT)
	fps := cap.Get(CAP_PROP_FPS)
	fourcc := int(cap.Get(CAP_PROP_FOURCC))

	fmt.Printf("CAP_PROP_FRAME_WIDTH: %f\n", frame_width)
	fmt.Printf("CAP_PROP_FRAME_HEIGHT: %f\n", frame_height)
	fmt.Printf("CAP_PROP_FPS: %f\n", fps)
	fmt.Printf("CAP_PROP_FOURCC: %c%c%c%c\n", fourcc&0xff, fourcc>>8&0xff, fourcc>>16&0xff, fourcc>>24&0xff)

	if !cap.IsOpened() { // check if we succeeded
		panic("Couldn't open capture.")
	}
	mat, _ := NewMat()
	defer mat.Release()

	for {
		err := cap.Read(mat)
		if err != nil {
			panic(err)
		}
		w.ShowImage(mat)

		if key := WaitKey(20); key > 0 {
			if key == int8('q') {
				break
			} else if key == int8('c') {
				Imwrite("save.png", mat)
				break
			}
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
		if WaitKey(33) > 0 {
			break
		}
	}
}

func TestGaussianBlur(t *testing.T) {
	w_in := NewWindow("Example 2-5-in", WINDOW_AUTOSIZE)
	defer w_in.Destory()

	w_out := NewWindow("Example 2-5-out", WINDOW_AUTOSIZE)
	defer w_out.Destory()

	img, _ := NewMat()
	defer img.Release()

	err := ReadImage(SAMPLE_FILE, IMREAD_UNCHANGED, img)
	if err != nil {
		panic(err)
	}
	w_in.ShowImage(img)

	out := GaussianBlur(img, 5, 5, 3, 3)
	defer out.Release()

	out1 := GaussianBlur(out, 5, 5, 3, 3)
	defer out1.Release()

	w_out.ShowImage(out1)

	WaitKey(0)
}

func TestPyrDown(t *testing.T) {
	in := NewWindow("Example 2-6-in", WINDOW_AUTOSIZE)
	defer in.Destory()

	out := NewWindow("Example 2-6-out", WINDOW_AUTOSIZE)
	defer out.Destory()

	img1, _ := NewMat()
	defer img1.Release()

	img2, _ := NewMat()
	defer img2.Release()

	err := ReadImage(SAMPLE_FILE, IMREAD_UNCHANGED, img1)
	if err != nil {
		panic(err)
	}
	in.ShowImage(img1)

	cvPyrDown(img1, img2)
	out.ShowImage(img2)

	WaitKey(0)
}

func TestCanny(t *testing.T) {
	in := NewWindow("Example Gray", WINDOW_AUTOSIZE)
	defer in.Destory()

	out := NewWindow("Example Canny", WINDOW_AUTOSIZE)
	defer out.Destory()

	rgb, err := Imread(SAMPLE_FILE)
	if err != nil {
		panic(err)
	}
	defer rgb.Release()

	gry := CvtColor(rgb, COLOR_BGR2GRAY)
	defer gry.Release()

	in.ShowImage(gry)

	cny := Canny(gry, 10, 100, 3, true)
	defer cny.Release()

	out.ShowImage(cny)

	WaitKey(0)

}

func TestSimplerAPIs(t *testing.T) {
	mat, _ := Imread(SAMPLE_FILE)
	defer mat.Release()

	w := Imshow("Load Image", mat)
	defer w.Destory()

	gray := CvtColor(mat, COLOR_BGR2GRAY)
	defer gray.Release()

	w1 := Imshow("Gray Image", gray)
	defer w1.Destory()

	WaitKey(0)
}

func TestLaplacian(t *testing.T) {
	mat, _ := Imread(SAMPLE_FILE)
	defer mat.Release()

	w := Imshow("Original Image", mat)
	defer w.Destory()

	gray := CvtColor(mat, COLOR_BGR2GRAY)
	defer gray.Release()

	const MEDIAN_BLUR_FILTER_SIZE = 7
	blurred := MedianBlur(gray, MEDIAN_BLUR_FILTER_SIZE)
	defer blurred.Release()

	const LAPLACIAN_FILTER_SIZE = 5
	edges := Laplacian(blurred, CV_8U, LAPLACIAN_FILTER_SIZE)
	defer edges.Release()

	const EDGES_THRESHOLD = 80
	mask := Threshold(edges, EDGES_THRESHOLD, 255, THRESH_BINARY_INV)
	defer mask.Release()

	w1 := Imshow("Filtered Image", mask)
	defer w1.Destory()

	WaitKey(0)
}

func TestResize(t *testing.T) {
	srcColor, _ := Imread(SAMPLE_FILE)
	defer srcColor.Release()

	width, height := srcColor.Size()
	smallSize := CvSize{
		w: width / 2,
		h: height / 2,
	}

	smallImg, _ := NewMatFromSize(smallSize, CV_8UC3)
	defer smallImg.Release()

	Resize(srcColor, smallImg, smallSize, 0, 0, INTER_LINEAR)
	out := Imshow("resided picture", smallImg)
	defer out.Destory()

	WaitKey(0)
}
