package esbuild

import (
	"syscall"
	"unsafe"
)

func getTerminalWidth() int {
	var result struct {
		_   uint16
		Col uint16
		_   uint16
		_   uint16
	}
	returnCode, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&result)),
	)
	if int(returnCode) == -1 {
		panic(errno)
	}
	return int(result.Col)
}
