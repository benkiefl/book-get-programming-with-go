package main

import (
	"fmt"
	"strings"
)

func count(text string) map[string]int {
	var words []string = strings.Fields(strings.ToLower(text))
	frequency := make(map[string]int)
	for _, word := range words {
		word = strings.Trim(word, ` ,./;'"!?`)
		frequency[word]++
	}
	return frequency
}

func main() {

	var quote string = `Of all tyrannies, a tyranny sincerely exercised for the good of its victims may be the most oppressive. It would be better to live under robber barons than under omnipotent moral busybodies. The robber baron's cruelty may sometimes sleep, his cupidity may at some point be satiated; but those who torment us for our own good will torment us without end for they do so with the approval of their own conscience. They may be more likely to go to Heaven yet at the same time likelier to make a Hell of earth. This very kindness stings with intolerable insult. To be "cured" against one's will and cured of states which we may not regard as disease is to be put on a level of those who have not yet reached the age of reason or those who never will; to be classed with infants, imbeciles, and domestic animals.`

	map_words := count(quote)
	map_words_sorted := make(map[int][]string)

	for k, v := range map_words {
		map_words_sorted[v] = append(map_words_sorted[v], k)
	}

	fmt.Println(map_words_sorted)
}
