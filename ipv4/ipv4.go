package ipv4

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Atoi(address string) (integer uint32, err error) {
	if len(address) > 15 {
		return integer, errors.New("invalid address string")
	}

	octs := strings.Split(address, ".")

	if len(octs) != 4 {
		return integer, errors.New("4 octects required")
	}

	for i := 0; i < 4; i++ {
		oct, err := strconv.Atoi(octs[i])

		if err != nil {
			return integer, fmt.Errorf("bad octect: %s", octs[i])
		}

		integer += uint32(oct) << uint32((4 - 1 - i) * 8)
	}
	return
}

func Itoa(integer uint32) (address string, err error) {
	var buf bytes.Buffer

	for i := 0; i < 4; i++ {
		oct := (integer >> uint32((4 - 1 - i) * 8)) & 0xff
		buf.WriteString(strconv.Itoa(int(oct)))

		if i < 3 {
			buf.WriteByte(46)
		}

	}
	return buf.String(), err
}
