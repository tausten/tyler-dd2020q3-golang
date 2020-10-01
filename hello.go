package main

import "fmt"

func HelloWorld() string {
	return Hello(worldName, "")
}

const (
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

const (
	frenchLanguage  = "French"
	spanishLanguage = "Spanish"
)

const worldName = "world"

func Hello(name string, language string) string {
	if name == "" {
		name = worldName
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case frenchLanguage:
		prefix = frenchHelloPrefix
	case spanishLanguage:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(HelloWorld())
}
