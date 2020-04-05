package goavx

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const avxPath = "\\goavx"
const avxModule = "goavx"

var LoadedDLL *windows.DLL

func init() {
	loadWinDLL()
}

func initDllPath(kernel32 syscall.Handle) {
	setDllDirectory, err := syscall.GetProcAddress(kernel32, "SetDllDirectoryA")
	if err != nil {
		// we can't do anything since SetDllDirectoryA is missing - fall back to use first wpcap.dll we encounter
		return
	}
	getSystemDirectory, err := syscall.GetProcAddress(kernel32, "GetSystemDirectoryA")
	if err != nil {
		// we can't do anything since SetDllDirectoryA is missing - fall back to use first wpcap.dll we encounter
		return
	}
	buf := make([]byte, 4096)
	r, _, _ := syscall.Syscall(getSystemDirectory, 2, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)), 0)
	if r == 0 {
		// we can't do anything since SetDllDirectoryA is missing
		return
	}
	copy(buf[r:], avxPath)
	_, _, _ = syscall.Syscall(setDllDirectory, 1, uintptr(unsafe.Pointer(&buf[0])), 0, 0)
	// ignore errors here - we just fallback to load wpcap.dll from default locations
}

func loadWinDLL() error {
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		return fmt.Errorf("couldn't load kernel32.dll")
	}
	defer syscall.FreeLibrary(kernel32)
	initDllPath(kernel32)

	LoadedDLL = windows.MustLoadDLL(avxModule)

	return nil
}

func byteSliceToString(bval []byte) string {
	for i := range bval {
		if bval[i] == 0 {
			return string(bval[:i])
		}
	}
	return string(bval[:])
}

func bytePtrToString(r uintptr) string {
	if r == 0 {
		return ""
	}
	bval := (*[1 << 30]byte)(unsafe.Pointer(r))
	return byteSliceToString(bval[:])
}
