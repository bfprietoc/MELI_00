package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"meli_00/config"
	m "meli_00/model"
	"net/http"
	"regexp"
	"strings"
)

// Expresion regular para las validaciones iniciales.
var re = regexp.MustCompile(`[^ACGT]`)

//Metodo POST que recibe la informacion a procesar.
func Mutant(w http.ResponseWriter, r *http.Request) {
	var response m.ResponseDna
	dna := m.Dna{}
	db := config.Connect()
	defer db.Close()

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&dna)
	if err != nil {
		fmt.Println(err)
		return
	}

	var mutant bool
	chain := dna.Dna
	chainString := strings.Join(chain, "")

	//Validar cantidad de caracteres (36) y que no sean diferentes de ACGT con la expresion regular
	if len(chainString) != 36 {
		err := "DNA ingresado no cumple con el tamano"
		panic(err)
	} else if re.MatchString(chainString) {
		err := "DNA ingresado posee caracteres incorrectos"
		panic(err)
	}

	arr := ToMatrix(chain)

	if IsMutant(arr) {
		mutant = true
		fmt.Println(chain, mutant)
		//Prepara la query validando errores
		insForm, err := db.Prepare("INSERT INTO stats (chain, mutant) VALUES (?, ?)")
		if err != nil {
			panic(err)
		} else {
			log.Println("Insert mutan succes")
		}
		//Ejecuta la query
		insForm.Exec(chainString, mutant)
		response.Status = 200
		response.Message = "HTTP 200-OK"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	} else {
		mutant = true
		insForm, err := db.Prepare("INSERT INTO stats (chain, non_mutant) VALUES (?, ?)")
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Insert no-mutan succes")
		}
		insForm.Exec(chainString, mutant)
		response.Status = 403
		response.Message = "403-Forbidden"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}

//Metodo GET que muestra las estadisticas
func Stats(w http.ResponseWriter, r *http.Request) {
	var response m.ResponseStats
	var stats m.Stats
	var arrStats []m.Stats

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT count(mutant),count(non_mutant),(count(mutant)/count(non_mutant)) from stats;")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&stats.Count_mutant_dna, &stats.Count_human_dna, &stats.Ratio)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrStats = append(arrStats, stats)
		}
	}

	fmt.Println(arrStats)

	response.Status = 200
	response.Message = "Success"
	response.Stats = arrStats

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
