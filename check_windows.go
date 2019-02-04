package neterr

import (
	"net"
	"os"
	"syscall"
)

const (
	// Windows Sockets Error Codes
	// https://msdn.microsoft.com/en-us/library/windows/desktop/ms740668.aspx
	WSAECONNREFUSED syscall.Errno = 10061
)

func CheckErr(err error) error {
	if err == nil {
		return nil
	}
	if opErr, ok := err.(*net.OpError); ok {
		if syscallErr, ok := opErr.Err.(*os.SyscallError); ok {
			if errNo, ok := syscallErr.Err.(syscall.Errno); ok {
				switch errNo {
				case WSAECONNREFUSED:
					return Refused
				}
			}
		}
	}
	return err
}
