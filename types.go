package nanorpc

import (
	"fmt"
	"math/big"
	"strconv"
)

type Raw big.Int

var nanoDecimalDivision, _ = big.NewInt(0).SetString("1000000000000000000000000000000", 10)

func (r *Raw) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	if string(s) == "null" {
		return
	}
	s, err = strconv.Unquote(s)
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		err = fmt.Errorf("Unable to parse raw: %v", s)
		return
	}
	*r = Raw(*n)
	return
}

func (r Raw) String() string {
	v := big.Int(r)
	return v.String()
}

func (r Raw) AsNanoString() string {
	v := big.Int(r)
	abs := big.NewInt(0).Abs(&v)
	integerDigits, fractionalDigits := big.NewInt(0).QuoRem(abs, nanoDecimalDivision, big.NewInt(0))

	if v.Sign() == -1 {
		return fmt.Sprintf("-%s.%030s", integerDigits.String(), fractionalDigits.String())
	}
	return fmt.Sprintf("%s.%030s", integerDigits.String(), fractionalDigits.String())
}
func (r Raw) AsNanoFloat() *big.Float {
	v := big.Int(r)
	f := new(big.Float).SetInt(&v)
	return f.Quo(f, new(big.Float).SetInt(nanoDecimalDivision))
}

type Block string

type BlockType string

const (
	BlockTypeSend    BlockType = "send"
	BlockTypeRecieve BlockType = "receive"
)
