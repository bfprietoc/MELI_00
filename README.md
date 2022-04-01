

# Ejercicio Mutantes - Mercadolibre_TEST
Prueba tecnica para para MercadoLibre (Backend). 

### Para tener en cuenta
Me propuse realizar este reto en el lenguaje GOLANG, fue una decision que tome teniendo en cuenta que ya habia realizado este tipo de implementaciones en lenguajes como JAVA (Spring-boot), node.Js, Python o Ruby (On rails) y  me parecio una oportunidad genial para demostrar mis capacidades de aprendizaje y adaptacion a problemas rapidamente. Se uso una base de datos MySql para almacenar la informacion.

## Para ejecutar en una maquina local:

## Base de datos (Mysql)
1. Es necesario tener una base de datos Mysql
2. Conectarse a la base de datos mysql
3. Crear la base de datos y la tabla con los campos necesarios, para eso se ejecutan las siquientes querys
```
CREATE DATABASE meli;
USE meli;
CREATE TABLE `stats` (`chain` varchar(255) NOT NULL PRIMARY KEY, `mutant` BOOLEAN, `non_mutant` BOOLEAN);
```
4. La base de datos se expone por defecto en el puerto 3306, asi que usamos el mismo para la conexion.
5. Finalmente se creo una instancia RDS de AWS para la creacion de la base de datos mysql.

## App Go
1. Descargar o clonar el repositorio a la maquina donde se desea correr
2. Ubicarse en el directorio MELI_00/
3. Ejecutar el comando go build
4. Luego se puede correr la app con el comando ./meli_00 o go run main.go
5. La aplicacion se ejecutara en la ruta http://localhost:1234/


## Docker 
1. Se creo una imagen de docker de la aplicacion desarrollada en GO, se subio la imagen a dockerhub.
3. Se creo una instancia EC2 en AWS y se desplego la imagen de docker alli.
4. Opcionalmente se puede descargar la imagen de docker para su uso local con el siguiente comando docker pull bfprietoc/meli
5. Para correr la imagen: docker run -p 8080:1234 bfprietoc/meli:1.0



## Implementacion y tecnologias usadas

- [Golang 1.8](https://go.dev/)
- [MySql](https://www.mysql.com)
- [Mux](https://github.com/gorilla/mux)
- [Docker](https://www.docker.com)
- [AWS](https://aws.amazon.com/)




Enlaces de AWS donde se encuentran los dos servicios principales:

- [IsMutant](http://ec2-34-224-56-190.compute-1.amazonaws.com:8080/mutant) http://ec2-34-224-56-190.compute-1.amazonaws.com:8080/mutant
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

- [stats](http://ec2-34-224-56-190.compute-1.amazonaws.com:8080/stats) : http://ec2-34-224-56-190.compute-1.amazonaws.com:8080/stats

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

### Test Unitarios

Se realizaron test unitarios con la libreria "testing" de go, se creo una estructura con 12 diferentes entradas, todas pasaron exitosamente con un tiempo de respuesta de 0.100 ms aproximadamente.









