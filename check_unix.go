// +build darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package neterr

import (
	"net"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func CheckErr(err error) error {
	if err == nil {
		return nil
	}
	if opErr, ok := err.(*net.OpError); ok {
		if syscallErr, ok := opErr.Err.(*os.SyscallError); ok {
			if errNo, ok := syscallErr.Err.(syscall.Errno); ok {
				switch errNo {
				case unix.ECONNREFUSED:
					return Refused
				case unix.EHOSTUNREACH:
					return Unreachable
				}
			}
		}
	}
	return err
}
