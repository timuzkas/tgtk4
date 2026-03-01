package tgtk4

import (
	"fmt"
	"math"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func IconBtnContent(iconName, labelText string) *gtk.Box {
	content := gtk.NewBox(gtk.OrientationHorizontal, 6)
	content.SetHAlign(gtk.AlignCenter)
	img := gtk.NewImageFromIconName(iconName)
	content.Append(img)
	if labelText != "" {
		lbl := gtk.NewLabel(labelText)
		content.Append(lbl)
	}
	return content
}

func IconBtn(iconName, labelText string) *gtk.Button {
	btn := gtk.NewButton()
	btn.SetChild(IconBtnContent(iconName, labelText))
	btn.AddCSSClass("action-btn")
	return btn
}

func MiniActionBtn(icon, label string, onClick func()) *gtk.Button {
	btn := gtk.NewButton()
	box := gtk.NewBox(gtk.OrientationHorizontal, 4)
	box.Append(gtk.NewImageFromIconName(icon))
	if label != "" {
		box.Append(gtk.NewLabel(label))
	}
	btn.SetChild(box)
	btn.AddCSSClass("mini-action-btn")
	if onClick != nil {
		btn.ConnectClicked(onClick)
	}
	return btn
}

func NewHeader(title string) *gtk.HeaderBar {
	header := gtk.NewHeaderBar()
	titleLabel := gtk.NewLabel("// " + title)
	titleLabel.AddCSSClass("title")
	header.SetTitleWidget(titleLabel)
	return header
}

func NewProgressBar() *gtk.ProgressBar {
	pb := gtk.NewProgressBar()
	pb.AddCSSClass("progressview")
	return pb
}

func NewLabeledSlider(label string, min, max, step float64) (*gtk.Box, *gtk.Scale) {
	box := gtk.NewBox(gtk.OrientationVertical, 0)
	lbl := gtk.NewLabel(label)
	lbl.SetHAlign(gtk.AlignStart)
	lbl.AddCSSClass("header-label")
	box.Append(lbl)
	
	scale := gtk.NewScaleWithRange(gtk.OrientationHorizontal, min, max, step)
	scale.SetHExpand(true)
	scale.SetDrawValue(false)
	scale.SetHasTooltip(true)
	box.Append(scale)
	return box, scale
}

func NewPicture(path string, w, h int) *gtk.Picture {
	pic := gtk.NewPicture()
	pic.SetContentFit(gtk.ContentFitContain)
	if path != "" {
		go func() {
			pb, err := gdkpixbuf.NewPixbufFromFileAtScale(path, w, h, true)
			if err == nil {
				glib.IdleAdd(func() {
					pic.SetPaintable(gdk.NewTextureForPixbuf(pb))
				})
			}
		}()
	}
	return pic
}

type Lightbox struct {
	Overlay *gtk.Box
	Image   *gtk.Picture
}

func NewLightbox(parent *gtk.Overlay) *Lightbox {
	box := gtk.NewBox(gtk.OrientationVertical, 0)
	box.AddCSSClass("lightbox-overlay")
	box.SetHAlign(gtk.AlignFill)
	box.SetVAlign(gtk.AlignFill)
	box.SetVisible(false)

	img := gtk.NewPicture()
	img.SetContentFit(gtk.ContentFitContain)
	img.AddCSSClass("lightbox-image")
	img.SetHExpand(true); img.SetVExpand(true)
	img.SetHAlign(gtk.AlignCenter); img.SetVAlign(gtk.AlignCenter)
	box.Append(img)

	click := gtk.NewGestureClick()
	click.ConnectPressed(func(n int, x, y float64) {
		box.RemoveCSSClass("active")
		glib.TimeoutAdd(250, func() bool {
			box.SetVisible(false)
			return false
		})
	})
	box.AddController(click)

	parent.AddOverlay(box)
	return &Lightbox{Overlay: box, Image: img}
}

func (l *Lightbox) Show(path string) {
	go func() {
		pb, err := gdkpixbuf.NewPixbufFromFile(path)
		if err == nil {
			glib.IdleAdd(func() {
				l.Image.SetPaintable(gdk.NewTextureForPixbuf(pb))
				l.Overlay.SetVisible(true)
				glib.TimeoutAdd(10, func() bool {
					l.Overlay.AddCSSClass("active")
					return false
				})
			})
		}
	}()
}

type AnimatedPicture struct {
	*gtk.DrawingArea
	progress float64
	accent   [3]float64
	pixbuf   *gdkpixbuf.Pixbuf
}

func NewAnimatedPicture(path string, w, h int, accentHex string) *gtk.DrawingArea {
	da := gtk.NewDrawingArea()
	da.SetSizeRequest(w, h)
	ap := &AnimatedPicture{DrawingArea: da, progress: 0}
	if len(accentHex) == 7 && accentHex[0] == '#' {
		var r, g, b uint8
		fmt.Sscanf(accentHex, "#%02x%02x%02x", &r, &g, &b)
		ap.accent = [3]float64{float64(r)/255, float64(g)/255, float64(b)/255}
	} else {
		ap.accent = [3]float64{0.88, 0.31, 0.16}
	}
	go func() {
		pb, err := gdkpixbuf.NewPixbufFromFileAtScale(path, w, h, true)
		if err == nil {
			glib.IdleAdd(func() {
				ap.pixbuf = pb
				ap.StartAnimation()
			})
		}
	}()
	da.SetDrawFunc(ap.Draw)
	return da
}

func (ap *AnimatedPicture) StartAnimation() {
	glib.TimeoutAdd(16, func() bool {
		ap.progress += 0.025
		ap.QueueDraw()
		return ap.progress < 1.8
	})
}

func (ap *AnimatedPicture) Draw(_ *gtk.DrawingArea, cr *cairo.Context, w, h int) {
	if ap.pixbuf == nil { return }
	
	fw, fh := float64(w), float64(h)
	tw, th := float64(ap.pixbuf.Width()), float64(ap.pixbuf.Height())
	tx, ty := (fw-tw)/2, (fh-th)/2

	gridSize := 32.0
	cols, rows := math.Ceil(fw/gridSize), math.Ceil(fh/gridSize)
	
	for c := 0.0; c < cols; c++ {
		for r := 0.0; r < rows; r++ {
			cx, cy := c*gridSize, r*gridSize
			dx, dy := (c - cols/2), (r - rows/2)
			dist := math.Sqrt(dx*dx + dy*dy) / (math.Sqrt(cols*cols+rows*rows) / 2)
			
			lp := (ap.progress - dist*0.7) / 0.4
			if lp <= 0 { continue }
			if lp > 1 { lp = 1 }

			cr.Save()
			cr.Rectangle(cx, cy, gridSize, gridSize)
			cr.Clip()

			if lp < 0.5 {
				cr.SetSourceRGBA(ap.accent[0], ap.accent[1], ap.accent[2], lp/0.5)
				cr.Paint()
			} else {
				fade := (lp - 0.5) / 0.5
				if cx+gridSize >= tx && cx <= tx+tw && cy+gridSize >= ty && cy <= ty+th {
					gdk.CairoSetSourcePixbuf(cr, ap.pixbuf, tx, ty)
					cr.PaintWithAlpha(fade)
				}
				cr.SetSourceRGBA(ap.accent[0], ap.accent[1], ap.accent[2], (1.0-fade)*0.8)
				cr.Paint()
			}
			cr.Restore()
		}
	}
	
	if ap.progress > 1.6 {
		gdk.CairoSetSourcePixbuf(cr, ap.pixbuf, tx, ty)
		cr.Paint()
	}
}

func SetupTheme(c Colors, extraCSS string) {
	css := gtk.NewCSSProvider()
	css.LoadFromData(BuildBaseCSS(c) + `
.lightbox-overlay { background-color: rgba(0, 0, 0, 0); transition: background-color 0.35s ease; }
.lightbox-overlay.active { background-color: rgba(0, 0, 0, 0.96); }
.lightbox-image { margin: 48px; opacity: 0; transform: scale(0.98); transition: all 0.45s cubic-bezier(0.2, 0.8, 0.2, 1); }
.lightbox-overlay.active .lightbox-image { opacity: 1; transform: scale(1); }
` + extraCSS)
	gtk.StyleContextAddProviderForDisplay(gdk.DisplayGetDefault(), css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)
}

func NewTag(label string) *gtk.Label {
	lbl := gtk.NewLabel(label); lbl.AddCSSClass("gallery-tag"); return lbl
}

type MenuAction struct {
	Label string; Icon string; Danger bool; OnClick func()
}
