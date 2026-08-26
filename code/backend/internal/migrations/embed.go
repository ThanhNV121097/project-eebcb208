package migrations

import "embed"

// Files contains SQL migrations embedded for runtime startup.
//go:embed *.sql
var Files embed.FS
