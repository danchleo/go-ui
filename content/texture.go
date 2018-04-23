package content

import (
	"image"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Texture struct {
	Rgba    *image.RGBA
	Width   int
	Height  int
	IsAlpha bool
	texture uint32
	ctx    *ZContext
}

func NewTexture(ctx *ZContext, rgba *image.RGBA) (*Texture) {
	p := rgba.Bounds().Size()
	t := &Texture{Rgba: rgba, Width: p.X, Height: p.Y, ctx: ctx}

	t.Init()
	return t
}

func (t *Texture) Init() {
	if t.texture != 0 {
		return
	}
	var texture uint32
	gl.Enable(gl.TEXTURE_2D)
	gl.GenTextures(1, &texture)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(t.Rgba.Rect.Size().X),
		int32(t.Rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(t.Rgba.Pix))
	t.texture = texture
}
func (t *Texture) Recycle() {
	gl.DeleteTextures(1, &t.texture)
	t.texture = 0
}

var xx float32

func (t *Texture) Draw(c ZCanvas, x, y float64) {
	tran := c.GetTranslate()
	if tran != nil {
		x = x + tran.X
		y = y + tran.Y
	}
	winWidth, winHeight := (t.ctx.Value("window").(*glfw.Window)).GetSize()
	x, y = AppCoordinate2OpenGL(winWidth, winHeight, x, y)
	w, h := AppWidthHeight2OpenGL(winWidth, winHeight, (float64(t.Width)), (float64(t.Height)))
	if t.IsAlpha {
		gl.Enable(gl.BLEND);
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);
	}

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	gl.BindTexture(gl.TEXTURE_2D, t.texture)
	gl.LineWidth(0)
	gl.PointSize(0)
	gl.Begin(gl.QUADS)

	gl.Normal3f(float32(x), float32(y-h), 0) //
	gl.TexCoord2f(1, 1)
	gl.Vertex3f(float32(x+w), float32(y-h), 0) //

	gl.TexCoord2f(1, 0)
	gl.Vertex3f(float32(x+w), float32(y), 0) //
	gl.TexCoord2f(0, 0)
	gl.Vertex3f(float32(x), float32(y), 0) //
	gl.TexCoord2f(0, 1)
	gl.Vertex3f(float32(x), float32(y-h), 0) //

	gl.End()

	gl.PopMatrix()

}
func AppCoordinate2OpenGL(w, h int, x, y float64) (float64, float64) {

	return x*6/float64(w) - 3, -y*6/float64(h) + 3
}
func AppWidthHeight2OpenGL(w, h int, x, y float64) (float64, float64) {
	return x * 6 / float64(w), y * 6 / float64(h)
}
