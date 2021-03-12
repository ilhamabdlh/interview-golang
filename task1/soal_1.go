package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("input Number:")
	for scanner.Scan() {
		fmt.Println(MissingNumber(scanner.Text()))
		//fmt.Println(pecahAngka(scanner.Text(), 2))
	}
}

func MissingNumber(input string) (result int) {
	pangkatAwal := 1
	pangkatAkhir := 6
	splitYangBenar := []string{}
	angkaYangBenar := 0
	// cari tau ini masuk ke pangkat yang mana
	for i := pangkatAwal; i <= pangkatAkhir; i++ {
		arraySplitted := pecahAngka(input, i)

		angkaPertama, _ := strconv.Atoi(arraySplitted[0])
		angkaKedua, _ := strconv.Atoi(arraySplitted[1])
		angkaSebelumTerakhir, _ := strconv.Atoi(arraySplitted[len(arraySplitted)-2])
		angkaTerakhir, _ := strconv.Atoi(arraySplitted[len(arraySplitted)-1])

		if angkaPertama+1 == angkaKedua || angkaSebelumTerakhir+1 == angkaTerakhir {
			splitYangBenar = arraySplitted
		}
	}

	for index := 0; index < len(splitYangBenar)-1; index++ {
		angkaSebelumnya, _ := strconv.Atoi(splitYangBenar[index])
		angkaSetelahnya, _ := strconv.Atoi(splitYangBenar[index+1])

		if angkaSetelahnya-angkaSebelumnya > 1 {
			angkaYangBenar = angkaSebelumnya + 1
		}
	}
	return angkaYangBenar
}

func pecahAngka(input string, perBerapaAngka int) []string {
	result := []string{}
	runes := []rune(input)

	for i := 0; i < len(input); i = i + perBerapaAngka {
		susbtring := string(runes[i : i+perBerapaAngka])
		result = append(result, susbtring)
	}
	return result
}
