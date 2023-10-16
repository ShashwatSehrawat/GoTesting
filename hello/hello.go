package main

import (
	"fmt"
)

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "hello, "
	spanishHelloPrefix = "hola, "
	frenchHelloPrefix  = "bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := greetingPrefix(language)
	return prefix + name
}
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
func main() {
	fmt.Println(Hello("Chris", "English"))
}
