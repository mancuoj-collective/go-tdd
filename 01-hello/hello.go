package hello

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	chineseHelloPrefix = "你好, "
	spanish            = "Spanish"
	french             = "French"
	chinese            = "Chinese"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	switch language {
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	case chinese:
		return chineseHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}
