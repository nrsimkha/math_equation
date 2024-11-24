package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя входного файла: ")
	fi_name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fi_name = strings.TrimSpace(fi_name)

	reader = bufio.NewReader(os.Stdin)
	fmt.Print("Введите имя файла для вывода результатов: ")
	fo_name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fo_name = strings.TrimSpace(fo_name)

	// open output file
	fo, err := os.Create(fo_name)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	writer := bufio.NewWriter(fo)
	contentBytes, err := os.ReadFile(fi_name)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile(`([0-9]+)([\+\-\*\/]{1})([0-9]+)=\?`)

	submatches := re.FindAllStringSubmatch(string(contentBytes), -1)

	for _, s := range submatches {
		equationResult := calculateEquation(s[1], s[3], s[2])
		writer.Write([]byte(s[1] + s[2] + s[3] + "=" + equationResult + "\n"))
	}
	writer.Flush()
}

func calculateEquation(firstValue string, secondValue string, operator string) string {
	firstValue_int, err := strconv.Atoi(firstValue)
	if err != nil {
		panic(err)
	}
	secondValue_int, err := strconv.Atoi(secondValue)
	if err != nil {
		panic(err)
	}
	switch operator {
	case "+":
		return strconv.Itoa(firstValue_int + secondValue_int)
	case "-":
		return strconv.Itoa(firstValue_int - secondValue_int)
	case "*":
		return strconv.Itoa(firstValue_int * secondValue_int)
	case "/":
		return strconv.Itoa(firstValue_int / secondValue_int)
	default:
		panic("Unknown operator string")
	}
}
