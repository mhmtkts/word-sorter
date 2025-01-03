package main

import (
    "fmt"
    "word-sorter/src/sorter"
)

func main() {
    words := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
    sortedWords := sorter.SortByACount(words)
    fmt.Println("Input:", words)
    fmt.Println("Output:", sortedWords)
}