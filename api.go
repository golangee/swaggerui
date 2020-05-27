package swaggerui

import (
	"github.com/golangee/bundle"
	"net/http"
)

// Handler returns a new handler which provides a swaggerui instance.
// Use path to create a prefix for the delivered resources and specDoc with the
// json openAPI specification.
func Handler(path string, specDoc string) func(http.ResponseWriter, *http.Request) {
	myBundle := Bundle.Put(bundle.NewResourceFromBytes("/apidoc.json", []byte(specDoc)))
	return myBundle.Handler(path)
}
