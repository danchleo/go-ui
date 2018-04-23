package content

type Rect struct {
	X      int
	Y      int
	Width  int
	Height int
}

func (this *Rect) Set(x, y, w, h int) {
	this.X = x
	this.Y = y
	this.Width = w
	this.Height = h
}
func (this *Rect) Empty() {
	this.X = 0
	this.Y = 0
	this.Width = 0
	this.Height = 0
}

type Viewer interface {
	Init()
	Measure(w, h int)
	SetMeasureSize(w, h int)
	GetMeasureSize() (w, h int)
	GetBounds() *Rect
	GetChildren() []Viewer
	GetChildrenSize() int
	AddChild(v Viewer)
	RemoveChild(v Viewer)
	Layout(x, y, w, h int)
	Draw(canvas ZCanvas)
	Recycle()

	setParent(v Viewer)
	getParent() Viewer
}
type View struct {
	measureWidth  int
	measureHeight int
	bounds        *Rect
	children      []Viewer
	parent        Viewer
}

func (this *View) Init() {
	this.bounds = &Rect{}
	this.children = make([]Viewer, 0)
}
func (this *View) Measure(w, h int) {

	this.SetMeasureSize(w, h);
}

func (this *View) SetMeasureSize(w, h int) {
	this.measureWidth = w
	this.measureHeight = h
}

func (this *View) GetMeasureSize() (w, h int) {
	return this.measureWidth, this.measureHeight
}
func (this *View) AddChild(v Viewer) {
	if v.getParent() == nil {
		this.children = append(this.children, v)
		v.setParent(this)
	}
}
func (this *View) RemoveChild(v Viewer) {
	for i := 0; i < len(this.children); i++ {
		if this.children[i] == v {

			if i < len(this.children)-1 {
				this.children = append(this.children[:i], this.children[i+1:]...)
			} else {
				this.children = this.children[:i]
			}
			v.Recycle()
			v.setParent(nil)
			return
		}
	}
}

func (this *View) GetBounds() *Rect {
	return this.bounds;
}
func (this *View) GetChildren() []Viewer {
	return this.children
}
func (this *View) GetChildrenSize() int {
	return len(this.children)
}
func (this *View) Layout(x, y, w, h int) {
	this.bounds.Set(x, y, w, h)
}

func (this *View) Draw(canvas ZCanvas) {
	children := this.children
	for i := 0; i < len(children); i++ {
		c := children[i]
		canvas.Save()
		b := c.GetBounds()
		canvas.SetTranslate(float64(b.X), float64(b.Y))
		c.Draw(canvas)
		canvas.Restore()
	}
}

func (this *View) Recycle() {

}

func (this *View) setParent(v Viewer) {
	this.parent = v
}
func (this *View) getParent() Viewer {
	return this.parent
}
