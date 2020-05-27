// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
