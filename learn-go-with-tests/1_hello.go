package main

import "fmt"

func main() {

	fmt.Println(Hello())
}

func Hello() string {

	return "Hello"
}

const (
	HELLO = "Hello "
	EN    = "en"
	FA    = "fa"
	FR    = "fr"
)

func Hello1(name string) string {
	return HELLO + name
}

func Hello2(name string) string {

	if name == "" {
		name = "world"
	}
	return HELLO + name
}

func Hello3(name string, language string) string {

	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + " " + HELLO + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case FR:
		prefix = FR
	case FA:
		prefix = FA
	default:
		prefix = EN
	}
	return
}
