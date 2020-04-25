package cv

import "testing"

func TestCartoonifyImage(t *testing.T) {
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

	srcColor, _ := Imread(SAMPLE_FILE)
	defer srcColor.Release()

	width, height := srcColor.Size()
	smallSize := Size{
		w: width / 2,
		h: height / 2,
	}

	smallImg, _ := NewMatFromSize(smallSize, CV_8UC3)
	defer smallImg.Release()

	Resize(srcColor, smallImg, smallSize, 0, 0, INTER_LINEAR)

	tmp, _ := NewMatFromSize(smallSize, CV_8UC3)
	defer tmp.Release()

	repetitions := 7

	for i := 0; i < repetitions; i++ {
		ksize := 9
		sigmaColor := 9.0
		sigmaSpace := 7.0

		BilateralFilter(smallImg, tmp, ksize, sigmaColor, sigmaSpace, BORDER_DEFAULT)
		BilateralFilter(tmp, smallImg, ksize, sigmaColor, sigmaSpace, BORDER_DEFAULT)

	}

	bigImg, _ := NewMat()
	defer bigImg.Release()

	Resize(smallImg, bigImg, Size{
		w: width,
		h: height,
	}, 0, 0, INTER_LINEAR)

	dst, _ := NewMatWith(int(width), int(height), CV_8UC3)
	defer dst.Release()

	dst.SetTo(0)

	bigImg.CopyTo(dst, mask)

	out := Imshow("final picture", dst)
	defer out.Destory()

	WaitKey(0)
}
