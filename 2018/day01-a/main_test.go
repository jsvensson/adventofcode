package main

import "testing"

func Test_ParseValue_Positive(t *testing.T) {
	if v := parseValue("+10"); v != 10 {
		t.Fail()
	}
}

func Test_ParseValue_Negative(t *testing.T) {
	if v := parseValue("-10"); v != -10 {
		t.Fail()
	}
}
