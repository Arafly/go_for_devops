package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

// Pass in language as a parameter
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

// greetingPrefix returns the appropriate greeting prefix based on the specified language.
// If the language is "Español", it returns the Spanish greeting prefix.
// If the language is "Français", it returns the French greeting prefix.
// For any other language, it returns the English greeting prefix.
func greetingPrefix(language string) (prefix string){
	switch language {
	case "Español":
		prefix = spanishHelloPrefix
	case "Français":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return

}

func main() {
	fmt.Println(Hello("world", ""))
}