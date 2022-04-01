

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

### Test 

Se realizaron test unitarios con la libreria "testing" de go, se creo una estructura con 12 diferentes entradas, todas pasaron exitosamente con un tiempo de respuesta de 0.100 ms aproximadamente.


### Casos de prueba POSTMAN:
```

{
    "dna":["ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"]
}

{
    "dna":["ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"]
}

{
    "dna":["AATAGA", "GGGGCC", "TAGTGG", "AAAAGG", "CCGCTG", "GGGGTA"]
}

{
    "dna":["AATAGA", "GAGGCC", "TAGTGG", "AAAAGG", "CCGCTG", "GGGGTA"]
}


{
    "dna":["GGTAGA", "GGGTCC", "GGGTGG", "GGAAGG", "CCGCTG", "GGTGTT"]
}


{
    "dna":["AATAGA", "GAGGCC", "TAGTGG", "AGAAGG", "CCGCTG", "GGTGTG"]
}


{
    "dna":["AATAGA", "GAGGCC", "TAGTGG", "AGAAGG", "CCGCTG", "GGTGTA"]
}


{
    "dna":["AATAGA", "GAGTCC", "TATTTG", "AGATGG", "CCGCTG", "GGTGTT"]
}


{
    "dna":["AATAGA", "AATACC", "TTATGT", "ATAAGG", "CGCCTA", "GTGGTG"]
}


{
    "dna":["AATAGA", "AAGGCC", "AAGTAG", "AGAAAG", "CCGCAG", "GGTGAG"]
}


{
    "dna":["AATAGA", "GAGGCC", "TAGTGG", "AGAAGG", "CTGCTG", "GGTGTA"]
}


{
    "dna":["TTGCGA", "CAGTGC", "TTATGT", "AGAACG", "CGCCTA", "TCACTC"]
}


```


## Notas.

El algoritmo que busca las coincidencias es simpre, pero en el caso de querer optimizarse y aprovechando la facilidad de crear threds (gorutines) con GOLANG, es posible mejorar los tiempos de ejecucion ya que se podria crear una gorutina para los analisis en cada una de las direcciones, pero tener en cuenta el consumo para la creacion de las gorutinas.

Al momento de procesar y hacer la busqueda, el algoritmo implementado no realiza ninguna copia de la matriz entonces el almacenamiento en memoria es optimo.

Analizando el peor de los casos ( Cuando no hay ninguna coincidencia) el algoritmo debe recorrer la matriz en las direcciones establecidad ( Horizontal, vertical y diagonal) por ello en terminos de tiempo de ejecucion tenemos que es lineal O(n) ya que a medida que vayamos aumentando la matriz, este tiempo aumentara tambien.




### Ejercicio 

Magneto quiere reclutar la mayor cantidad de mutantes para poder luchar
contra los X-Men.

Te ha contratado a ti para que desarrolles un proyecto que detecte si un
humano es mutante basándose en su secuencia de ADN.

Para eso te ha pedido crear un programa con un método o función con la siguiente firma (En
alguno de los siguiente lenguajes: Java / Golang / C-C++ / Javascript (node) / Python / Ruby):
* boolean isMutant(String[] dna); // Ejemplo Java

En donde recibirás como parámetro un array de Strings que representan cada fila de una tabla
de (NxN) con la secuencia del ADN. Las letras de los Strings solo pueden ser: (A,T,C,G), las
cuales representa cada base nitrogenada del ADN.

### No-Mutante 
A T G C G A 

C A G T G C

T T A T T T

A G A C G G

G C G T C A

T C A C T G


### Mutante
A T G C G A

C A G T G C

T T A T G T

A G A A G G

C C C C T A

T C A C T G



Sabrás si un humano es mutante, si encuentras más de una secuencia de cuatro letras
iguales​, de forma oblicua, horizontal o vertical.
Ejemplo (Caso mutante):

String[] dna = {"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"};


En este caso el llamado a la función isMutant(dna) devuelve “true”.
Desarrolla el algoritmo de la manera más eficiente posible.
Desafíos:


## Nivel 1:
Programa (en cualquier lenguaje de programación) que cumpla con el método pedido por
Magneto.

## Nivel 2:
Crear una API REST, hostear esa API en un cloud computing libre (Google App Engine,
Amazon AWS, etc), crear el servicio “/mutant/” en donde se pueda detectar si un humano es
mutante enviando la secuencia de ADN mediante un HTTP POST con un Json el cual tenga el
siguiente formato:

POST → /mutant/
{
“dna”:["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}

En caso de verificar un mutante, debería devolver un HTTP 200-OK, en caso contrario un
403-Forbidden

## Nivel 3:
Anexar una base de datos, la cual guarde los ADN’s verificados con la API.
Solo 1 registro por ADN.

Exponer un servicio extra “/stats” que devuelva un Json con las estadísticas de las
verificaciones de ADN: {"count_mutant_dna" : 40, "count_human_dna" : 100, "ratio" : 0.4 }
Tener en cuenta que la API puede recibir fluctuaciones agresivas de tráfico (Entre 100 y 1
millón de peticiones por segundo).
Test-Automáticos, Code coverage > 80%.


