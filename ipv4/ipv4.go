// IPv4 address utils for golang.
// Copyright 2015. Chao Wang <hit9@icloud.com>

package ipv4

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

// Atoi returns the uint32 representation of an ipv4 address string value.
// Example:
//
//	Atoi("192.168.0.1")   // 3232312315
//
func Atoi(address string) (uint32, error) {
	var sum uint32 = 0

	if len(address) > 15 {
		return sum, errors.New("address too long")
	}

	octs := strings.Split(address, ".")

	if len(octs) != 4 {
		return sum, errors.New("requires 4 octects")
	}

	for i := 0; i < 4; i++ {
		oct, err := strconv.ParseUint(octs[i], 10, 0)
		if err != nil {
			return sum, errors.New("bad octect " + octs[i])
		}
		sum += uint32(oct) << uint32((4-1-i)*8)
	}
	return sum, nil
}

// Itoa returns the string representation of an ipv4 address uint32 value.
// Example:
//
//	Itoa(3232312315)  // "192.168.0.1"
//
func Itoa(integer uint32) string {
	var buf bytes.Buffer

	for i := 0; i < 4; i++ {
		oct := (integer >> uint32((4-1-i)*8)) & 0xff
		buf.WriteString(strconv.FormatUint(uint64(oct), 10))
		if i < 3 {
			buf.WriteByte('.')
		}
	}
	return buf.String()
}
