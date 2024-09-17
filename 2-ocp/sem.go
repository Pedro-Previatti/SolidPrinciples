package ocp

type ContratoClt struct {
	salarioBase float64
}

func (c *ContratoClt) Salario() float64 {
	return c.salarioBase
}

type Estagio struct {
	bolsaAuxilio float64
}

func (e *Estagio) BolsaAuxilio() float64 {
	return e.bolsaAuxilio
}

type FolhaDePagamento struct {
	saldo float64
}

func (f *FolhaDePagamento) Calcular(funcionario interface{}) {
	switch v := funcionario.(type) {
	case *ContratoClt:
		f.saldo = v.Salario()
	case *Estagio:
		f.saldo = v.BolsaAuxilio()
	}
}
