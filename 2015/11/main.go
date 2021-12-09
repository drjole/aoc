package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputBytes, err := os.ReadFile("2015/11/input.txt")
	if err != nil {
		panic(err)
	}
	password := string(inputBytes)

	p := next(password)
	fmt.Println(p)
	fmt.Println(next(p))
}

func next(password string) string {
	for {
		// Incrementing is just like counting with numbers: xx, xy, xz, ya, yb, and so on.
		// Increase the rightmost letter one step; if it was z, it wraps around to a, and repeat with the next letter to
		// the left until one doesn't wrap around.
		passwordBytes := []byte(password)
		i := len(passwordBytes) - 1
		for {
			passwordBytes[i] = (((passwordBytes[i] + 1) - 'a') % 26) + 'a'
			if passwordBytes[i] != 'a' {
				break
			}
			i--
		}
		password = string(passwordBytes)

		// Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and
		// are therefore confusing.
		if strings.ContainsAny(password, "iol") {
			continue
		}

		straight := 1
		pairs := 0
		lastPair := 0
		gotPairs := make(map[byte]struct{})
		for i := 0; i < len(password); i++ {
			// Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on,
			// up to xyz. They cannot skip letters; abd doesn't count.
			if i > 0 && straight < 3 {
				if password[i] == password[i-1]+1 {
					straight++
				} else {
					straight = 1
				}
			}

			// Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
			if i > 0 && i > lastPair+1 && password[i] == password[i-1] {
				if _, ok := gotPairs[password[i]]; !ok {
					lastPair = i
					pairs++
					gotPairs[password[i]] = struct{}{}
				}
			}
		}

		if straight == 3 && pairs >= 2 {
			break
		}
	}

	return password
}
