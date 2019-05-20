// Given a dictionary i.e., a set of strings, and a name, design an efficient
// algorithm that checks whether the name is the concatenation of a sequence
// of dictionary words. If such a concatenation exists, return it. A dictionary
// word may appear more than once in the sequence, e.g., "a", "man", "a",
// "plan", "a", "canal" is a valid sequence for "amanaplanacanal".

package main

import "fmt"

func FindWords(name string, dict map[string]struct{}) (bool, []string) {
	if len(name) == 0 {
		return true, nil
	}

	if len(dict) == 0 {
		return false, nil
	}

	memo := make(map[int]bool)
	words := make([]string, 0)

	var find func(int) bool
	find = func(index int) bool {
		// All done.
		if index == len(name) {
			return true
		}

		// If the check is already performed for this index.
		if valid, ok := memo[index]; ok {
			return valid
		}

		// For each valid range, run a DFS.
		for i := index; i < len(name); i++ {
			if _, ok := dict[name[index:i+1]]; ok {
				if find(i+1) {
					memo[index] = true
					words = append(words, name[index:i+1])
					return true
				}
			}
		}

		return false
	}

	return find(0), words
}

func BuildDictionary(dict map[string]struct{}, str ...string) {
	for _, s := range str {
		dict[s] = struct{}{}
	}
}

func main() {
	dict := make(map[string]struct{})
	BuildDictionary(dict, "a", "and", "bat", "bath", "bed", "beyond",
		"canal", "hand", "man", "plan")

	for _, s := range []string{"amanaplanacanal", "bedbathandbeyond"} {
		valid, words := FindWords(s, dict)
		fmt.Printf("Is %v a valid sequence of words? %v %v.\n",
			s, valid, words)
	}

}
