func findSubstring(s string, words []string) []int {
	wordLength := 0
	wordListLength := len(words)
	fuckingStringLength := len(s)
	numberOfEachWord := make(map[string]int)
	startingIndices := make([]int, 0)
	matchedRange := make([]int, 2)
	prevStateWithoutFirstWord := make(map[string]int)

	if wordListLength > 0 {
		wordLength = len(words[0])
	} else {
		return startingIndices
	}

	for _, word := range words {
		numberOfEachWord[word] += 1
	}

	for i := 0; i <= fuckingStringLength-wordLength; i++ {
		for _, word := range words {
			startingIndice := i

			if s[i:i+wordLength] == word {
				wordCounter := 0
				numberOfEachWordClone := make(map[string]int)

				if i < matchedRange[1] {
					copyMap(prevStateWithoutFirstWord, numberOfEachWordClone)
				} else {
					copyMap(numberOfEachWord, numberOfEachWordClone)
					numberOfEachWordClone[word]--
				}

				// Check left part
				if i >= matchedRange[1] {
					c1 := i

					for c1 >= wordLength && c1 >= i-wordLength*(fuckingStringLength-1) {
						checkedWord := s[c1-wordLength : c1]

						if numberOfEachWordClone[checkedWord] >= 1 {
							wordCounter++
							numberOfEachWordClone[checkedWord]--
							startingIndice = c1 - wordLength
							c1 -= wordLength
						} else {
							break
						}
					}
				}

				// Skip if index exists
				if icontains(startingIndices, startingIndice) != -1 {
					continue
				}

				// Check right part
				c2 := i + wordLength

				if i < matchedRange[1] {
					variance := matchedRange[1] - i
					c2 = (variance + 1) * wordLength
				}

				remainingWordListLenght := fuckingStringLength - wordCounter

				for c2 <= len(s)-wordLength && c2 <= i+wordLength*remainingWordListLenght {
					checkedWord := s[c2 : c2+wordLength]

					if numberOfEachWordClone[checkedWord] >= 1 {
						wordCounter++
						numberOfEachWordClone[checkedWord]--
						c2 += wordLength
					} else {
						break
					}
				}

				if wordCounter == wordListLength-1 {
					matchedRange[0] = startingIndice
					matchedRange[1] = startingIndice + wordCounter
					prevStateWithoutFirstWord = numberOfEachWordClone
					prevStateWithoutFirstWord[s[startingIndice:startingIndice+wordLength]] = 1

					startingIndices = append(startingIndices, startingIndice)
				}

				break
			}
		}
	}

	return startingIndices
}

func icontains(s []int, str int) int {
	for i, v := range s {
		if v == str {
			return i
		}
	}

	return -1
}

func copyMap(originalMap map[string]int, newMap map[string]int) {
	for k, v := range originalMap {
		newMap[k] = v
	}
}