// Copyright 2015. Chao Wang <hit9@icloud.com>

// Ipv4 address utils for golang.
package ipv4

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type Net struct {
	address   string // address
	bitmask   uint8  // bitmask
	mask      string // mask
	hostmask  string // hostmask
	broadcast string // broadcast
	first     string // first
	last      string // last
	size      uint32 //size
}

// Atoi returns the uint32 representation of an ipv4 addr string value.
//
// Example:
//
//	Atoi("192.168.0.1")   // 3232312315
//
func Atoi(addr string) (sum uint32, err error) {
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
//
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
func Or(addra string, addrb string) (addr string, err error) {
	ia, err := Atoi(addra)
	if err != nil {
		return addr, err
	}

	ib, err := Atoi(addrb)
	if err != nil {
		return addr, err
	}

	return Itoa(ia | ib), err
}

// Example:
//
//	Xor("0.255.255.255", "192.255.255.255")  // "192.0.0.0"
//
func Xor(addra string, addrb string) (addr string, err error) {
	ia, err := Atoi(addra)
	if err != nil {
		return addr, err
	}

	ib, err := Atoi(addrb)
	if err != nil {
		return addr, err
	}

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

// Returns information for a netblock.
func Network(block string) (net Net, err error) {
	if len(block) > 18 {
		return net, errors.New("block too long")
	}

	list := strings.Split(block, "/")
	if len(list) != 2 {
		return net, errors.New("invalid block")
	}

	// address
	net.address = list[0]

	// bitmask
	bitmask, err := strconv.ParseUint(list[1], 10, 0)
	if err != nil {
		return net, err
	}
	if bitmask&31 != bitmask {
		return net, errors.New("invalid bitmask")
	}
	net.bitmask = uint8(bitmask)

	// mask
	net.mask = Itoa(0xffffffff >> (32 - net.bitmask) << (32 - net.bitmask))
	net.hostmask, err = Not(net.mask)
	if err != nil {
		return net, err
	}

	// broadcast
	net.broadcast, err = Or(net.address, net.hostmask)
	if err != nil {
		return net, err
	}

	// first
	addr, err := Xor(net.hostmask, net.broadcast)
	if err != nil {
		return net, err
	}

	net.first, err = Next(addr)
	if err != nil {
		return net, err
	}

	// last
	net.last, err = Prev(net.broadcast)
	if err != nil {
		return net, err
	}

	// size
	i, err := Atoi(net.last)
	if err != nil {
		return net, err
	}

	j, err := Atoi(net.first)
	if err != nil {
		return net, err
	}

	net.size = i - j + 1
	return net, nil
}
