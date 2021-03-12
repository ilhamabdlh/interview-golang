package main

import (
	"fmt"
	"testing"
)

var output string = "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3N14 3A13"

func TestBooksLibrary(t *testing.T) {
	checkLibrary := BooksLibrary("0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3A13")
	if checkLibrary == output {
		fmt.Println("ordinal :", checkLibrary)
	} else {
		t.Logf("find %d expect %d", checkLibrary, 9)
	}
}
