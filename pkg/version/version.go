// Package version provides build-info version detection.
package version

import "runtime/debug"

// Detect returns the module version from build info, or "local" for dev builds.
// It returns the real tag when installed via `go install module@vX.Y.Z`,
// and "local" for `go run` or local builds.
func Detect() string {
	if info, ok := debug.ReadBuildInfo(); ok &&
		info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return "local"
}
