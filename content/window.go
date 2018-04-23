package content

import (
	"runtime"
	"log"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"context"
)

type ZContext struct {
	context.Context
	Cancel context.CancelFunc
}
type ZPanel struct {
	Width   int
	Height  int
	Title   string
	Window  *glfw.Window
	Canvas  ZCanvas
	Context *ZContext

	Pages []Pager
}

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}
func NewZPanel(w, h int, title string) *ZPanel {
	return &ZPanel{w, h, title, nil, nil, nil, make([]Pager, 0)}
}
func (w *ZPanel) GetCurrentWidthHeight() (int, int) {
	return w.Window.GetSize()
}
func (w *ZPanel) StartPage(p Pager) {
	if p.GetSate() != PageStateInit {
		return
	}
	p.init(w, w.Context)
	if ( len(w.Pages) > 0) {
		previousPage := w.Pages[len(w.Pages)-1]
		previousPage.Pause()
	}
	w.Pages = append(w.Pages, p)

}
func (w *ZPanel) FinishPage(p Pager) {
	for i := 0; i < len(w.Pages); i++ {
		if w.Pages[i] == p {

			if i < len(w.Pages)-1 {
				w.Pages = append(w.Pages[:i], w.Pages[i+1:]...)
			} else {
				w.Pages = w.Pages[:i]
			}
			p.Pause()
			p.Destroy()
			return
		}
	}
}
func (w *ZPanel) Run() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	var err error
	w.Window, err = glfw.CreateWindow(int(w.Width), int(w.Height), w.Title, nil, nil)
	if err != nil {
		panic(err)
	}
	w.Context = &ZContext{}
	w.Context.Context = context.WithValue(context.Background(), "window", w.Window)
	w.Canvas = NewGL2Canvas(w.Context)
	w.Window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		panic(err)
	}

	gl.ClearColor(0.5, 0.5, 0.5, 0.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)
	ambient := []float32{0.5, 0.5, 0.5, 1}
	diffuse := []float32{1, 1, 1, 1}
	lightPosition := []float32{-5, 5, 10, 0}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &lightPosition[0])
	gl.Enable(gl.LIGHT0)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	//gl.Frustum(-1, 1, -1, 1, 1.0, 5.0)
	gl.Ortho(-3, 3, -3, 3, -3.0, 100.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	for !w.Window.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.ClearColor(1, 1, 1, 0)

		if len(w.Pages) > 0 {

			p := w.Pages[len(w.Pages)-1]
			if p.GetSate() == PageStateInit {
				p.Create(p.getBundle())
			}
			rootView := p.GetContentView()
			if rootView != nil {
				b := rootView.GetBounds()
				c := w.Canvas
				c.Save()
				c.SetTranslate(float64(b.X), float64(b.Y))
				rootView.Draw(c)
				c.Restore()
			}
			if p.GetSate() == PageStateCreated {
				p.Resume()
			} else if p.GetSate() == PageStatePaused {
				p.Resume()
			}
		}

		w.Window.SwapBuffers()
		glfw.PollEvents()
	}
}
