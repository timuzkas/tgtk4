package tgtk4

import (
	"time"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type App struct {
	ID     string
	Title  string
	Config Colors
	Win    *gtk.ApplicationWindow
	GtkApp *gtk.Application

	StatusLabel     *gtk.Label
	StatusContainer *gtk.Box
}

func NewApp(id, title string) *App {
	return &App{
		ID:     id,
		Title:  title,
		Config: DefaultColors(),
	}
}

func (a *App) NewStatusBar() *gtk.Box {
	a.StatusContainer = gtk.NewBox(gtk.OrientationHorizontal, 0)
	a.StatusContainer.AddCSSClass("status-bar")
	a.StatusLabel = gtk.NewLabel("")
	a.StatusLabel.SetHAlign(gtk.AlignStart)
	a.StatusContainer.Append(a.StatusLabel)
	return a.StatusContainer
}

func (a *App) SetStatus(msg string, isError bool) {
	if a.StatusLabel != nil {
		a.StatusLabel.SetText("  " + msg)
		if isError {
			a.StatusContainer.AddCSSClass("error")
		} else {
			a.StatusContainer.RemoveCSSClass("error")
		}

		go func() {
			time.Sleep(5 * time.Second)
			glib.IdleAdd(func() {
				a.StatusContainer.RemoveCSSClass("error")
				a.StatusLabel.SetText("")
			})
		}()
	}
}

func (a *App) Run(onActivate func()) int {
	a.GtkApp = gtk.NewApplication(a.ID, 0)
	a.GtkApp.ConnectActivate(func() {
		SetupTheme(a.Config, "")
		a.Win = gtk.NewApplicationWindow(a.GtkApp)
		a.Win.SetTitle(a.Title)
		a.Win.SetDefaultSize(1100, 700)

		if onActivate != nil {
			onActivate()
		}

		a.Win.Show()
	})

	return a.GtkApp.Run(nil)
}
