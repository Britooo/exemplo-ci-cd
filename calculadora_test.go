package main

import "testing"

type cenario struct {
	a        int
	b        int
	esperado int
}

func TestSomar(t *testing.T) {

	cenarios := []cenario{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, c := range cenarios {
		resultado := Somar(c.a, c.b)
		if resultado != c.esperado {
			t.Errorf("resultado '%d', esperado '%d'",
				resultado, c.esperado)
		}
	}
}
