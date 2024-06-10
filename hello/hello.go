package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix =  "Bonjur, "

func Hello(name ,language string)string{
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string)(prefix string){
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix =  spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
