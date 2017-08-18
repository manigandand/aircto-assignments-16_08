package main

import (
	"database/sql"
	_ "database/sql/driver"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var createTableStatements = []string{
	`CREATE DATABASE IF NOT EXISTS aricto DEFAULT CHARACTER SET = 'utf8' DEFAULT COLLATE 'utf8_general_ci';`,
	`USE aricto;`,
	`CREATE TABLE IF NOT EXISTS user (
	  id int(11) NOT NULL AUTO_INCREMENT,
	  email varchar(45) NOT NULL,
	  user_name varchar(45) NOT NULL,
	  first_name varchar(45) DEFAULT NULL,
	  last_name varchar(45) DEFAULT NULL,
	  password varchar(45) DEFAULT NULL,
	  access_token varchar(250) DEFAULT NULL,
	  PRIMARY KEY (id),
	  UNIQUE KEY email_UNIQUE (email)
	) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1`,
	`LOCK TABLES user WRITE`,
	`INSERT INTO user VALUES (1,'manigandan.jeff@gmail.com','manigandanjeff','Manigandan','Dharmalingam','qwerty@123','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.1Gvg0ahLLUKTdyBBR-KMOEOu8fnl24UF2_71MiVZdKU'),(2,'manigandan2293@gmail.com','employee.one','Employee','One','qwerty123','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.1Gvg0ahLLUKTdyBBR-KMOEOu8fnl24UF2_71MiVZdKU'),(3,'manigandan.jeff1@gmail.com','emplyee.two','Employee Two','Aricto','qwerty12345','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.1Gvg0ahLLUKTdyBBR-KMOEOu8fnl24UF2_71MiVZdKU'),(4,'manigandaninnovator@gmail.com','aricto','Aricto','Admin','qwerty!123','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.1Gvg0ahLLUKTdyBBR-KMOEOu8fnl24UF2_71MiVZdKU')`,
	`UNLOCK TABLES`,
	`CREATE TABLE IF NOT EXISTS issues (
	  id int(11) NOT NULL AUTO_INCREMENT,
	  title varchar(45) NOT NULL,
	  description varchar(250) NOT NULL,
	  assigned_to int(11) NOT NULL,
	  created_by int(11) NOT NULL,
	  status varchar(45) NOT NULL,
	  PRIMARY KEY (id),
	  UNIQUE KEY id_UNIQUE (id)
	) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1`,
	`LOCK TABLES issues WRITE`,
	`INSERT INTO issues VALUES (1,'Test Issue','This issue is created for test.',2,1,'open'),(2,'Updated Title','updated discription.',3,1,'closed'),(3,'Wrong Salary Info','The candidate salary information showing wrong data.',3,1,'open'),(4,'Wrong Company Info','Some Company Info showing wrongly.',4,1,'open'),(5,'Assign to Mani','Assigning this issue to mani to handle.',1,2,'open')`,
	`UNLOCK TABLES`,
}

/**
* Global Variables
 */
var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "homestead:secret@tcp([192.168.11.11]:3306)/")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("MIGRATION START ---------")
	i := 1
	for _, stmt := range createTableStatements {
		fmt.Println("MIGRATION STATEMENT", i, " ---------")
		_, err := db.Exec(stmt)
		if err != nil {
			panic(err.Error())
		}
		i++
	}

	fmt.Println("MIGRATION COMPLETED ---------")
}
