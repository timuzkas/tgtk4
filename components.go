package tgtk4

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

// ─── DropZone ────────────────────────────────────────────────────────────────

type DropZone struct {
	*gtk.Box
	Overlay *gtk.Overlay
	Hint    *gtk.Label
	OnDrop  func(paths []string)
}

func NewDropZone(text string, onDrop func([]string)) *DropZone {
	overlay := gtk.NewOverlay()
	overlay.SetHExpand(true)
	overlay.SetVExpand(true)
	overlay.AddCSSClass("drop-zone-container")

	innerBox := gtk.NewBox(gtk.OrientationVertical, 0)
	innerBox.AddCSSClass("drop-zone")
	innerBox.SetSizeRequest(300, 300)
	innerBox.SetHAlign(gtk.AlignCenter)
	innerBox.SetVAlign(gtk.AlignCenter)

	hint := gtk.NewLabel(text)
	hint.AddCSSClass("hint")
	hint.SetVExpand(true)
	hint.SetHAlign(gtk.AlignCenter)
	innerBox.Append(hint)
	overlay.SetChild(innerBox)

	// Hover overlay
	hoverOverlay := gtk.NewBox(gtk.OrientationVertical, 0)
	hoverOverlay.AddCSSClass("drop-zone-hover")
	hoverOverlay.SetHAlign(gtk.AlignFill)
	hoverOverlay.SetVAlign(gtk.AlignFill)

	centerBox := gtk.NewBox(gtk.OrientationVertical, 0)
	centerBox.SetHAlign(gtk.AlignCenter)
	centerBox.SetVAlign(gtk.AlignCenter)
	centerBox.AddCSSClass("hover-box")
	centerBox.SetVExpand(true)

	icon := gtk.NewImageFromIconName("document-open-symbolic")
	icon.SetPixelSize(32)
	icon.SetMarginBottom(16)
	centerBox.Append(icon)

	hoverLbl := gtk.NewLabel("select or drop files")
	hoverLbl.AddCSSClass("hover-label")
	centerBox.Append(hoverLbl)

	hoverOverlay.Append(centerBox)
	overlay.AddOverlay(hoverOverlay)

	dz := &DropZone{
		Box:     gtk.NewBox(gtk.OrientationVertical, 0),
		Overlay: overlay,
		Hint:    hint,
		OnDrop:  onDrop,
	}
	dz.Append(overlay)

	target := gtk.NewDropTarget(glib.TypeString, gdk.ActionCopy)
	target.ConnectDrop(func(val *glib.Value, x, y float64) bool {
		str := val.String()
		lines := strings.Split(str, "\r\n")
		var paths []string
		for _, l := range lines {
			l = strings.Trim(l, "\"' \r\n")
			if l == "" {
				continue
			}
			if strings.HasPrefix(l, "file://") {
				if u, err := url.Parse(l); err == nil {
					l = u.Path
				} else {
					l = strings.TrimPrefix(l, "file://")
				}
			}
			if _, err := os.Stat(l); err == nil {
				paths = append(paths, l)
			}
		}
		if len(paths) > 0 && dz.OnDrop != nil {
			dz.OnDrop(paths)
			return true
		}
		return false
	})
	overlay.AddController(target)

	return dz
}

func (dz *DropZone) SetActive(active bool) {
	innerBox := dz.Overlay.Child().(*gtk.Box)
	if active {
		innerBox.AddCSSClass("active")
	} else {
		innerBox.RemoveCSSClass("active")
	}
}

// ─── LogView ─────────────────────────────────────────────────────────────────

type LogView struct {
	*gtk.ScrolledWindow
	List *gtk.ListBox
}

func NewLogView() *LogView {
	scroll := gtk.NewScrolledWindow()
	scroll.SetPolicy(gtk.PolicyNever, gtk.PolicyAutomatic)
	scroll.SetMinContentHeight(100)

	list := gtk.NewListBox()
	list.AddCSSClass("log-list")
	list.SetSelectionMode(gtk.SelectionNone)
	scroll.SetChild(list)

	lv := &LogView{
		ScrolledWindow: scroll,
		List:           list,
	}
	lv.SetVExpand(true)
	NewSmoothScroller(scroll, &list.Widget)
	return lv
}

func (lv *LogView) Log(level LogLevel, text string) {
	ts := time.Now().Format("15:04:05")
	var prefix, cssClass string
	switch level {
	case LogInfo:
		prefix = "·"; cssClass = "log-info"
	case LogWarn:
		prefix = "▲"; cssClass = "log-warn"
	case LogErr:
		prefix = "✗"; cssClass = "log-err"
	case LogOK:
		prefix = "✓"; cssClass = "log-ok"
	}
	full := fmt.Sprintf("%s  %s %s", ts, prefix, text)

	row := gtk.NewListBoxRow()
	row.AddCSSClass("log-row")
	lbl := gtk.NewLabel(full)
	lbl.AddCSSClass("log-entry")
	lbl.AddCSSClass(cssClass)
	lbl.SetHAlign(gtk.AlignStart)
	lbl.SetXAlign(0)
	lbl.SetWrap(true)
	lbl.SetMaxWidthChars(80)
	row.SetChild(lbl)
	row.AddCSSClass("pop-in")
	lv.List.Prepend(row)
}

func (lv *LogView) Clear() {
	for child := lv.List.FirstChild(); child != nil; child = lv.List.FirstChild() {
		lv.List.Remove(child)
	}
}

type LogLevel int

const (
	LogInfo LogLevel = iota
	LogWarn
	LogErr
	LogOK
)

// ─── SidePanel ───────────────────────────────────────────────────────────────

type SidePanel struct {
	*gtk.Box
	Visible bool
}

func NewSidePanel(title string) *SidePanel {
	root := gtk.NewBox(gtk.OrientationVertical, 0)
	root.AddCSSClass("side-panel")
	root.SetHAlign(gtk.AlignEnd)
	root.SetVAlign(gtk.AlignFill)

	titleRow := gtk.NewBox(gtk.OrientationHorizontal, 0)
	titleRow.SetMarginTop(20)
	titleRow.SetMarginBottom(16)
	titleRow.SetMarginStart(16)
	titleRow.SetMarginEnd(16)

	titleLbl := gtk.NewLabel("// " + title)
	titleLbl.AddCSSClass("settings-title")
	titleRow.Append(titleLbl)
	root.Append(titleRow)

	sep := gtk.NewSeparator(gtk.OrientationHorizontal)
	root.Append(sep)

	return &SidePanel{Box: root}
}

func (p *SidePanel) Toggle() {
	p.Visible = !p.Visible
	if p.Visible {
		p.SetVisible(true)
		glib.TimeoutAdd(10, func() bool {
			p.AddCSSClass("visible")
			return false
		})
	} else {
		p.RemoveCSSClass("visible")
		glib.TimeoutAdd(500, func() bool {
			if !p.Visible {
				p.SetVisible(false)
			}
			return false
		})
	}
}

// ─── Fundamental UI Blocks ──────────────────────────────────────────────────

func NewCheck(label string) *gtk.CheckButton {
	cb := gtk.NewCheckButtonWithLabel(label)
	cb.AddCSSClass("cyber-check")
	return cb
}

func NewSwitch() *gtk.Switch {
	s := gtk.NewSwitch()
	s.AddCSSClass("cyber-switch")
	s.SetVAlign(gtk.AlignCenter)
	return s
}

func NewSettingsRow(label string, widget gtk.Widgetter, vertical bool) *gtk.Box {
	var row *gtk.Box
	if vertical {
		row = gtk.NewBox(gtk.OrientationVertical, 4)
	} else {
		row = gtk.NewBox(gtk.OrientationHorizontal, 12)
	}
	row.AddCSSClass("setting-row")

	lbl := gtk.NewLabel(label)
	lbl.AddCSSClass("setting-label")
	lbl.SetHAlign(gtk.AlignStart)
	row.Append(lbl)

	if !vertical {
		spacer := gtk.NewBox(gtk.OrientationHorizontal, 0)
		spacer.SetHExpand(true)
		row.Append(spacer)
	}

	row.Append(widget)
	return row
}

// ─── AdaptiveBox ─────────────────────────────────────────────────────────────

type AdaptiveBox struct {
	*gtk.Box
	Breakpoint int
}

func NewAdaptiveBox(breakpoint int) *AdaptiveBox {
	box := gtk.NewBox(gtk.OrientationHorizontal, 0)
	box.SetHExpand(true)
	box.SetVExpand(true)
	ab := &AdaptiveBox{Box: box, Breakpoint: breakpoint}

	box.Connect("size-allocate", func() {
		w := box.Width()
		if w < ab.Breakpoint {
			if box.Orientation() == gtk.OrientationHorizontal {
				box.SetOrientation(gtk.OrientationVertical)
			}
		} else {
			if box.Orientation() == gtk.OrientationVertical {
				box.SetOrientation(gtk.OrientationHorizontal)
			}
		}
	})

	return ab
}

// ─── Animation Loop Helper ──────────────────────────────────────────────────

type Canvas struct {
	*gtk.DrawingArea
	OnDraw func(cr *cairo.Context, w, h int, phase float64)
	Phase  float64
	Timer  glib.SourceHandle
}

func NewCanvas(size int, onDraw func(*cairo.Context, int, int, float64)) *Canvas {
	da := gtk.NewDrawingArea()
	if size > 0 {
		da.SetSizeRequest(size, size)
	}
	c := &Canvas{DrawingArea: da, OnDraw: onDraw}

	da.SetDrawFunc(func(_ *gtk.DrawingArea, cr *cairo.Context, w, h int) {
		if c.OnDraw != nil {
			c.OnDraw(cr, w, h, c.Phase)
		}
	})

	return c
}

func (c *Canvas) Start() {
	if c.Timer != 0 {
		return
	}
	c.Timer = glib.TimeoutAdd(16, func() bool {
		c.Phase += 0.05
		c.QueueDraw()
		return true
	})
}

func (c *Canvas) Stop() {
	if c.Timer != 0 {
		glib.SourceRemove(c.Timer)
		c.Timer = 0
	}
}
