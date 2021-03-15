// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The gotip command compiles and runs the go command from the development tree.
//
// To install, run:
//
//     $ go get golang.org/dl/gotip
//     $ gotip download
//
// And then use the gotip command as if it were your normal go command.
//
// To update, run "gotip download" again.
// To download a specific CL, run "gotip download NUMBER".
package main

import (
	"golang.org/dl/internal/version"
)

func main() {
	version.RunTip()
}
