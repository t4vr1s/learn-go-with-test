package main

import "fmt"

const englishHelloPrefix string = "Hello, "
const spanishHolaPrefix string = "Hola, "
const frenchBonjourPrefix string = "Bonjour, "
const japanKonishiwaPrefix string = "Konishiwa, "
const spanishLanguage string = "Spanish"
const frenchLanguage string = "French"
const japanLanguage string = "Japan"

func main() {
	fmt.Println(Hello("Eduardo", "Spanish"))
}

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	prefix := englishHelloPrefix

	switch language {
	case spanishLanguage:
		prefix = spanishHolaPrefix
	case frenchLanguage:
		prefix = frenchBonjourPrefix
	case japanLanguage:
		prefix = japanKonishiwaPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}
