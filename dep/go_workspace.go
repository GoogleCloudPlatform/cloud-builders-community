// Copyright 2016 Google, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"go/build"
	"os"
)

func main() {
	// This command looks for an import directive in a go source file in the
	// current directory. https://golang.org/s/go14customimport has details of how
	// import directives are specified.
	pkg, err := build.ImportDir(".", build.ImportComment)
	if err != nil {
		fmt.Printf("Could not parse source in current directory: %v\n", err)
		os.Exit(1)
	}
	if pkg.ImportComment == "" {
		os.Exit(1)
	}
	fmt.Print(pkg.ImportComment)
}
