package main

import (
  "fmt"
  contas "banc/contas"
  //clientes "banc/clientes"
)

func PagarBoleto(conta VerificarConta, valor float64){
  conta.Sacar(valor)
}

type VerificarConta interface{
  Sacar(valor float64) string
}

func main(){
  countP := contas.ContaPoupanca{}
  countP.Depositar(7000)
  PagarBoleto(&countP, 7000)
  fmt.Println(countP.Saldo())

  countL := contas.ContaCorrente{}
  countL.Depositar(10000)
  PagarBoleto(&countL, 10000)
  fmt.Println(countL.Saldo())
}
