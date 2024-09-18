package lsp

type BirdLSP interface {
	MakeSound() string
}

type FlyingBird interface {
	BirdLSP
	Fly() string
}

type PigeonLSP struct{}

func (p *PigeonLSP) MakeSound() string {
	return "Coo"
}

func (p *PigeonLSP) Fly() string {
	return "Pigeon is flying."
}

type ChickenLSP struct{}

func (p *ChickenLSP) MakeSound() string {
	return "Cocorico"
}
