package lsp

type Bird interface {
	Fly() string
}

type Pigeon struct{}

func (p *Pigeon) Fly() string {
	return "Pigeon is flying."
}

type Chicken struct{}

func (p *Chicken) Fly() string {
	return "Chicken is flying."
}
