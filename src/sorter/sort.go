package sorter

// SortByACount sorts words by 'a' count and length
func SortByACount(words []string) []string {
    wordInfos := make([]WordInfo, len(words))
    for i, word := range words {
        aCount := 0
        for _, char := range word {
            if char == 'a' {
                aCount++
            }
        }
        wordInfos[i] = WordInfo{
            Word:   word,
            ACount: aCount,
            Length: len(word),
        }
    }

    for i := 0; i < len(wordInfos); i++ {
        for j := i + 1; j < len(wordInfos); j++ {
            if shouldSwap(wordInfos[i], wordInfos[j]) {
                wordInfos[i], wordInfos[j] = wordInfos[j], wordInfos[i]
            }
        }
    }

    result := make([]string, len(words))
    for i, info := range wordInfos {
        result[i] = info.Word
    }
    return result
}

type WordInfo struct {
    Word   string
    ACount int
    Length int
}

func shouldSwap(a, b WordInfo) bool {
    // Descending sort by the number 'a' first
    if a.ACount != b.ACount {
        return a.ACount < b.ACount
    }
    
    // Descending order by length for the same number 'a'
    if a.Length != b.Length {
        return a.Length < b.Length
    }
    
    // Alphabetical sorting of the same length
    return a.Word > b.Word
}