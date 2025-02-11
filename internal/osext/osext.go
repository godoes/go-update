// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package osext Extensions to the standard "os" package.
package osext

import "path/filepath"

// Executable returns an absolute path that can be used to
// re-invoke the current program.
// It may not be valid after the current program exits.
func Executable() (string, error) {
	p, err := executable()
	return filepath.Clean(p), err
}

// ExecutableFolder returns same path as Executable, returns just the folder
// path. Excludes the executable name and any trailing slash.
func ExecutableFolder() (string, error) {
	p, err := Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(p), nil
}
