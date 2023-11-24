package main

import "fmt"

const englishPrefix = "Hello"
const polishPrefix = "Witam"
const chinesePrefix = "你好"

func greetingPrefix(lang string) (prefix string) {

	switch lang {
	case "polish":
		prefix = polishPrefix
	case "中文":
		prefix = chinesePrefix
	default:
		prefix = englishPrefix
	}

	return
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + ", " + name + "!"
}

func main() {
	fmt.Println(Hello("Stefan", "english"))
}
