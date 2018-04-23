package content

import (
	"image"
	"github.com/fogleman/gg"
	"math"
	"image/draw"
)

type GL2ZCanvas struct {
	Alpha    float64
	Restores []GL2ZCanvas

	Scale     *ZScale
	Rotate    *ZRotate
	Translate *ZTranslate
	ctx    *ZContext
}

func NewGL2Canvas(ctx    *ZContext) ZCanvas {
	return &GL2ZCanvas{
		Restores: make([]GL2ZCanvas, 0),
		ctx:ctx,
	}
}
func (c *GL2ZCanvas) clone() GL2ZCanvas {
	return GL2ZCanvas{
		Alpha:     c.Alpha,
		Scale:     c.Scale,
		Rotate:    c.Rotate,
		Translate: c.Translate,
	}
}

func (c *GL2ZCanvas) DrawCircle(dx, dy, radius float64, p ZPaint) {
	rgba := image.NewRGBA(image.Rect(0, 0, int(2*radius), int(2*radius)))
	dc := gg.NewContextForRGBA(rgba)
	dc.DrawCircle(radius, radius, radius)
	dc.SetColor(p.Color)
	dc.Fill()
	t := NewTexture(c.ctx,rgba)
	t.IsAlpha = true
	t.Draw(c, dx-radius, dy-radius);
}
func (c *GL2ZCanvas) DrawLine(x1, y1, x2, y2 float64, p ZPaint) {
	w := int(math.Abs(x2 - x1))
	if w == 0 {
		w = int(p.StrokeWidth)
	}
	h := int(math.Abs(y2 - y1))
	if h == 0 {
		h = int(p.StrokeWidth)
	}
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	dc := gg.NewContextForRGBA(rgba)
	dc.DrawLine(0, 0, math.Abs(x2-x1), math.Abs(y2-y1))
	dc.SetLineWidth(p.StrokeWidth)
	dc.SetColor(p.Color)
	dc.Stroke()
	t := NewTexture(c.ctx,rgba)
	t.IsAlpha = true
	t.Draw(c, math.Min(x1, x2), math.Min(y1, y2));
}

func (c *GL2ZCanvas) DrawRect(x1, y1, x2, y2, rounCorner float64, p ZPaint) {


	rgba := image.NewRGBA(image.Rect(0, 0, int(math.Abs(x2-x1)), int(math.Abs(y2-y1))))
	dc := gg.NewContextForRGBA(rgba)
	dc.DrawRoundedRectangle(0, 0, math.Abs(x2-x1), math.Abs(y2-y1), rounCorner)
	dc.SetColor(p.Color)
	dc.Fill()
	t := NewTexture(c.ctx,rgba)
	t.IsAlpha = true
	t.Draw(c, math.Min(x1, x2), math.Min(y1, y2));
}
func (c *GL2ZCanvas) DrawImage(x, y float64, img image.Image) {
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	t := NewTexture(c.ctx,rgba)
	t.IsAlpha = true
	t.Draw(c, x, y);
}

func (c *GL2ZCanvas) SetTranslate(x, y float64) {
	if c.Translate != nil {
		c.Translate.X = x
		c.Translate.Y = y

	} else {
		c.Translate = &ZTranslate{x, y}
	}

}
func (c *GL2ZCanvas) SetScale(sx, sy float64) {
	if c.Scale != nil {
		c.Scale.ScaleX = sx
		c.Scale.ScaleY = sy

	} else {
		c.Scale = &ZScale{sx, sy}
	}
}
func (c *GL2ZCanvas) SetRotate(angle, dx, dy, x, y, z float64) {
	if c.Rotate != nil {
		c.Rotate.RotateAngle = angle
		c.Rotate.Dx = dx
		c.Rotate.Dy = dy
		c.Rotate.X = x
		c.Rotate.Y = y
		c.Rotate.Z = z

	} else {
		c.Rotate = &ZRotate{angle, dx, dy, x, y, z}
	}

}

func (c *GL2ZCanvas) GetTranslate() *ZTranslate {
	return c.Translate
}
func (c *GL2ZCanvas) GetScale() *ZScale {
	return c.Scale

}
func (c *GL2ZCanvas) GetRotate() *ZRotate {
	return c.Rotate
}

func (c *GL2ZCanvas) Save() {
	c.Alpha = 0
	c.Scale = nil
	c.Translate = nil
	c.Rotate = nil

	c.Restores = append(c.Restores, c.clone())
}
func (c *GL2ZCanvas) Restore() {
	c1 := c.Restores[len(c.Restores)-1]
	c.Restores = c.Restores[:len(c.Restores)-1]
	c.Alpha = c1.Alpha;
	c.Translate = c1.Translate;
	c.Rotate = c1.Rotate;
	c.Scale = c1.Scale

}

func (c *GL2ZCanvas) SetAlpha(a float64) {
	c.Alpha = a;
}

func (c *GL2ZCanvas) GetAlpha() float64 {
	return c.Alpha
}
