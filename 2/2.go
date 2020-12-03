package main

import ("fmt"
		"strings"
		"strconv"
		"io/ioutil")

func main() {
	passwords := readLines("2.input")
	validPasswords1 := 0
	validPasswords2 := 0
	for i := 0; i < len(passwords); i++ {
		if validatePasswordOccurences(splitPasswordLine(passwords[i])) {
			validPasswords1++
		}
		if validatePasswordExclusiveOne(splitPasswordLine(passwords[i])) {
			validPasswords2++
		}
	}
	fmt.Println(validPasswords1)
	fmt.Println(validPasswords2)
}

func readLines(filePath string) []string {
	content, _ := ioutil.ReadFile(filePath)
	return strings.Split(string(content), "\n")
}

func validatePasswordOccurences(min int, max int, char string, password string) bool {
	occurences := 0
	for i := 0; i < len(password); i++ {
		if char == string(password[i]) {
			occurences += 1
		}
	}
	return (min <= occurences) && (occurences <= max)
}

func validatePasswordExclusiveOne(min int, max int, char string, password string) bool {
	existsFirst := char == string(password[min-1])
	existsSecond := char == string(password[max-1])
	return (existsFirst || existsSecond) && !(existsFirst && existsSecond)
}

func splitPasswordLine(line string) (int, int, string, string) {
	splitLine := strings.Split(line, ":")
	rest := splitLine[0]
	password := strings.TrimSpace(splitLine[1])
	splitLine = strings.Split(rest, " ")
	rest = splitLine[0]
	char := splitLine[1]
	splitLine = strings.Split(rest, "-")
	min, _ := strconv.Atoi(splitLine[0])
	max, _ := strconv.Atoi(splitLine[1])
	return min, max, char, password
}