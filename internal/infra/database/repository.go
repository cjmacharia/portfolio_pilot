package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type RepositoryDB struct {
	Client *sqlx.DB
}

func RunMigration(url string, dbSource string) {
	migration, err := migrate.New("file://"+url, "postgres://postgres:@127.0.0.1:5432/portfolio_pilot?sslmode=disable")

	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}
	version, dirty, err := migration.Version()
	if err != nil {
		log.Fatal("Failed to get migration version:", err)
	}

	if dirty {
		log.Printf("Dirty database version %d. Forcing version.\n", version)
		if err := migration.Force(int(version)); err != nil {
			log.Fatal("Failed to force migration version:", err)
		}
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to run migration: ", err)
	}
	fmt.Println("migrations successful")

}

func GetDbClient() *sqlx.DB {
	connectionString := "user=postgres dbname=portfolio_pilot sslmode=disable"
	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	RunMigration("internal/infra/migrations", connectionString)

	return db
}
