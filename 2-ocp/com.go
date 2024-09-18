package ocp

type Funcionario interface {
	CalcularPagamento() float64
}

type ContratoCltOCP struct {
	salarioBase float64
}

func (c *ContratoCltOCP) CalcularPagamento() float64 {
	return c.salarioBase
}

type EstagioOCP struct {
	bolsaAuxilio float64
}

func (e *EstagioOCP) CalcularPagamento() float64 {
	return e.bolsaAuxilio
}

type FolhaDePagamentoOCP struct {
	saldo float64
}

func (f *FolhaDePagamentoOCP) Calcular(funcionario Funcionario) {
	f.saldo = funcionario.CalcularPagamento()
}
