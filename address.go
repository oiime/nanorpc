package nanorpc

import (
	"fmt"
	"regexp"
)

type Address string

var addrRegex, _ = regexp.Compile("^(nano|xrb)_[13]{1}[13456789abcdefghijkmnopqrstuwxyz]{59}$")

func MustAddressFromString(orig string) Address {
	addr, err := AddressFromString(orig)
	if err != nil {
		panic(err)
	}
	return addr
}
func AddressFromString(orig string) (Address, error) {
	addr := Address(orig)
	if !addr.Valid() {
		return addr, fmt.Errorf("Invalid nano address: %s", orig)
	}
	return addr, nil
}

func (a Address) Valid() bool {
	return addrRegex.MatchString(string(a))
}

func (a Address) String() string {
	return string(a)
}
