//go:build !bazel
// +build !bazel

package bazel

// This file contains stub implementations for non-bazel builds.
// See bazel.go for full documentation on the contracts of these functions.

// BuiltWithBazel returns true if this library was built with Bazel.
func BuiltWithBazel() bool {
	return false
}

// FindBinary is not implemented.
func FindBinary(pkg, name string) (string, bool) {
	panic("not build with Bazel")
}

// Runfile is not implemented.
func Runfile(string) (string, error) {
	panic("not built with Bazel")
}

// RunfilesPath is not implemented.
func RunfilesPath() (string, error) {
	panic("not built with Bazel")
}

// TestTmpDir is not implemented.
func TestTmpDir() string {
	panic("not built with Bazel")
}

// NewTmpDir is not implemented.
func NewTmpDir(prefix string) (string, error) {
	panic("not built with Bazel")
}

// RelativeTestTargetPath is not implemented.
func RelativeTestTargetPath() string {
	panic("not built with Bazel")
}

// SetGoEnv is not implemented.
func SetGoEnv() {
	panic("not built with Bazel")
}