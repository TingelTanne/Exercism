package airportrobot

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(visitorsName string, language Greeter) string {
	greeting := ("I can speak " +
		language.LanguageName() + ": " +
		language.Greet(visitorsName))
	return greeting
}

type Italian struct {
	languageName, greeting string
}

type Portuguese struct {
	languageName, greeting string
}

func (Italian) LanguageName() string {
	return "Italian"

}

func (Italian) Greet(visitorsName string) string {
	return "Ciao " + visitorsName + "!"
}

func (Portuguese) LanguageName() string {
	return "Portuguese"
}

func (Portuguese) Greet(visitorsName string) string {
	return "Olá " + visitorsName + "!"
}
