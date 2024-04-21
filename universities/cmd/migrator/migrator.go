package main

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"os"
)

var (
	schemasDir string
)

func main() {

	flag.StringVar(&schemasDir, "dir", "migrations", "path to schemas dir")

	flag.Parse()

	fmt.Printf("schemas dir: %s\n", schemasDir)

	host := os.Getenv("UNIVERSITY_DATABASE_HOST")
	port := os.Getenv("UNIVERSITY_DATABASE_PORT")
	user := os.Getenv("UNIVERSITY_DATABASE_USER")
	pass := os.Getenv("UNIVERSITY_DATABASE_PASS")
	name := os.Getenv("UNIVERSITY_DATABASE_NAME")

	cs := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, name)
	fmt.Printf("cs: %s\n", cs)

	m, err := migrate.New("file://"+schemasDir, cs)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		panic(err)
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			fmt.Printf("err: %s\n", err)
			panic(err)
		}
		fmt.Println("no change")
	}
	fmt.Println("migrations done")
}
