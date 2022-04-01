

# Ejercicio Mutantes - Mercadolibre_TEST

Prueba tecnica para para MercadoLibre (Backend). 

Enlaces de AWS donde se encuentran los dos servicios principales:

- [IsMutant](http://ec2-34-224-56-190.compute-1.amazonaws.com:8080/mutant).
* /mutant/
  * Tipo: POST
  * Respuesta para DNA **mutante: 200 (OK)**
  * Respuesta para DNA **humano: 403 (Forbidden)**
  * Acciones, valida el DNA y lo guarda en la base de datos (si es que no se encontraba ya en la misma)
  * formato esperado (JSON): Objeto que contiene un Array de strings llamado dna

Request: 

Request body (caso ADN mutante):

```
  {"dna":["ATGCGA", "CAGGGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"]}
```

Response:

```
  200 OK
```
Request body (caso ADN humano):

```
  {"dna":["AATACT", "CCCAGA", "GGGATT", "AATTCC", "GGATCG", "TCACTG"]}
```

Response:

```
  403 Forbidden
```

- [stats](http://ec2-34-224-56-190.compute-1.amazonaws.com:8080/stats)

* /stats/
  * Tipo: GET
  * Respuesta: 200 OK
  * Acciones: devuelve las estadisticas de cantidad de mutantes y humanos en la base de datos (si es que los hay)

Response: 200 (application/json)

```
{
    count_mutant_dna: 4,
    count_human_dna: 1,
    ratio: "0.8"
}
```

## Para ejecutar en una maquina local:

1. Descargar o clonar el repositorio a la maquina donde se desea correr
2. Ubicarse en el directorio MELI_00/
3. Ejecutar el comando go build
4. Luego correr el archivo main.go con el comando go run main.go
5. La aplicacion se ejecutara en la ruta http://localhost:1234/


## Implementacion y tecnologias usadas

- [Golang 1.8](https://go.dev/)
- [MySql](https://www.mysql.com)
- [Mux](https://github.com/gorilla/mux)
- [Docker](https://www.docker.com)
- [AWS](https://aws.amazon.com/)


## Ejercicio

### Para tener en cuenta
Me propuse realizar este reto en el lenguaje GOLANG, fue una desicion que tome teniendo en cuenta que ya habia realizado este tipo de implementaciones en lenguajes como JAVA (Spring-boot), node.Js, Python o Rubi  (On rails), me parecio una oportunidad genial de mostrar mis capacidades de aprendizaje.


### Test Unitarios










