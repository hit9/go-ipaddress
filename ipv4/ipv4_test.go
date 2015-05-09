// Copyright 2015. Chao Wang <hit9@icloud.com>

package ipv4_test

import (
	. "github.com/hit9/go-ipv4/ipv4"
	"testing"
)

func TestAtoi(t *testing.T) {
	var except uint32 = 3232235521
	got, err := Atoi("192.168.0.1")
	if err != nil {
		t.Errorf("unexcepted error %v", err)
	}
	if except != got {
		t.Errorf("except=%v but got=%v", except, got)
	}
}

func TestItoa(t *testing.T) {
	except := "192.168.0.1"
	got := Itoa(3232235521)
	if except != got {
		t.Errorf("except=%v but got=%v", except, got)
	}
}
