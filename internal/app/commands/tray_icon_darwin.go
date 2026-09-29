//go:build darwin && cgo

package commands

import _ "embed"

// Menu bar template icons: a Nerd Font glyph (Font Awesome play/clock) rendered
// to a monochrome PNG with alpha. macOS ignores the color of a template image
// and recolors the shape for the current appearance, so only the alpha
// channel matters.
//
//go:embed assets/tray_play.png
var trayPlayIconPNG []byte

//go:embed assets/tray_idle.png
var trayIdleIconPNG []byte
