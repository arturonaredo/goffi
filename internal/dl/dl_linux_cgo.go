//go:build linux && !android && cgo

package dl

import (
	"fmt"

	"github.com/go-webgpu/goffi/internal/cbridge"
)

// RTLD constants from <dlfcn.h> for dynamic library loading on Linux.
const (
	// RTLD_LAZY performs relocations at an implementation-dependent time.
	RTLD_LAZY = 0x00001

	// RTLD_NOW resolves all symbols when loading the library (recommended).
	RTLD_NOW = 0x00002

	// RTLD_GLOBAL makes all symbols available for relocation processing of other modules.
	// NOTE: Different from macOS (0x8) - Linux uses 0x00100
	RTLD_GLOBAL = 0x00100

	// RTLD_LOCAL makes symbols not available for relocation processing by other modules.
	RTLD_LOCAL = 0x00000
)

// RTLD_DEFAULT is a pseudo-handle for dlsym to search for any loaded symbol.
// NOTE: Different from macOS (1<<64 - 2) - Linux uses 0
const RTLD_DEFAULT = 0x00000

// Dlopen loads a shared library.
func Dlopen(path string, mode int) (uintptr, error) {
	handle, errMessage := cbridge.Dlopen(path, mode)
	if handle == 0 {
		if errMessage == "" {
			errMessage = "unknown error"
		}
		return 0, fmt.Errorf("dlopen failed: %s", errMessage)
	}
	return handle, nil
}

// Dlsym returns the address of a symbol in a loaded library.
func Dlsym(handle uintptr, name string) (uintptr, error) {
	symbol, errMessage := cbridge.Dlsym(handle, name)
	if symbol == 0 {
		if errMessage == "" {
			errMessage = "unknown error"
		}
		return 0, fmt.Errorf("dlsym failed: %s", errMessage)
	}
	return symbol, nil
}

// Dlclose intentionally retains process-lifetime mappings, matching the
// existing implementation's semantics.
func Dlclose(uintptr) error { return nil }
