//go:build dev

package embed

import "io/fs"

// DistFS is a no-op placeholder when building with -tags dev
// (i.e. during development when web/dist does not exist yet).
var DistFS fs.FS
