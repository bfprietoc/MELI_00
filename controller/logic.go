package controller

import (
	"fmt"
	"log"
	"strings"
)

//Funcion que dado un arreglo retorna una matriz NxN
func ToMatrix(dna []string) [][]string {
	arr := [][]string{}
	for _, s := range dna {
		split := strings.Split(s, "")
		arr = append(arr, split)
	}
	return arr
}

//Funcion que valida si en un arreglo de strings existe una determinada cadena
//Retorna true si lo encuentra o false si no.
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func IsMutant(dna [][]string) bool {
	var horVal bool = false
	var vertVal bool = false
	var diagVal bool = false
	var doubleVal bool = false
	var horCount int = 0
	var vertCount int = 0
	var diagCount int = 0
	arrHorCount := []string{}
	arrVertCount := []string{}
	arrDiagCount := []string{}

	//Encuentra repeticiones en las horizontales
	for i := 0; i < len(dna); i++ {
		for j := 0; j < len(dna[0])-3; j++ {
			if dna[i][j] == dna[i][j+1] && dna[i][j] == dna[i][j+2] && dna[i][j] == dna[i][j+3] {
				var coinc string = dna[i][j] + dna[i][j+1] + dna[i][j+2] + dna[i][j+3]
				if Contains(arrHorCount, coinc) {
					log.Println("Error palabra mas de una vez")
					horVal = false
					doubleVal = true
				} else {
					arrHorCount = append(arrHorCount, coinc)
					horCount++
				}

			}

		}

		if doubleVal == false && horCount != 0 {
			horVal = true
		} else {
			horVal = false
		}

		if doubleVal == true { //Sale del loop doonde exista un doble registro
			horVal = false
			break
		}

	}

	//Encuentra repeticiones en las Verticales
	for j := 0; j < len(dna); j++ {
		for i := 0; i < len(dna[0])-3; i++ {
			if dna[i][j] == dna[i+1][j] && dna[i][j] == dna[i+2][j] && dna[i][j] == dna[i+3][j] {
				var coinc string = dna[i][j] + dna[i+1][j] + dna[i+2][j] + dna[i+3][j]
				if Contains(arrVertCount, coinc) {
					log.Println("Error palabra mas de una vez")
					vertVal = false
					doubleVal = true
				} else {
					arrVertCount = append(arrVertCount, coinc)
					vertCount++
				}
			}
		}

		if doubleVal == false && vertCount != 0 {
			vertVal = true
		} else {
			vertVal = false
		}

		if doubleVal == true { //Sale del loop doonde exista un doble registro
			vertVal = false
			break
		}
	}

	//Encuentra repeticiones en la diagonal principal
	for i := 0; i < len((dna[0]))-3; i++ {
		for j := 0; j < len(dna[0])-3; j++ {
			if dna[i][j] == dna[i+1][j+1] && dna[i][j] == dna[i+2][j+2] && dna[i][j] == dna[i+3][j+3] {
				fmt.Println("Diagonal Principal: "+dna[i][j], "con pos:", i, j, dna[i+1][j+1], "con pos: ", i+1, j+1, dna[i+2][j+2], "con pos: ", i+2, j+2, dna[i+3][j+3], "con pos: ", i+3, j+3)
				var coinc string = dna[i][j] + dna[i+1][j+1] + dna[i+2][j+2] + dna[i+3][j+3]
				if Contains(arrDiagCount, coinc) {
					log.Println("Error palabra mas de una vez")
					diagVal = false
					doubleVal = true
				} else {
					arrDiagCount = append(arrDiagCount, coinc)
					diagCount++
				}
			}
		}

		if doubleVal == false && diagCount != 0 {
			diagVal = true
		} else {
			diagVal = false
		}

		if doubleVal == true { //Sale del loop doonde exista un doble registro
			diagVal = false
			break
		}

	}

	if (horVal == true || vertVal == true || diagVal == true) && doubleVal == false {
		return true
	} else {
		return false
	}
}
