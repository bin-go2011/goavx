package cv

type Window struct {
	name string
}

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
