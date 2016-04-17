package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func readLines(fileToRead string) ([]string, error) {
	file, err := os.Open(fileToRead)
	if err != nil {
		fmt.Println("retard")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	var password bytes.Buffer
	rand.Seed(time.Now().UTC().UnixNano())

	args := os.Args[1:]
	fmt.Println(args[0], args[1])

	words, err := readLines(args[0])
	if err != nil {
		fmt.Println("Something went wrong in words")
	}

	wordCount, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Something went wrong in wordCount")
	}

	for i := 0; i < wordCount; i++ {
		n := rand.Intn(len(words))
		password.WriteString(words[n])
	}
	fmt.Println(password.String())
}
