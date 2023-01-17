package database

import (
	"database/sql"
	"final-project/service"
	"fmt"

	packr "github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	Connection *sql.DB
)

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./migration"),
	}
	number, err := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	service.CheckError(err)

	Connection = dbParam
	fmt.Println("Applied", number, "migrations")
}
