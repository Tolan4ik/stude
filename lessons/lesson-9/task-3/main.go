package main

import "fmt"

func countWords(words []string, limit int) map[string]int {

	counts := make(map[string]int)
	for _, w := range words {
		_, ok := counts[w]
		if !ok && len(counts) >= limit {

			for k := range counts {
				delete(counts, k)
				break
			}
		}

		counts[w]++
	}
	return counts
}

func main() {

	slice := []string{"go", "map", "go", "code", "Natasha", "Tola", "Olga", "Ana", "Sona", "Sona", "Olga"}
	sli := countWords(slice, 3)
	fmt.Println(sli)
}
