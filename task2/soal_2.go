package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//scan input
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("input Number:")
	for scanner.Scan() {
		fmt.Println(MissingNumber(scanner.Text()))
	}
}

//handle input
func MissingNumber(input string) (result int) {
	//create var for first and last number
	firstHandle := 1
	closeHandle := 6
	splitTrue := []string{}
	numTrue := 0
	// find unit numbers, tens, hundreds, thousands, etc.
	for i := firstHandle; i <= closeHandle; i++ {
		arraySplitted := splitNum(input, i)

		firstNumber, _ := strconv.Atoi(arraySplitted[0])
		secondNumber, _ := strconv.Atoi(arraySplitted[1])
		befLastNumber, _ := strconv.Atoi(arraySplitted[len(arraySplitted)-2])
		lastNumber, _ := strconv.Atoi(arraySplitted[len(arraySplitted)-1])

		if firstNumber+1 == secondNumber || befLastNumber+1 == lastNumber {
			splitTrue = arraySplitted
		}
	}

	//looking for the missing number after finding the fraction
	for index := 0; index < len(splitTrue)-1; index++ {
		prevNumber, _ := strconv.Atoi(splitTrue[index])
		afterNumber, _ := strconv.Atoi(splitTrue[index+1])

		if afterNumber-prevNumber > 1 {
			numTrue = prevNumber + 1
		}
	}
	return numTrue
}
func splitNum(input string, perBerapaAngka int) []string {
	result := []string{}
	runes := []rune(input)

	for i := 0; i < len(input); i = i + perBerapaAngka {
		susbtring := string(runes[i : i+perBerapaAngka])
		result = append(result, susbtring)
	}
	return result
}
