package util

// Permutations implements Heap's algorithm to generate all permutations of an array of strings.
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
//
// Use this function by passing a 'new([][]string)', the input collection and the length of the input collection.
func Permutations(result *[][]string, l []string, k int) {
	if k == 1 {
		c := make([]string, len(l))
		copy(c, l)
		*result = append(*result, c)
		return
	}

	Permutations(result, l, k-1)

	for i := 0; i < k-1; i++ {
		if k%2 == 0 {
			l[i], l[k-1] = l[k-1], l[i]
		} else {
			l[0], l[k-1] = l[k-1], l[0]
		}
		Permutations(result, l, k-1)
	}
}
