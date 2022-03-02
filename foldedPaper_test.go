package foldedPaper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestFolded(t *testing.T) {
	content := []byte("8\nkrloo vi\nnmw Iaok\nie. On o\neeT!ghp \n s hoeua\nkacniob\n. t. ns.\nsoskril")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile
	if err := foldedPaper(); err != nil {
		t.Errorf("userInput failed: %v", err)
	}
	test := foldedPaper()
	log.Println(test)
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func TestSwapEachNChar(t *testing.T) {
	str := "21436587"
	n := 2

	if swapEachNChar(n, str) != "12345678" {
		t.Error("faild")
	}
}

func TestPairLine(t *testing.T) {
	str1 := "abcd"
	str2 := "efgh"
	finale := "efabghcd"
	nb := 2

	if pairLine(str1, str2, nb) != finale {
		t.Error("failed")
	}
}

func TestFoldTop(t *testing.T) {
	tab := []string{
		"ikvr loo",
		"knomawI ",
		"oi en.O ",
		" epehTg!",
		"a use oh",
		" kbaocin",
		"..s nt .",
		"l isroks",
	}
	test := foldToTop(8, tab)

	finale := []string{
		" liksivror lskoo",
		"..kn somtnaw. I ",
		"k oiab econ.niO ",
		" a esupe ehThog!"}

	for pos, line := range finale {
		if line != test[pos] {
			t.Error("fail: ", test[pos], " should be ", line)
			break
		}

	}

}

func TestFoldLeft(t *testing.T) {
	res := foldLeft("krloo vi")
	fmt.Println(res)
}
