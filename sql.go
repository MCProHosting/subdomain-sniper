package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var connection *sql.DB = nil

func getDsn() string {
	config := getConfig()

	dsn := config.MysqlUser
	if config.MysqlPass != "" {
		dsn += "+" + config.MysqlPass
	}
	dsn += "@tcp(" + config.MysqlHost + ")/" + config.MysqlDb

	return dsn
}

func getConnection() *sql.DB {
	if connection == nil {
		db, err := sql.Open("mysql", getDsn())

		if err != nil {
			log.Fatal(err)
		}

		connection = db
	}

	return connection
}

func getStatements() (stmtDomain *sql.Stmt, stmtDelete *sql.Stmt) {
	db := getConnection()

	stmtDomain, err := db.Prepare("SELECT `id` FROM `msql_subdomain_domains` WHERE `domain` = ?")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	stmtDelete, _ = db.Prepare("DELETE FROM `msql_subdomains` WHERE `domain_id` = ? AND `name` = ?")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return stmtDomain, stmtDelete
}

func deleteSqlSubdomain(subdomain string, zone string) {
	domain, delete := getStatements()
	defer domain.Close()
	defer delete.Close()

	var id int
	err := domain.QueryRow(zone).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	if id == 0 {
		log.Println("Could not find the domain in the database.")
		return
	}

	delete.Exec(id, subdomain)
	log.Println("Removed the subdomain from the database.")
}
