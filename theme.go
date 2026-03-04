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
  transition-property: background-color, border-color, opacity, transform;
  transition-duration: 200ms;
}

*:focus {
  outline: none;
  box-shadow: none;
}

window {
  background-color: %[1]s;
}

headerbar {
  background-color: rgba(0, 0, 0, 0.6);
  background-image: none;
  border-bottom: 1px solid %[6]s;
  min-height: 48px;
  padding: 0 8px;
}

headerbar label.title {
  color: %[5]s;
  font-family: "JetBrains Mono", monospace;
  font-weight: 800;
  font-size: 11px;
  letter-spacing: 0.2em;
  text-transform: uppercase;
}

/* Cyber Buttons - PERFECTLY SQUARE & CENTERED */
.action-btn, .mini-action-btn, button {
  background-color: #0c0c0c;
  background-image: none;
  border: 1px solid %[6]s;
  border-radius: 0;
  font-family: "JetBrains Mono", monospace;
  font-size: 9px;
  font-weight: 600;
  color: %[4]s;
  padding: 0;
  margin: 0;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: none;
  outline: none;
  min-width: 36px;
  min-height: 36px;
}

.action-btn:hover, .mini-action-btn:hover, button:hover {
  color: %[3]s;
  background-color: #121212;
  border-color: %[5]s;
}

.action-btn:active, button:active {
  background-color: #080808;
  transform: scale(0.92);
}

/* Center icons in buttons */
button image {
  margin: 0;
  padding: 0;
}

headerbar button {
  margin: 4px;
}

/* Inputs */
entry {
  background-color: #080808;
  border: 1px solid %[6]s;
  border-radius: 0;
  color: %[3]s;
  font-family: "JetBrains Mono", monospace;
  font-size: 11px;
  padding: 12px 16px;
  caret-color: %[5]s;
  transition: all 0.3s cubic-bezier(0.2, 0.8, 0.2, 1);
}

entry:focus {
  border-color: %[5]s;
  background-color: #0c0c0c;
  box-shadow: 0 0 16px rgba(224, 78, 42, 0.08);
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
  border: 1px solid %[5]s;
  border-radius: 0;
  padding: 0;
  box-shadow: 0 16px 32px rgba(0,0,0,0.6);
  animation: popover-in 0.25s cubic-bezier(0.2, 0.8, 0.2, 1);
}

@keyframes popover-in {
  from { opacity: 0; transform: translateY(-10px); }
  to { opacity: 1; transform: translateY(0); }
}

popover listview, popover listitem {
  background-color: %[2]s;
  color: %[4]s;
}

popover listitem:hover {
  background-color: %[5]s;
  color: %[1]s;
}

/* Fundamental Cyber Widgets */
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

switch {
  background-color: #0c0c0c;
  border: 1px solid %[6]s;
  border-radius: 0;
  min-width: 32px;
  min-height: 16px;
  transition: all 0.2s ease;
}

switch:checked {
  background-color: %[5]s;
  border-color: %[5]s;
}

switch slider {
  background-color: %[4]s;
  border: 1px solid %[6]s;
  border-radius: 0;
  min-width: 14px;
  min-height: 14px;
  margin: 1px;
}

switch:checked slider {
  background-color: %[2]s;
}

/* Row Styling */
.setting-row {
  padding: 14px 16px;
  border-bottom: 1px solid %[6]s;
}

.setting-label {
  font-family: "JetBrains Mono", monospace;
  font-size: 8px;
  font-weight: 700;
  letter-spacing: 0.1em;
  color: %[4]s;
  text-transform: uppercase;
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
  padding: 6px 16px;
}

.status-bar label {
  font-family: "JetBrains Mono", monospace;
  font-size: 8px;
  font-weight: 600;
  color: %[4]s;
  letter-spacing: 0.05em;
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
  animation: pop-in 0.5s cubic-bezier(0.2, 0.8, 0.2, 1) both;
}

@keyframes pop-in {
  0% { transform: translateY(12px); opacity: 0; }
  100% { transform: translateY(0); opacity: 1; }
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
  border: 1px solid %[5]s;
  border-radius: 0;
  padding: 6px 10px;
}

tooltip label {
  font-family: "JetBrains Mono", monospace;
  font-size: 10px;
  color: %[3]s;
}

/* Side Panel / Settings Overlay */
.side-panel {
  background-color: rgba(13, 12, 10, 0.99);
  border-left: 1px solid %[6]s;
  min-width: 320px;
  padding-bottom: 24px;
  transform: translateX(100%%);
  transition: transform 0.5s cubic-bezier(0.2, 0.8, 0.2, 1), opacity 0.4s ease;
  opacity: 0;
  box-shadow: -16px 0 48px rgba(0,0,0,0.7);
}
.side-panel.visible {
  transform: translateX(0);
  opacity: 1;
}

/* Log Rows / Chat bubbles */
.log-list {
  background-color: transparent;
  padding: 24px 0;
}
.log-row {
  padding: 0;
  margin: 6px 24px;
  border: none;
  background: transparent;
}
.log-entry {
  font-family: "JetBrains Mono", monospace;
  font-size: 11px;
  font-weight: 400;
  color: %[4]s;
  padding: 14px 20px;
  line-height: 1.6;
  letter-spacing: 0.01em;
  border-radius: 2px;
  transition: all 0.3s ease;
}
.log-info { /* USER */
  color: %[3]s;
  background-color: rgba(255, 255, 255, 0.04);
  border-right: 4px solid %[5]s;
  margin-left: 64px;
}
.log-ok { /* AI */
  color: %[7]s;
  background-color: rgba(122, 158, 110, 0.06);
  border-left: 4px solid %[7]s;
  margin-right: 64px;
}
.log-err { /* ERROR */
  color: %[9]s;
  background-color: rgba(201, 57, 26, 0.08);
  border-left: 4px solid %[9]s;
}
.log-warn { /* SYSTEM / HINT */
  color: #e0b44a;
  font-size: 9px;
  text-transform: uppercase;
  letter-spacing: 0.12em;
  opacity: 0.8;
  margin-top: 16px;
  margin-bottom: 8px;
}
`,
		EnsureHex(c.Bg), EnsureHex(c.Bg2), EnsureHex(c.Fg),
		EnsureHex(c.Muted), EnsureHex(c.Accent), EnsureHex(c.Border),
		EnsureHex(c.Green), EnsureHex(c.Dim), EnsureHex(c.Danger),
	)
}
