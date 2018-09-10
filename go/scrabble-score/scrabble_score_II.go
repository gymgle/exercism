package scrabble

// Score compute the scrabble score for the given word
func Score(word string) int {
	total := 0
	for _, letter := range word {
		switch letter {
		case 'A', 'a', 'E', 'e', 'I', 'i', 'O', 'o', 'U', 'u', 'L', 'l', 'N', 'n', 'R', 'r', 'S', 's', 'T', 't':
			total++
		case 'D', 'd', 'G', 'g':
			total += 2
		case 'B', 'b', 'C', 'c', 'M', 'm', 'P', 'p':
			total += 3
		case 'F', 'f', 'H', 'h', 'V', 'v', 'W', 'w', 'Y', 'y':
			total += 4
		case 'K', 'k':
			total += 5
		case 'J', 'j', 'X', 'x':
			total += 8
		case 'Q', 'q', 'Z', 'z':
			total += 10
		}
	}

	return total
}
