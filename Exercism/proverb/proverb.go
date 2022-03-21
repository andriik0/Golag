// Package proverb make proverbs
package proverb

import "fmt"

// Proverb gets rhyme as slice of string and returns proverb as slice of string 
func Proverb(rhyme []string) []string {
	rhymelen := len(rhyme)
	if rhymelen == 0 {
		return []string{}
	}
	var proverb []string = make([]string, 0)
	if rhymelen > 1 {
		for i := 0; i < rhymelen-1; i++ {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}
	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return proverb
}
