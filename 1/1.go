package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"
)

func main() {
  nums := readIntegerLines("1.input")
  fmt.Println(find_2(nums))
  fmt.Println(find_3(nums))
}

func readIntegerLines(filePath string) []int {
  content, _ := ioutil.ReadFile(filePath)
  asStrings := strings.Split(string(content), "\n")
  integers := make([]int, 0)
  for i := 0; i < len(asStrings); i++ {
    asInteger, _ := strconv.Atoi(asStrings[i])
    integers = append(integers, asInteger)
  }
  return integers
}

func find_2(nums []int) int {
  for i := 0; i < len(nums); i++ {
    a := nums[i]
    for j := i + 1; j < len(nums); j++ {
      b := nums[j]
      if a+b == 2020 {
        return a * b
      }
    }
  }
  return -1
}

func find_3(nums []int) int {
  for i := 0; i < len(nums); i++ {
    a := nums[i]
    for j := i + 1; j < len(nums); j++ {
      b := nums[j]
      for k := j + 1; k < len(nums); k++ {
        c := nums[k]
        if a+b+c == 2020 {
          return a * b * c
        }
      }
    }
  }
  return -1
}
