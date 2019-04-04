// +build windows

package nf

import (
	"fmt"
	"syscall"
	"unsafe"
)

type ASIODriver struct {
	Name  string
	CLSID string
	GUID  *GUID

	ASIO *IASIO
}

func (drv *ASIODriver) Open() (err error) {
	disp, err := CreateInstance(drv.GUID, drv.GUID)
	if err != nil {
		return
	}
	drv.ASIO = (*IASIO)(unsafe.Pointer(disp))

	//drv.ASIO.AsIUnknown().AddRef()

	ok := drv.ASIO.Init(uintptr(0))
	if !ok {
		return fmt.Errorf("Could not init ASIO driver")
	}

	return
}

func (drv *ASIODriver) Close() {
	drv.ASIO.AsIUnknown().Release()
}

func newDriver(key syscall.Handle, keynameUTF16 winUTF16string) (drv *ASIODriver, err error) {
	var subkey syscall.Handle
	err = syscall.RegOpenKeyEx(key, keynameUTF16.Addr(), 0, syscall.KEY_READ, &subkey)
	if err != nil {
		return nil, err
	}
	defer syscall.RegCloseKey(subkey)

	clsidName, err := syscall.UTF16PtrFromString("clsid")
	if err != nil {
		return nil, err
	}

	// Get CLSID of driver impl:
	clsidUTF16, datatype, datasize := make([]uint16, 128, 128), uint32(syscall.REG_SZ), uint32(256)
	err = syscall.RegQueryValueEx(subkey, clsidName, nil, &datatype, (*byte)(unsafe.Pointer(&clsidUTF16[0])), &datasize)
	if err != nil {
		return nil, err
	}

	// Convert the subkey name from UTF-16 to a string:
	keyname := keynameUTF16.String()
	drv = &ASIODriver{
		Name:  keyname,
		CLSID: syscall.UTF16ToString(clsidUTF16),
	}

	drv.GUID, err = CLSIDFromStringUTF16(&clsidUTF16[0])
	if err != nil {
		return nil, err
	}

	return drv, nil
}

// Enumerate list of ASIO drivers registered on the system
func ListDrivers() (drivers map[string]*ASIODriver, err error) {
	var key syscall.Handle
	key, err = RegOpenKey(syscall.HKEY_LOCAL_MACHINE, "Software\\ASIO", syscall.KEY_ENUMERATE_SUB_KEYS)
	if err != nil {
		return
	}
	defer syscall.RegCloseKey(key)

	drivers = make(map[string]*ASIODriver)

	// Enumerate subkeys:
	index := uint32(0)
	for err == nil {
		keynameUTF16 := winUTF16string{
			utf16:  make([]uint16, 128),
			length: uint32(128),
		}

		// Get next subkey:
		err = syscall.RegEnumKeyEx(key, index, keynameUTF16.Addr(), &keynameUTF16.length, nil, nil, nil, nil)
		// Determine when to stop:
		if err != nil {
			if errno, ok := err.(syscall.Errno); ok {
				// 259 is "No more data" error; aka end of enumeration.
				if uintptr(errno) == uintptr(259) {
					err = nil
					break
				}
			}
			fmt.Println(err)
			return
		}

		index++

		// Create an ASIODriver based on the key:
		drv, err := newDriver(key, keynameUTF16)
		if err != nil {
			continue
		}

		drivers[drv.Name] = drv
	}

	return drivers, nil
}
