package tgtk4

import "fmt"

type Colors struct {
	Bg     string `toml:"bg"`
	Bg2    string `toml:"bg2"`
	Fg     string `toml:"fg"`
	Muted  string `toml:"muted"`
	Accent string `toml:"accent"`
	Border string `toml:"border"`
	Green  string `toml:"green"`
	Dim    string `toml:"dim"`
	Danger string `toml:"danger"`
}

func DefaultColors() Colors {
	return Colors{
		Bg:     "#0d0c0a",
		Bg2:    "#141210",
		Fg:     "#e2ddd4",
		Muted:  "#6b6560",
		Accent: "#e04e2a",
		Border: "#222018",
		Green:  "#7a9e6e",
		Dim:    "#3a3530",
		Danger: "#c9391a",
	}
}

func EnsureHex(s string) string {
	if s == "" {
		return "#000000"
	}
	if s[0] != '#' {
		return "#" + s
	}
	return s
}

func BuildBaseCSS(c Colors) string {
	return fmt.Sprintf(`
/* Global Focus Reset - NO BLUE EVER */
* {
  outline: none;
  outline-width: 0;
  box-shadow: none;
  -gtk-outline-radius: 0;
}

*:focus {
  outline: none;
  box-shadow: none;
}

window {
  background-color: %[1]s;
}

headerbar {
  background-color: rgba(0, 0, 0, 0.4);
  background-image: none;
  border-bottom: 1px solid %[6]s;
  min-height: 64px;
  padding: 0 16px;
}

headerbar label.title {
  color: %[5]s;
  font-family: "JetBrains Mono", monospace;
  font-weight: 700;
  font-size: 10px;
  letter-spacing: 0.15em;
}

/* Slightly Darker Buttons */
.action-btn, .conv-btn, .mini-action-btn, button {
  background-color: #0c0c0c;
  background-image: none;
  border: 1px solid %[6]s;
  border-radius: 0;
  font-family: "JetBrains Mono", monospace;
  font-size: 9px;
  font-weight: 600;
  color: %[4]s;
  padding: 6px 12px;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: none;
  outline: none;
}

.action-btn:hover, .mini-action-btn:hover, button:hover {
  color: %[3]s;
  background-color: #121212;
  background-image: none;
  border-color: %[5]s;
  box-shadow: none;
}

.action-btn:active, button:active {
  background-color: #080808;
  transform: scale(0.98);
}

dropdown {
  background-color: #0c0c0c;
  background-image: none;
  border: 1px solid %[6]s;
  border-radius: 0;
  font-family: "JetBrains Mono", monospace;
  font-size: 9px;
  font-weight: 600;
  color: %[3]s;
  padding: 0;
  margin: 0 4px;
}

dropdown:hover {
  border-color: %[5]s;
  background-color: #121212;
}

dropdown > button {
  padding: 4px 12px;
  background: none;
  background-image: none;
  border: none;
  box-shadow: none;
}

popover > contents {
  background-color: %[2]s;
  border: 1px solid %[6]s;
  border-radius: 0;
  padding: 0;
  box-shadow: none;
}

popover listview, popover listitem {
  background-color: %[2]s;
  color: %[4]s;
}

popover listitem:hover {
  background-color: %[5]s;
  color: %[1]s;
}

checkbutton {
  color: %[4]s;
  font-family: "JetBrains Mono", monospace;
  font-size: 9px;
  font-weight: 600;
}

checkbutton check {
  -gtk-icon-source: none;
  background-image: none;
  background-color: #0c0c0c;
  border: 1px solid %[6]s;
  border-radius: 0;
  margin-right: 8px;
  min-width: 16px;
  min-height: 16px;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  outline: none;
}

checkbutton:checked check {
  background-color: %[5]s;
  background-image: none;
  border-color: %[5]s;
}

progressbar trough {
  background-color: rgba(0, 0, 0, 0.2);
  min-height: 4px;
  min-width: 0;
  border-radius: 0;
}

progressbar progress {
  background-color: %[5]s;
  background-image: none;
  border-radius: 0;
  border: none;
  min-height: 0;
  min-width: 0;
}

/* Button-as-Progressbar styles */
.conv-btn {
  padding: 0 !important;
  overflow: hidden;
}

.conv-btn overlay {
  min-width: 140px;
  min-height: 44px;
}

.conv-btn progressbar {
  transition: all 0.2s ease;
}

.conv-btn progressbar trough {
  background-color: transparent;
  min-height: 44px;
}

.conv-btn progressbar progress {
  min-height: 44px;
  background-color: %[5]s;
}

.conv-btn label {
  color: %[5]s;
  z-index: 10;
  /* Use a bit of shadow for legibility if needed, or just switch color */
}

.conv-btn.processing label {
  color: #ffffff;
  text-shadow: 0 1px 2px rgba(0,0,0,0.8);
}

progressbar {
  transition: opacity 0.8s ease-in-out;
  opacity: 1;
}

progressbar.faded {
  opacity: 0;
}

scale {
  outline: none;
  box-shadow: none;
}

scale trough {
  background-color: #0c0c0c;
  background-image: none;
  min-height: 4px;
  min-width: 0;
  border: 1px solid %[6]s;
  border-radius: 0;
  outline: none;
}

scale highlight, scale fill, scale progress {
  background-color: %[5]s;
  background-image: none;
  border-radius: 0;
  border: none;
  min-height: 0;
  min-width: 0;
}

scale slider {
  background-color: %[3]s;
  background-image: none;
  border: 1px solid %[6]s;
  min-width: 12px;
  min-height: 12px;
  margin: -4px;
  border-radius: 0;
  box-shadow: none;
  outline: none;
}

scale:hover slider {
  background-color: %[5]s;
  border-color: %[5]s;
}

.status-bar {
  background-color: %[2]s;
  border-top: 1px solid %[6]s;
  padding: 4px 16px;
}

.status-bar label {
  font-family: "JetBrains Mono", monospace;
  font-size: 8px;
  font-weight: 500;
  color: %[4]s;
}

.drop-zone {
  background-color: rgba(255, 255, 255, 0.005);
  border: 1px solid %[6]s;
  border-radius: 0;
  transition: all 0.3s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.drop-zone:hover {
  background-color: rgba(255, 255, 255, 0.015);
  border-color: %[4]s;
}

.drop-zone.active {
  background-color: rgba(224, 78, 42, 0.01);
  border-color: %[5]s;
}

.overlay-actions {
  background-color: #0c0c0c;
  border: 1px solid %[6]s;
  padding: 2px;
  opacity: 0;
  transition: all 0.25s cubic-bezier(0.2, 0.8, 0.2, 1);
}

overlay:hover .overlay-actions {
  opacity: 1;
}

.mini-action-btn {
  border: 1px solid transparent;
  padding: 4px 12px;
  font-family: "JetBrains Mono", monospace;
  font-size: 9px;
  font-weight: 600;
  color: %[4]s;
}

.mini-action-btn:hover {
  color: %[5]s;
  background-color: #121212;
  border-color: %[6]s;
}

.pop-in {
  animation: pop-in 0.4s cubic-bezier(0.2, 0.8, 0.2, 1);
}

@keyframes pop-in {
  0% { transform: scale(0.97); opacity: 0; }
  100% { transform: scale(1); opacity: 1; }
}

.ripple-react {
  animation: ripple-react 0.6s cubic-bezier(0.2, 0.8, 0.2, 1);
}

@keyframes ripple-react {
  0% { transform: scale(1); }
  20% { transform: scale(1.002); }
  100% { transform: scale(1); }
}

scrollbar slider {
  background-color: %[8]s;
  border-radius: 0;
  min-width: 4px;
}

tooltip {
  background-color: %[2]s;
  border: 1px solid %[6]s;
  border-radius: 0;
  padding: 4px 8px;
}

tooltip label {
  font-family: "JetBrains Mono", monospace;
  font-size: 10px;
  color: %[3]s;
}

headerbar button {
  min-width: 36px;
  min-height: 36px;
  border-radius: 0;
}

headerbar button:not(.titlebutton) {
  padding: 0;
}

.action-btn {
  min-width: 36px;
  min-height: 36px;
  padding: 0;
}

.action-btn > * {
  min-width: 20px;
  min-height: 20px;
}
`,
		EnsureHex(c.Bg), EnsureHex(c.Bg2), EnsureHex(c.Fg),
		EnsureHex(c.Muted), EnsureHex(c.Accent), EnsureHex(c.Border),
		EnsureHex(c.Green), EnsureHex(c.Dim), EnsureHex(c.Danger),
	)
}
