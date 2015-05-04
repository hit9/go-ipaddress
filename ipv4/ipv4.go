// IPv4 addr utils for golang.
// Copyright 2015. Chao Wang <hit9@icloud.com>

package ipv4

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

// Atoi returns the uint32 representation of an ipv4 addr string value.
// Example:
//
//	Atoi("192.168.0.1")   // 3232312315
//
func Atoi(addr string) (uint32, error) {
	var sum uint32 = 0
	if len(addr) > 15 {
		return sum, errors.New("addr too long")
	}
	octs := strings.Split(addr, ".")
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

// Itoa returns the string representation of an ipv4 addr uint32 value.
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

// Example:
//
//	Not("0.0.255.255")  // "255.255.0.0"
//
func Not(addr string) (string, error) {
	i, err := Atoi(addr)
	return Itoa(i ^ 0xffffffff), err
}

// Example:
//
//	Or("0.0.1.1", "1.1.0.0")  // "1.1.1.1"
//
func Or(addra string, addrb string) (string, error) {
	ia, err := Atoi(addra)
	ib, err := Atoi(addrb)
	return Itoa(ia | ib), err
}

// Example:
//
//	Xor("0.255.255.255", "192.255.255.255")  // "192.0.0.0"
//
func Xor(addra string, addrb string) (string, error) {
	ia, err := Atoi(addra)
	ib, err := Atoi(addrb)
	return Itoa(ia ^ ib), err
}

// Example:
//
//	Next("192.168.0.1")  // "192.168.0.2"
//
func Next(addr string) (string, error) {
	i, err := Atoi(addr)
	return Itoa(i + 1), err
}

// Example:
//
//	Prev("192.168.0.1")  // "192.168.0.0"
//
func Prev(addr string) (string, error) {
	i, err := Atoi(addr)
	return Itoa(i - 1), err
}
