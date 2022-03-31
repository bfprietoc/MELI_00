package config

import (
	"database/sql"
)

//Tabla en Mysql que guarda los registros a procesar.
//CREATE DATABASE meli;
//USE meli;
//CREATE TABLE `stats` (`chain` varchar(255) NOT NULL PRIMARY KEY, `mutant` BOOLEAN, `non_mutant` BOOLEAN);
//Funcion de conexion a la base de datos Mysql
func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "meli-test"
	dbName := "meli"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(meli-test.chg1woop28kd.us-east-1.rds.amazonaws.com:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
