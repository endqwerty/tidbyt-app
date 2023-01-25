// Package airthings provides details for the AirThings applet.
package tidbytapp

import (
	_ "embed"

	"tidbyt.dev/community/apps/manifest"
)

//go:embed airthings.star
var source []byte

// New creates a new instance of the AirThings applet.
func New() manifest.Manifest {
	return manifest.Manifest{
		ID:          "tidbytapptest",
		Name:        "tidbytapptest",
		Author:      "endqwerty",
		Summary:     "tidbyt app test",
		Desc:        "Test Description",
		FileName:    "tidbyt-app.star",
		PackageName: "tidbytapptest",
		Source:      source,
	}
}
