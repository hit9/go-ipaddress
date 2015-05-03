package ipv4

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Atoi(address string) (integer uint32, err error) {
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
	octs := []string {}

	for i := 0; i < 4; i++ {
		oct := (integer >> uint32((4 - 1 - i) * 8)) & 0xff
		octs = append(octs, strconv.Itoa(int(oct)))
	}

	fmt.Println(len(octs))

	return strings.Join(octs, "."), err
}
