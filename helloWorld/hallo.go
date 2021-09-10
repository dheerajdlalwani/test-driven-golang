package main

import "fmt"

const english = "English"
const french = "French"
const germanHalloPrefix = "Hallo "
const englishHalloPrefix = "Hello "
const frenchHalloPrefix = "Bonjour "

func Hallo(name, language string) string {
	if name == "" {
		name = "Ingenieur"
	}
    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case french:
			prefix = frenchHalloPrefix
		case english:
			prefix = englishHalloPrefix
		default:
			prefix = germanHalloPrefix
	}
	return
}

func main() {
    fmt.Println(Hallo("welt", ""))
}