package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(a string) string
}

type Italian struct {
}

type Portuguese struct {
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s!", greeter.LanguageName(), greeter.Greet(name))
}

func (i Italian) Greet(a string) string {
	return fmt.Sprintf("Ciao %s", a)
}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Portuguese) Greet(a string) string {
	return fmt.Sprintf("Olá %s", a)
}

func (i Portuguese) LanguageName() string {
	return "Portuguese"
}
