package main

import (
	"strings"
)

// given a list of genres returned by spotify that classify an artist
// return a signular "overarching" genre that fits as many
// subgenres as possible
// do this by:
// 1. if list is empty, return "other"
// 2. return the genre string that appears the most in the other genre strings
//		e.g. ["dance pop", "pop", "pop rap"] -> "pop"
// 3. if no such string exists, return the word that appears the most
//		e.g. ["alternative rock", "indie rock", "modern rock"] -> "rock"
func getMainGenre(genres []string) string {

	// 1. if list is empty, return "Other" 
	if len(genres) == 0 {
		return "other"
	}

	main := ""
	mainCount := 0

	// 2. return the genre string that appears the most in the other genre strings
	for _, candidate := range genres {
		count := 0
		for _, genre := range genres {
			if strings.Contains(genre, candidate) && candidate != genre{
				count += 1
			}
		}

		if count > mainCount{
			main = candidate
			mainCount = count
		}
	}

	if mainCount > 0 {
		return main
	}

	// 3. if no such string exists, return the word that appears the most
	wordMap := map[string]int{}
	for _, genre := range genres {
		for _, word := range strings.Split(genre, " "){
			if _, ok := wordMap[word]; ok {
				wordMap[word] += 1
			} else {
				wordMap[word] = 1
			}
		}
	}

	for word := range wordMap {
		if wordMap[word] > mainCount{
			main = word
			mainCount = wordMap[word]
		}
	}

	return main
}
