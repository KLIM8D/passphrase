package password_generator

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"math/big"
	"os"
	"strings"
)

type Symbols uint8

const (
	ALL Symbols = iota
	MINIMAL
	NONE
)

type PasswordGenerator struct {
	Words   []string
	Symbols []string
}

func New(symbolsToUse Symbols, filePath string) *PasswordGenerator {
	words, err := readLines(filePath)
	if err != nil {
		panic(err)
	}

	var symbols []string
	switch symbolsToUse {
	case MINIMAL:
		symbols = []string{"!", "@", "$", "%", "^", "&", "*", "0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9"}

	case ALL:
		symbols = []string{"!", "@", "$", "%", "^", "&", "*", "0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9", "`", "~", "#", "-", "=", "+", "[", "{", "]", "}"}
	}

	return &PasswordGenerator{words, symbols}
}

func (pg *PasswordGenerator) GeneratePassphrase(wordCount int) string {
	var password bytes.Buffer
	var nWords *big.Int
	var nSymbols *big.Int

	nWords = big.NewInt(int64(len(pg.Words)))
	if pg.Symbols != nil {
		nSymbols = big.NewInt(int64(len(pg.Symbols)))
	}

	for i := 0; i < wordCount; i++ {
		n, err := rand.Int(rand.Reader, nWords)
		if err != nil {
			panic(err)
		}

		word := strings.Title(pg.Words[n.Int64()])
		if pg.Symbols != nil {
			n0, err := rand.Int(rand.Reader, nSymbols)
			if err != nil {
				panic(err)
			}
			word += pg.Symbols[n0.Int64()]
		}

		password.WriteString(word)
	}

	return password.String()
}

func readLines(fileToRead string) ([]string, error) {
	file, err := os.Open(fileToRead)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
