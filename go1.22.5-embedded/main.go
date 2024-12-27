// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.22.5-embedded command runs the go command from Go 1.22.5.
//
// To install, run:
//
//	$ go install github.com/clktmr/dl/go1.22.5-embedded@latest
//	$ go1.22.5-embedded download
//
// And then use the go1.22.5-embedded command as if it were your normal go
// command.
//
// See the release notes at https://go.dev/doc/devel/release#go1.22.5.
//
// File bugs at https://go.dev/issue/new.
package main

import "github.com/clktmr/dl/internal/version"

func main() {
	version.RunTag("go1.22.5-embedded")
}
