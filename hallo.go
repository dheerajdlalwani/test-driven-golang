package main

import "fmt"


const germanHalloPrefix = "Hallo "

func Hallo(name string) string {
	if name == "" {
		name = "Ingenieur"
	}
    return germanHalloPrefix + name
}

func main() {
    fmt.Println(Hallo("welt"))
}