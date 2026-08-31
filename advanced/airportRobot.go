package advanced

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

type GermanGreeter struct {
}

func (gg GermanGreeter) Greet(name string) string {
	return "Hallo " + name + " !"
}

func (gg GermanGreeter) LanguageName() string {
	return "German"
}

type Italian struct {
}

func (it Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

func (it Italian) LanguageName() string {
	return "Italian"
}

type Portuguese struct {
}

func (pt Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}

func (pt Portuguese) LanguageName() string {
	return "Portuguese"
}

func SayHello(name string, greeter Greeter) string {
	return "I can speak " + greeter.LanguageName() + ": " + greeter.Greet(name)
}
