package helper

// Converts a string to runes
func ToRune(str string) []rune {
	runes := make([]rune, len(str))
	for i, e := range str {
		runes[i] = e
	}
	return runes
}

type Frequencies map[rune]int

// Get a map of character frequencies of a string
func ToFrequency(str string) Frequencies {
	freq := make(map[rune]int)
	for _, r := range str {
		freq[r] += 1
	}
	return freq
}

// Converts the frequency map to a set
func (f Frequencies) ToSet() Set {
	s := Set{}
	for k := range f {
		s.Add(k)
	}
	return s
}

// Converts the frequency map to an array
func (f Frequencies) ToArray() []int {
	freqs := make([]int, 0, len(f))
	for _, v := range f {
		freqs = append(freqs, v)
	}
	return freqs
}
