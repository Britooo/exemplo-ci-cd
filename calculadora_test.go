package main

import "testing"

func TestSomar(t *testing.T) {
    resultado := Somar(2, 3)
    esperado := 5

    if resultado != esperado {
        t.Errorf("resultado '%d', esperado '%d'", resultado, esperado)
    }
}