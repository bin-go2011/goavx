package cv

type Window struct {
	name string
	open bool
}

func NewWindow(name string, flags int) *Window {
	CvNamedWindow(name, flags)
	return &Window{
		name: name,
		open: true,
	}
}

func (w *Window) WaitKey(delay int) {
	CvWaitKey(delay)
}

func (w *Window) Show(mat Mat) {
	CvImshow(w.name, mat)
}
