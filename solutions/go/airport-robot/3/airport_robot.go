package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(visitorsName string, l Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", l.LanguageName(), l.Greet(visitorsName))
}

type Italian struct{}

func (Italian) LanguageName() string             { return "Italian" }
func (Italian) Greet(visitorsName string) string { return fmt.Sprintf("Ciao %s!", visitorsName) }

type Portuguese struct{}

func (Portuguese) LanguageName() string             { return "Portuguese" }
func (Portuguese) Greet(visitorsName string) string { return fmt.Sprintf("Olá %s!", visitorsName) }
