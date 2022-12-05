//go:build tools
// +build tools

package wbserver

// A list of all the tooling packages wandbctl depends on.
// Included here as explicit imports to work cleanly with
// go modules.
import (
	_ "github.com/gordonklaus/ineffassign"
	_ "github.com/segmentio/golines"
	_ "golang.org/x/tools/cmd/goimports"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
