package content

import (
	"image"
	"image/color"
)

type ZRotate struct {
	RotateAngle float64
	Dx          float64
	Dy          float64

	X float64
	Y float64
	Z float64
}
type ZScale struct {
	ScaleX float64
	ScaleY float64
}
type ZTranslate struct {
	X float64
	Y float64
}
type ZPaint struct {
	Color       color.Color
	Size        float64
	StrokeWidth float64
}

type ZCanvas interface {
	DrawCircle(dx, dy, radius float64, p ZPaint)
	DrawLine(x1, y1, x2, y2 float64, p ZPaint)
	DrawRect(x1, y1, x2, y2, rounCorner float64, p ZPaint)
	DrawImage(x, y float64, img image.Image)

	SetTranslate(x, y float64)
	SetScale(sx, sy float64)
	SetRotate(angle, dx, dy, x, y, z float64)

	Save()
	Restore()

	SetAlpha(a float64)

	GetTranslate() *ZTranslate
	GetScale() *ZScale
	GetRotate() *ZRotate
	GetAlpha() float64
}

func NewPaint() ZPaint {
	return ZPaint{
		Color:       color.Black,
		Size:        54.0,
		StrokeWidth: 5.0,
	}
}
