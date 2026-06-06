//go:build !dev

package embed

import (
	"embed"
	"io/fs"
)

//go:embed web/dist
var distFS embed.FS

// DistFS exposes the embedded frontend files as fs.FS.
var DistFS fs.FS = func() fs.FS {
	sub, err := fs.Sub(distFS, "web/dist")
	if err != nil {
		return nil
	}
	return sub
}()
