package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "

const Spanish = "Spanish"
const English = "English"

const World = "World"

func greetingPrefix(language string) (helloPrefix string) {
    switch language {
    case Spanish:
        helloPrefix = spanishHelloPrefix
    case English:
        helloPrefix = englishHelloPrefix
    default:
        helloPrefix = englishHelloPrefix
    }
    return
}

func Hello(name, language string) string {
    if name == "" {
        name = World
    }

    return greetingPrefix(language) + name
}

func main() {
    fmt.Println(Hello("World", ""))
}
