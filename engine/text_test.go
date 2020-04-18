package engine

import (
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	s1, s2, s3 := strings.NewReader(""), strings.NewReader("grenouille"), strings.NewReader("grenouille grenouille")
	r1, r2, r3 := Scan(s1), Scan(s2), Scan(s3)

	if len(r1) != 0 {
		t.Fail()
	}

	if r2["grenouille"] != 1 {
		t.Fail()
	}

	if r3["grenouille"] != 2 {
		t.Fail()
	}
}
