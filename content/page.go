package content

type PageState int

const (
	PageStateInit      PageState = 0
	PageStateCreated   PageState = 1
	PageStateResumed   PageState = 2
	PageStatePaused    PageState = 3
	PageStateDestroyed PageState = 4
)

type Pager interface {
	init(panel *ZPanel, ctx *ZContext)
	setBundle(b map[string]interface{})

	getBundle() map[string]interface{}
	GetSate() PageState
	GetContentView() Viewer
	SetContentView(v Viewer)
	Create(bundle map[string]interface{})
	Resume()
	Pause()
	Destroy()
	Finish()
	StartPage(p Pager)
}

type Page struct {
	Context  *ZContext
	state    PageState
	bundle   map[string]interface{}
	rootView Viewer
	panel    *ZPanel
}

func (this *Page) init(panel *ZPanel, ctx *ZContext) {
	this.panel = panel
	this.Context = ctx
	this.state = PageStateInit
}
func (this *Page) setBundle(b map[string]interface{}) {
	this.bundle = b
}
func (this *Page) getBundle() map[string]interface{} {
	return this.bundle
}
func (this *Page) GetSate() PageState {
	return this.state
}
func (this *Page) SetContentView(v Viewer) {
	this.rootView = v
}
func (this *Page) GetContentView() Viewer {
	return this.rootView
}
func (this *Page) Create(bundle map[string]interface{}) {
	this.state = PageStateCreated

}

func (this *Page) Resume() {
	this.state = PageStateResumed
}

func (this *Page) Pause() {
	this.state = PageStatePaused
}

func (this *Page) Destroy() {
	this.state = PageStateDestroyed
	this.Context.Cancel()
}
func (this *Page) Finish() {
	this.panel.FinishPage(this)
}
func (this *Page) StartPage(p Pager) {
	this.panel.StartPage(p)
}
