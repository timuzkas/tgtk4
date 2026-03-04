package tgtk4

import (
	"math"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Scroller struct {
	Scroll       *gtk.ScrolledWindow
	Velocity     float64
	Overshoot    float64
	InertiaTimer glib.SourceHandle
	MarginTarget *gtk.Widget // The widget whose margins are modified for overshoot
}

func NewSmoothScroller(scroll *gtk.ScrolledWindow, target *gtk.Widget) *Scroller {
	s := &Scroller{
		Scroll:       scroll,
		MarginTarget: target,
	}

	ctrl := gtk.NewEventControllerScroll(gtk.EventControllerScrollVertical)
	ctrl.ConnectScroll(func(dx, dy float64) bool {
		s.Velocity += dy * 14
		if s.InertiaTimer == 0 {
			s.InertiaTimer = glib.TimeoutAdd(16, func() bool {
				adj := s.Scroll.VAdjustment()
				min := adj.Lower()
				max := adj.Upper() - adj.PageSize()
				current := adj.Value()

				if (current <= min && s.Velocity < 0) || (current >= max && s.Velocity > 0) {
					s.Overshoot -= s.Velocity * 0.4
					s.Velocity *= 0.6
				}

				if math.Abs(s.Velocity) < 0.1 && math.Abs(s.Overshoot) < 0.5 {
					s.InertiaTimer = 0
					s.Velocity = 0
					s.Overshoot = 0
					if s.MarginTarget != nil {
						s.MarginTarget.SetMarginTop(0)
						s.MarginTarget.SetMarginBottom(0)
					}
					return false
				}

				adj.SetValue(current + s.Velocity)
				s.Velocity *= 0.90

				if math.Abs(s.Overshoot) > 0 && s.MarginTarget != nil {
					if s.Overshoot > 0 {
						s.MarginTarget.SetMarginTop(int(s.Overshoot))
						s.MarginTarget.SetMarginBottom(0)
					} else {
						s.MarginTarget.SetMarginTop(0)
						s.MarginTarget.SetMarginBottom(int(-s.Overshoot))
					}
					s.Overshoot *= 0.8
				}

				return true
			})
		}
		return true
	})
	scroll.AddController(ctrl)

	return s
}
