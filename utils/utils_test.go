package utils

import (
	"testing"
)

func TestNthElement(t *testing.T) {
	longList := []string{"foo", "bar", "bla", "blah"}
	assert(t, NthElement(longList, 0), "foo")
	assert(t, NthElement(longList, 1), "bar")
	assert(t, NthElement(longList, 2), "bla")
	assert(t, NthElement(longList, 5), "")
	assert(t, NthElement(longList, -10), "")
	assert(t, NthElement([]string{}, 0), "")
}

func assert(t *testing.T, var1 any, var2 any) {
	if var1 != var2 {
		t.Errorf("assert failed: %s != %s\n", var1, var2)
	}
}
