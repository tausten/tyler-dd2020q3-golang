package main

import "fmt"

func HelloWorld() string {
	return Hello(worldName)
}

const englishHelloPrefix = "Hello, "
const worldName = "world"

func Hello(name string) string {
	if name == "" {
		name = worldName
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(HelloWorld())
}
