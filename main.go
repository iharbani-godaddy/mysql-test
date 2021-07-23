package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	userName := flag.String("u", "", "MySQL username")
	dbName := flag.String("d", "", "MySQL database schema")
	host := flag.String("h", "", "MySQL database server")
	port := flag.Int("p", 0, "MySQL database server port")
	allowOldPasswordsToggle := flag.Int("a", 0, "Allow old passwords")
	timeout := flag.String("t", "30s", "Dial timeout")
	password := os.Getenv("MYSQL_PASSWORD")

	flag.Parse()

	if *userName == "" {
		log.Fatal("-u (username) parameter is empty")
	}

	if *dbName == "" {
		log.Fatal("-d (database name) parameter is empty")
	}

	if *host == "" {
		log.Fatal("-h (host) parameter is empty")
	}

	if *port == 0 {
		*port = 3306
	}

	if password == "" {
		log.Fatal("MYSQL_PASSWORD environment variable is empty")
	}

	allowOldPasswords := ""
	if *allowOldPasswordsToggle == 1 {
		allowOldPasswords = "allowOldPasswords=1"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s&%s", *userName, password, *host, *port, *dbName, *timeout, allowOldPasswords)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 1)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
