package main

import (
  "fmt"
  "io/ioutil"
  "strings"
)

func main() {
  lines := readLines("3.input")

  oneOne := countTrees(1, 1, lines)
  threeOne := countTrees(3, 1, lines)
  fiveOne := countTrees(5, 1, lines)
  sevenOne := countTrees(7, 1, lines)
  oneTwo := countTrees(1, 2, lines)

  fmt.Println("Right 1, down 1: ", oneOne)
  fmt.Println("Right 3, down 1: ", threeOne)
  fmt.Println("Right 5, down 1: ", fiveOne)
  fmt.Println("Right 7, down 1: ", sevenOne)
  fmt.Println("Right 1, down 2: ", oneTwo)
  fmt.Println(oneOne * threeOne * fiveOne * sevenOne * oneTwo)
}

func readLines(filePath string) []string {
  content, _ := ioutil.ReadFile(filePath)
  return strings.Split(string(content), "\n")
}

func countTrees(right int, down int, slope []string) int {
  numberOfTrees := 0
  breadthOfSlope := len(slope[0])
  for i := 0; i*down < len(slope); i++ {
    cellValue := string(slope[(i * down)][(i*right)%breadthOfSlope])
    if cellValue == "#" {
      numberOfTrees += 1
    }
  }

  return numberOfTrees
}
