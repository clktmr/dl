// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.24.0-n64 command compiles and runs the go command from Embedded Go
// 1.24.0
//
// To install, run:
//
//	$ go install github.com/clktmr/dl/go1.24.0-n64@latest
//	$ go1.24.0-n64 download
//
// And then use the go1.24.0-n64 command as if it were your normal go command.
package main

import "github.com/clktmr/dl/internal/version"

func main() {
	version.RunTag("go1.24.0-n64")
}
