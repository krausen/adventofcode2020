package main

import (
  "fmt"
  "testing"
)

var input = []string{
  "..##.......",
  "#...#...#..",
  ".#....#..#.",
  "..#.#...#.#",
  ".#...##..#.",
  "..#.##.....",
  ".#.#.#....#",
  ".#........#",
  "#.##...#...",
  "#...##....#",
  ".#..#...#.#",
}

func TestFiveOne(t *testing.T) {
  fiveOne := countTrees(5, 1, input)
  assertEquals(t, 3, fiveOne)

}

func TestOneOne(t *testing.T) {
  oneOne := countTrees(1, 1, input)
  assertEquals(t, 2, oneOne)
}

func TestThreeOne(t *testing.T) {
  threeOne := countTrees(3, 1, input)
  assertEquals(t, 7, threeOne)
}

func TestSevenOne(t *testing.T) {
  sevenOne := countTrees(7, 1, input)
  assertEquals(t, 4, sevenOne)
}

func TestOneTwo(t *testing.T) {
  oneOne := countTrees(1, 2, input)
  assertEquals(t, 2, oneOne)
}

func assertEquals(t *testing.T, expected int, actual int) {
  if actual != expected {
    var message string = fmt.Sprintf("Expected %d but go %d", expected, actual)
    t.Error(message)
  }
}
