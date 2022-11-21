package contas

import contas "banc/clientes"

type ContaPoupanca struct{
  Titular contas.Titular
  NumeroAgencia, NumeroConta   int
  saldo float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
    podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
    if podeSacar {
        c.saldo -= valorDoSaque
        return "Saque realizado com sucesso"
    } else {
        return "saldo insuficiente"
    }
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
    if valorDoDeposito > 0 {
        c.saldo += valorDoDeposito
        return "Deposito realizado com sucesso", c.saldo
    } else {
        return "Valor do deposito menor que zero", c.saldo
    }
}
func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
    if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
        c.saldo -= valorDaTransferencia
        contaDestino.Depositar(valorDaTransferencia)
        return true
    } else {
        return false
    }
}

func (c * ContaPoupanca) Saldo() float64{
  return c.saldo
}
