package main

import (
	"fmt"
	"testing"
)

func TestPalindromeMaster(t *testing.T) {
	testPalindrome := PalindromeMaster(1, 10)
	if testPalindrome == 9 {
		fmt.Println("palindrome :", testPalindrome)
	} else {
		t.Logf("find %d expect %d", testPalindrome, 9)
	}
}
