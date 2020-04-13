package cv

type Window struct {
	name string
}

const (
	WINDOW_NORMAL   = 0x00000000 //!< the user can resize the window (no constraint) / also use to switch a fullscreen window to a normal size.
	WINDOW_AUTOSIZE = 0x00000001 //!< the user cannot resize the window, the size is constrainted by the image displayed.
	WINDOW_OPENGL   = 0x00001000 //!< window with opengl support.
)

func NewWindow(name string, flags int) *Window {
	cvNamedWindow(name, flags)
	return &Window{
		name: name,
	}
}

func WaitKey(delay int) int8 {
	return cvWaitKey(delay)
}

func (w *Window) ShowImage(mat *Mat) {
	cvImshow(w.name, mat)
}

func (w *Window) Destory() {
	cvDestroyWindow(w.name)
}
