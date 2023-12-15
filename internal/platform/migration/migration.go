package migration

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbSource string) {
	migration, err := migrate.New("file://internal/platform/migration/files", dbSource)

	if err != nil {
		log.Fatal("ERR - Migration conn: ", err)
	}

	if err := migration.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migration change detected!")

			return
		}

		log.Fatal("ERR - Migration failed: ", err)
	}
}
