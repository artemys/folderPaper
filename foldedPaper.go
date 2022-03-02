package foldedPaper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func foldedPaper() error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var N int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &N)
	paper := []string{}
	leftFolded := []string{}

	for i := 0; i < N; i++ {
		scanner.Scan()
		L := scanner.Text()
		_ = L // to avoid unused error
		paper = append(paper, L)
	}
	for _, line := range paper {
		leftFolded = append(leftFolded, foldLeft(line))
		println(N, leftFolded)
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")
	// Write answer to stdout
	return nil

}

func foldToTop(size int, tab []string) []string {
	finaleTab := make([]string, size)
	for i, j := 0, size-1; i < size/2; i, j = i+1, j-1 {
		finaleTab[i] = pairLine(tab[i], swapEachNChar(2, tab[j]), 2)

	}
	return finaleTab
}

func pairLine(str1 string, str2 string, pairnb int) string {
	runes2 := []rune(str2)
	finale := []rune(str1 + str2)
	len := len(finale)
	var tmporigin *[]rune
	fmt.Println(len)

	tmporigin = &runes2

	for i, j, k := 0, 0, 0; i < len; i, j, k = i+1, j+1, k+1 {

		if j == pairnb {
			if strings.Compare(string(*tmporigin), str2) == 0 {

				*tmporigin = []rune(fmt.Sprintf("%s", str1))
				k = k - 2
				fmt.Println(k)

			} else if strings.Compare(string(*tmporigin), str1) == 0 {

				*tmporigin = []rune(fmt.Sprintf("%s", str2))
			}
			j = 0
		}

		tmp2 := *tmporigin
		finale[i] = tmp2[k]
	}
	return string(finale)
}

func swapLeftNChar(n int, str1 string, str2 string) string {
	runes := []rune(str1 + str2)

	len := len(runes)

	for i := 0; i < len; i++ {
		tmp := runes[i]
		runes[i] = runes[i+n]
		runes[i+1] = tmp
	}
	println(string(runes))
	return string(runes)
}

func swapEachNChar(n int, str string) string {
	runes := []rune(str)
	len := len(str)

	for i := 0; i < len; i = i + n {
		tmp := runes[i]
		runes[i] = runes[i+1]
		runes[i+1] = tmp
	}
	println(string(runes))
	return string(runes)
}

func foldLeft(str string) string {
	runes := []rune(str)
	runesFinal := []rune(str)
	len := len(str) - 1
	pair := 0

	for i := len; pair < len; i-- {
		//	fmt.Printf("i: %d\nrunes[len-i]:%s\nrunes[i]:%s\nrunesFinal[len-i+1]:%s\n", i, string(runes[len-i]), string(runes[i]), string(runesFinal[len-i+1]))

		tmp := runes[len-i]
		runesFinal[pair] = runes[i]

		runesFinal[pair+1] = tmp
		//	fmt.Printf("string(runesFinal):%s\n", string(runesFinal))

		pair = pair + 2
	}
	return string(runesFinal)
}
