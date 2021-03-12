package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//scan number in terminal
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("input Number:")
	for scanner.Scan() {
		input := scanner.Text()
		splitted := strings.Split(input, " ")
		input1, _ := strconv.Atoi(splitted[0])
		input2, _ := strconv.Atoi(splitted[1])
		fmt.Println(PalindromeMaster(input1, input2))
	}
}

var palFull int

//handle input with 2 param
func PalindromeMaster(num1, num2 int) (result int) {
	palindrome := num1 < num2
	//if inpu1 !< input2, continue
	if palindrome == true {
		for i := num1; i < num2; i++ {
			//set value
			palOne := 0
			palTwo := i
			palTree := palTwo

			//set compar
			for palTree > 0 {
				palFour := palTree % 10
				palOne = (palOne * 10) + palFour
				palTree = palTree / 10
			}
			if palTwo == palOne {
				palFull++
			}
		}
		result = palFull
		//reset value
		palFull = 0
	}
	return result
}
