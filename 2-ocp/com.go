package ocp

import "fmt"

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

func main() {
	clt := &ContratoCltOCP{salarioBase: 5000.0}
	estagiario := &EstagioOCP{bolsaAuxilio: 2000.0}

	folha := &FolhaDePagamentoOCP{}

	folha.Calcular(clt)
	fmt.Printf("Saldo do CLT: %.2f\n", folha.saldo)

	folha.Calcular(estagiario)
	fmt.Printf("Saldo do Estagiário: %.2f\n", folha.saldo)
}
