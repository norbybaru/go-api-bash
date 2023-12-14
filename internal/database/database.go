package database

import (
	"log"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

// DB represents a DB connection that can be used to run SQL queries.
type DB struct {
	query *dbx.DB
}

// Query returns the dbx.DB wrapped by this object.
func (db *DB) Query() *dbx.DB {
	return db.query
}

func (db *DB) Ping() error {
	return db.query.DB().Ping()
}

// New returns a new DB connection that wraps the given dbx.DB instance.
func New(db *dbx.DB) *DB {
	return &DB{db}
}

// Initialize db connection and return DB connection
func Init(uri string, driver string) *DB {
	db, err := dbx.MustOpen(driver, uri)

	if err != nil {
		log.Fatal("DB ERR: ", err)
	}

	return New(db)
}
