package example

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := Add(2, 4)
	if r != 6 {
		t.Fatalf("Add(2, 4) error, expect:%d, actual:%d", 6, r)
	}
	t.Logf("test Add succ")
}