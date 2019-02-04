package neterr

type netError string

var Refused = netError("connection refused")
var Unreachable = netError("destination is unreachable")

func (err netError) Error() string {
	return string(err)
}
