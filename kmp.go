package helper

// Using KMP algorithm to find the largest prefix that is also a suffix
// of a string and returns an array that keeps the length of a substring like that
// for each substring at [0:i] of the string
func LongestPrefixSuffix(arr string) []int {
	lps := make([]int, 0, len(arr))
	prevLps := 0
	i := 1
	for i < len(arr) {
		if arr[i] == arr[prevLps] {
			lps[i] = prevLps + 1
			i += 1
			prevLps += 1
			continue
		}
		if prevLps == 0 {
			lps[i] = 0
			i += 1
			continue
		}
		prevLps = lps[prevLps-1]
	}
	return lps
}
