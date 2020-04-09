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
	CvNamedWindow(name, flags)
	return &Window{
		name: name,
	}
}

func (w *Window) WaitKey(delay int) {
	CvWaitKey(delay)
}

func (w *Window) Show(img Mat) {
	CvImshow(w.name, img)
}

func (w *Window) Destory() {
	CvDestroyWindow(w.name)
}
