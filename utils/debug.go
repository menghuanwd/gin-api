package utils

import (
	"github.com/davecgh/go-spew/spew"
)

// SpewDump ...
func SpewDump(body ...interface{}) {
	spew.Dump(body)
}
