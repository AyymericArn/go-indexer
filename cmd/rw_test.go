package cmd

import (
	"fmt"
	"search/engine"
	"testing"
)

func TestReadWrite(t *testing.T) {
	c, err := engine.Dial()
	if err != nil {
		fmt.Println("failed at dial")
		t.Fail()
	}

	_, err = c.Do("FLUSHALL")
	if err != nil {
		fmt.Println("failed at flush")
		t.Fail()
	}

	err = engine.AddFile(c, "grenouille", "/test/grenouille", 1)
	if err != nil {
		fmt.Println("failed at add")
		t.Fail()
	}

	res, err := engine.Get(c, "grenouille")
	if err != nil || len(res) != 1 || res[0] != "/test/grenouille" {
		fmt.Println(err)
		t.Fail()
	}
}
