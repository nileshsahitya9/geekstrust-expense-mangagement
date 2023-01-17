package input

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func FormattedInput() [][]string {
	cliArgs := os.Args[1:]

	if validation(cliArgs) {
		return make([][]string, 0)
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")
		return make([][]string, 0)
	}

	defer file.Close()
	return processingInput(file)
}

func validation(cliArgs []string) bool {
	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")
		return true
	}
	return false
}

func processingInput(file io.Reader) [][]string {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text [][]string

	for scanner.Scan() {
		text = append(text, strings.Split(strings.ReplaceAll(scanner.Text(), " ", "/"), "/"))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}
