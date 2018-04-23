package main

import (
	."GO-UI/content"
	"image"
	"github.com/fogleman/gg"
	"image/color"
)

func main() {
	p := NewZPanel(500, 500, "test");
	page:=&TestPage{}
	p.StartPage(page)
	p.Run()
}
type TestButtonView struct {
	View
}

func (this *TestButtonView) Draw(canvas ZCanvas) {
	this.View.Draw(canvas)
	rect:=this.GetBounds()
	paint:=NewPaint()
	paint.Color = color.RGBA{0,225,225,255}
	canvas.DrawRect(0,0,float64(rect.Width),float64(rect.Height),10,paint)

	img:=image.NewRGBA(image.Rect(0,0,rect.Width,rect.Height))
	c:=gg.NewContextForRGBA(img)
	c.SetColor(color.White)
	c.DrawStringAnchored("button1",float64(rect.Width/2),float64(rect.Height/2),0.5,0.5)
	c.Fill()
	canvas.DrawImage(0,0,img)
}

type TestPage struct {
	Page
}

func (this *TestPage) Create(bundle map[string]interface{}) {
	this.Page.Create(bundle)
	v := &TestButtonView{}
	v.Init()
	v.Layout(20, 20, 130, 60)
	this.SetContentView(v)

}
func (this *TestPage) Resume() {
	this.Page.Resume()
}
