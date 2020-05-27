//go:generate go run main.go
package main

import (
	"github.com/golangee/bundle"
)

func main() {
	err := bundle.Embed(bundle.Options{
		TargetDir:   ".",
		PackageName: "swaggerui",
		Include:     []string{"./internal/dist"},
		IgnoreRegex: `.*\.map|^\..*`,
		StripPrefixes: []string{"/internal/dist"},
	})

	if err != nil {
		panic(err)
	}
}
