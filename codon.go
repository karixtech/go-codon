// Copyright © 2017 Grofers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "github.com/grofers/go-codon/cmd"

// All the generate directives here
//go:generate go-bindata -prefix bootstrap/golang/content/ -pkg golang -o bootstrap/golang/bindata.go bootstrap/golang/content/...
//go:generate go-bindata -prefix generator/golang/content/ -pkg golang -o generator/golang/bindata.go generator/golang/content/...

func main() {
	cmd.Execute()
}
