package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode/utf8"
)

type flagFunctionPair struct {
	flag     *bool
	function func(*os.File)
}

func main() {
	countBytesFlag := flag.Bool("c", false, "Count bytes in the file")
	countLinesFlag := flag.Bool("l", false, "Count lines in the file")
	countWordsFlag := flag.Bool("w", false, "Count words in the file")
	countCharsFlag := flag.Bool("m", false, "Count characters in the file")

	flag.Parse()

	var file *os.File
	var err error
	fileName := flag.Arg(0)

	if fileName == "" {
		file = os.Stdin

	} else {
		file, err = os.Open(fileName)
		if err != nil {
			panic("Error opening file")
		}
		defer file.Close()
	}

	flagFunctionPairs := []flagFunctionPair{
		{countBytesFlag, countBytes},
		{countLinesFlag, countLines},
		{countWordsFlag, countWords},
		{countCharsFlag, countChars},
	}

	anyFlagsSet := false

	for _, pair := range flagFunctionPairs {
		if *pair.flag {
			pair.function(file)
		}
	}

	if !anyFlagsSet {
		for _, pair := range flagFunctionPairs {
			pair.function(file)
		}
	}
}

func countBytes(file *os.File) {
	fileName := file.Name()
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic("Error reading file")
	}
	byteCount := len(data)
	fmt.Printf("%d %s\n", byteCount, fileName)
}

func countLines(file *os.File) {
	lineCount := 0
	s := bufio.NewScanner(file)

	for s.Scan() {
		lineCount++
	}
	file.Seek(0, 0)
	fmt.Printf("%d %s\n", lineCount, file.Name())
}

func countWords(file *os.File) {
	wordCount := 0
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanWords)

	for s.Scan() {
		wordCount++
	}
	file.Seek(0, 0)
	fmt.Printf("%d %s\n", wordCount, file.Name())
}

func countChars(file *os.File) {
	fileName := file.Name()
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic("Error reading file")
	}
	fmt.Printf("%d %s\n", utf8.RuneCount(data), fileName)
}
