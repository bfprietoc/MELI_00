package main

import (
	c "meli_00/controller"
	"testing"
)

func TestIsMutant(t *testing.T) {

	var arr00 = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"} // mutante del enunciado
	var test00 = c.ToMatrix(arr00)

	var arr01 = []string{"TTGCGA", "CAGTGC", "TTATGT", "AGAACG", "CGCCTA", "TCACTC"}
	var test01 = c.ToMatrix(arr01)

	var arr02 = []string{"AATAGA", "GAGGCC", "TAGTGG", "AGAAGG", "CTGCTG", "GGTGTA"}
	var test02 = c.ToMatrix(arr02)

	var arr03 = []string{"AATAGA", "AAGGCC", "AAGTAG", "AGAAAG", "CCGCAG", "GGTGAG"} //2 verticales
	var test03 = c.ToMatrix(arr03)

	var arr04 = []string{"AATAGA", "AATACC", "TTATGT", "ATAAGG", "CGCCTA", "GTGGTG"} // 1 diag inicio
	var test04 = c.ToMatrix(arr04)

	var arr05 = []string{"AATAGA", "GAGTCC", "TATTTG", "AGATGG", "CCGCTG", "GGTGTT"} //1 Diagonal al final
	var test05 = c.ToMatrix(arr05)

	var arr06 = []string{"AATAGA", "GAGGCC", "TAGTGG", "AGAAGG", "CCGCTG", "GGTGTA"} // No tiene conincidencias
	var test06 = c.ToMatrix(arr06)

	var arr07 = []string{"AATAGA", "GAGGCC", "TAGTGG", "AGAAGG", "CCGCTG", "GGTGTG"} // 1 G vertical final
	var test07 = c.ToMatrix(arr07)

	var arr08 = []string{"GGTAGA", "GGGTCC", "GGGTGG", "GGAAGG", "CCGCTG", "GGTGTT"} // 2 G vertical inicio y G diagonal
	var test08 = c.ToMatrix(arr08)

	var arr09 = []string{"AATAGA", "GAGGCC", "TAGTGG", "AAAAGG", "CCGCTG", "GGGGTA"} // 1 A horixontal 1 G horizontal
	var test09 = c.ToMatrix(arr09)

	var arr10 = []string{"AATAGA", "GGGGCC", "TAGTGG", "AAAAGG", "CCGCTG", "GGGGTA"} // 1 A horixontal 1 G horizontal
	var test10 = c.ToMatrix(arr10)

	var arr11 = []string{"ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"} // No mutante el enunciado
	var test11 = c.ToMatrix(arr11)

	var tests = []struct {
		input    [][]string
		expected bool
	}{
		{test00, true},
		{test01, false},
		{test02, false},
		{test03, false},
		{test04, true},
		{test05, true},
		{test06, false},
		{test07, true},
		{test08, false},
		{test09, true},
		{test10, false},
		{test11, false},
	}
	for _, test := range tests {
		if output := c.IsMutant(test.input); output != test.expected {
			t.Error("Test Failed: {} inputed, {} expected , recieved: {}", test.input, test.expected, output)

		}
	}
}
