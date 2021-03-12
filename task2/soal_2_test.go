package main

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	testNumber := MissingNumber("1213141517")
	if testNumber == 16 {
		fmt.Println("Missing Number :", testNumber)
	} else {
		t.Logf("find %d expect %d", testNumber, 9)
	}
}
