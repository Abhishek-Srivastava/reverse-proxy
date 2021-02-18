// +build tools

package tools

// this version locks the tooling to a release
import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "golang.org/x/tools/cmd/goimports"
)
