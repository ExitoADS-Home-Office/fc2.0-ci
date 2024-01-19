package main

import "testing"

func TestSoma(t *testing.T){
	total:=Soma(20,20)

	if total != 40 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado: %d", total, 40)
	}
}