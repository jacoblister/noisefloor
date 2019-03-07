package main

import (
	"syscall"
	"unsafe"
)

type winUTF16string struct {
	utf16  []uint16
	length uint32
}

func (utfstring *winUTF16string) String() string {
	return syscall.UTF16ToString(utfstring.utf16[:utfstring.length])
}

func (utfstring *winUTF16string) Addr() *uint16 {
	return &utfstring.utf16[0]
}

func RegOpenKey(key syscall.Handle, subkey string, desiredAccess uint32) (handle syscall.Handle, err error) {
	var subkeyUTF16 *uint16
	subkeyUTF16, err = syscall.UTF16PtrFromString(subkey)
	if err != nil {
		return syscall.InvalidHandle, err
	}

	err = syscall.RegOpenKeyEx(
		key,
		subkeyUTF16,
		uint32(0),
		desiredAccess,
		&handle,
	)
	return
}

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type pIUnknownVtbl struct {
	pQueryInterface uintptr
	pAddRef         uintptr
	pRelease        uintptr
}

type IUnknown struct {
	lpVtbl *pIUnknownVtbl
}

const (
	CLSCTX_INPROC_SERVER = 1
	CLSCTX_LOCAL_SERVER  = 4
)

var (
	ole32, _ = syscall.LoadLibrary("ole32.dll")

	procCoInitialize, _     = syscall.GetProcAddress(ole32, "CoInitialize")
	procCoUninitialize, _   = syscall.GetProcAddress(ole32, "CoUninitialize")
	procCoCreateInstance, _ = syscall.GetProcAddress(ole32, "CoCreateInstance")
	procCLSIDFromString, _  = syscall.GetProcAddress(ole32, "CLSIDFromString")

	IID_NULL = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
)

func CoInitialize(p uintptr) (hr uintptr) {
	hr, _, _ = syscall.Syscall(uintptr(procCoInitialize), 1, p, 0, 0)
	return
}

func CoUninitialize() {
	syscall.Syscall(uintptr(procCoUninitialize), 0, 0, 0, 0)
}

func CLSIDFromString(str string) (clsid *GUID, err error) {
	var guid GUID
	err = nil

	hr, _, _ := syscall.Syscall(uintptr(procCLSIDFromString), 2,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str))),
		uintptr(unsafe.Pointer(&guid)), 0)
	if hr != 0 {
		err = syscall.Errno(hr)
	}

	clsid = &guid
	return
}

func CLSIDFromStringUTF16(str *uint16) (clsid *GUID, err error) {
	var guid GUID
	err = nil

	hr, _, _ := syscall.Syscall(uintptr(procCLSIDFromString), 2,
		uintptr(unsafe.Pointer(str)),
		uintptr(unsafe.Pointer(&guid)), 0)
	if hr != 0 {
		err = syscall.Errno(hr)
	}

	clsid = &guid
	return
}

func CreateInstance(clsid *GUID, iid *GUID) (unk *IUnknown, err error) {
	hr, _, _ := syscall.Syscall6(uintptr(procCoCreateInstance), 5,
		uintptr(unsafe.Pointer(clsid)),
		0,
		CLSCTX_INPROC_SERVER,
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(&unk)),
		0)
	if hr != 0 {
		err = syscall.Errno(hr)
	}
	return
}

func (unk *IUnknown) AddRef() (r1 uintptr, err error) {
	var errno syscall.Errno
	r1, _, errno = syscall.Syscall(unk.lpVtbl.pAddRef, uintptr(1),
		uintptr(unsafe.Pointer(unk)),
		uintptr(0),
		uintptr(0))
	if errno != 0 {
		err = errno
	}
	return
}

func (unk *IUnknown) Release() (r1 uintptr, err error) {
	var errno syscall.Errno
	r1, _, errno = syscall.Syscall(unk.lpVtbl.pRelease, uintptr(1),
		uintptr(unsafe.Pointer(unk)),
		uintptr(0),
		uintptr(0))
	if errno != 0 {
		err = errno
	}
	return
}
