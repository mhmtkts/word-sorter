package tests

import (
    "testing"
    "word-sorter/src/sorter"
)

func TestSortWords(t *testing.T) {
    input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
    expected := []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"}

    t.Log("Initial input:", input)
    
    result := sorter.SortByACount(input)
    
    t.Log("Expected output:", expected)
    t.Log("Received output:", result)

    if !areSlicesEqual(result, expected) {
        t.Errorf("Test failed!\nExpected: %v\nReceived: %v", expected, result)
    } else {
        t.Log("Test passed successfully!")
    }
}

func areSlicesEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}