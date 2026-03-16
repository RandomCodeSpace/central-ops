// Package ui provides shared UI assets (CSS, fonts, icons).
package ui

import _ "embed"

//go:embed theme.css
var themeCSS []byte

// ThemeCSS returns the shared dark-theme CSS bytes.
func ThemeCSS() []byte { return themeCSS }
