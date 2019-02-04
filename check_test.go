package neterr

import (
	"net"
	"testing"
)

func connectionRefused() error {
	_, err := net.Dial("tcp", "localhost:62135")
	return err
}

func TestCheckErr(t *testing.T) {
	cases := []struct {
		f    func() error
		want error
	}{
		{connectionRefused, Refused},
	}
	for _, c := range cases {
		got := CheckErr(c.f())
		if got != c.want {
			t.Error(c.want)
		}
	}
}
