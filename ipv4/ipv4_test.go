// Copyright 2015. Chao Wang <hit9@icloud.com>

package ipv4_test

import (
	. "github.com/hit9/go-ipv4/ipv4"
	"testing"
)

var AtoiCases = []struct {
	addr   string
	except uint32
}{
	{"0.0.0.0", 0},
	{"0.0.1.0", 256},
	{"0.1.1.0", 256 + 256*256},
	{"1.1.1.0", 256 + 256*256 + 256*256*256},
	{"192.168.0.1", 3232235521},
}

func TestAtoi(t *testing.T) {
	for _, testCase := range AtoiCases {
		got, err := Atoi(testCase.addr)
		if err != nil {
			t.Errorf("unexcepted error %v", err)
		}
		if got != testCase.except {
			t.Errorf("except=%v but got=%v", testCase.except, got)
		}
	}
}

var ItoaCases = []struct {
	addr   uint32
	except string
}{
	{0, "0.0.0.0"},
	{256, "0.0.1.0"},
	{256 + 256*256, "0.1.1.0"},
	{256 + 256*256 + 256*256*256, "1.1.1.0"},
	{3232235521, "192.168.0.1"},
}

func TestItoa(t *testing.T) {
	for _, testCase := range ItoaCases {
		got := Itoa(testCase.addr)
		if got != testCase.except {
			t.Errorf("except=%v but got=%v", testCase.except, got)
		}
	}
}

var NotCases = []struct {
	addr   string
	except string
}{
	{"0.0.1.1", "255.255.254.254"},
	{"0.0.255.255", "255.255.0.0"},
}

func TestNot(t *testing.T) {
	for _, testCase := range NotCases {
		got, err := Not(testCase.addr)
		if err != nil {
			t.Errorf("unexcepted error %v", err)
		}
		if got != testCase.except {
			t.Errorf("except=%v but got=%v", testCase.except, got)
		}
	}
}

var OrCases = []struct {
	addra  string
	addrb  string
	except string
}{
	{"0.0.1.1", "1.1.0.0", "1.1.1.1"},
	{"0.0.1.2", "1.2.0.0", "1.2.1.2"},
	{"0.0.1.233", "1.2.0.0", "1.2.1.233"},
	{"0.0.1.233", "1.2.0.2", "1.2.1.235"},
}

func TestOr(t *testing.T) {
	for _, testCase := range OrCases {
		got, err := Or(testCase.addra, testCase.addrb)
		if err != nil {
			t.Errorf("unexcepted error %v", err)
		}
		if got != testCase.except {
			t.Errorf("except=%v but got=%v", testCase.except, got)
		}
	}
}

var XorCases = []struct {
	addra  string
	addrb  string
	except string
}{
	{"0.255.255.255", "192.255.255.255", "192.0.0.0"},
}

func TestXor(t *testing.T) {
	for _, testCase := range XorCases {
		got, err := Xor(testCase.addra, testCase.addrb)
		if err != nil {
			t.Errorf("unexcepted error %v", err)
		}
		if got != testCase.except {
			t.Errorf("except=%v but got=%v", testCase.except, got)
		}
	}
}

var PrevCases = []struct {
	addr   string
	except string
}{
	{"0.0.0.1", "0.0.0.0"},
	{"0.0.255.255", "0.0.255.254"},
	{"0.0.0.0", "255.255.255.255"},
	{"192.168.0.1", "192.168.0.0"},
}

func TestPrev(t *testing.T) {
	for _, testCase := range PrevCases {
		got, err := Prev(testCase.addr)
		if err != nil {
			t.Errorf("unexcepted error %v", err)
		}
		if got != testCase.except {
			t.Errorf("except=%v but got=%v", testCase.except, got)
		}
	}
}

var NextCases = []struct {
	addr   string
	except string
}{
	{"0.0.0.0", "0.0.0.1"},
	{"0.0.255.255", "0.1.0.0"},
	{"255.255.255.255", "0.0.0.0"},
	{"192.168.0.1", "192.168.0.2"},
}

func TestNext(t *testing.T) {
	for _, testCase := range NextCases {
		got, err := Next(testCase.addr)
		if err != nil {
			t.Errorf("unexcepted error %v", err)
		}
		if got != testCase.except {
			t.Errorf("except=%v but got=%v", testCase.except, got)
		}
	}
}

var NetworkCases = []struct {
	block  string
	except Net
}{
	{"192.168.0.0/24", Net{
		"192.168.0.0",
		24,
		"255.255.255.0",
		"0.0.0.255",
		"192.168.0.255",
		"192.168.0.1",
		"192.168.0.254",
		254}},
}

func TestNetwork(t *testing.T) {
	for _, testCase := range NetworkCases {
		got, err := Network(testCase.block)
		if err != nil {
			t.Errorf("unexcepted error %v", err)
		}
		if got != testCase.except {
			t.Errorf("except=%v but got=%v", testCase.except, got)
		}
	}
}
