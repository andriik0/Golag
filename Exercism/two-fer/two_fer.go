// Package twofer implements function ShareWith
package twofer

import "fmt"

// ShareWith given a name, return a string with the message
func ShareWith(name string) string {
	if len(name) > 0 {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
	return "One for you, one for me."
}
