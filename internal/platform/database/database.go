package database

import (
	"context"
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

type contextKey int

const (
	txKey contextKey = iota
)

// With returns a Builder that can be used to build and execute SQL queries.
// With will return the transaction if it is found in the given context.
// Otherwise it will return a DB connection associated with the context.
func (db *DB) With(ctx context.Context) dbx.Builder {
	if tx, ok := ctx.Value(txKey).(*dbx.Tx); ok {
		return tx
	}
	return db.query.WithContext(ctx)
}

// New returns a new DB connection that wraps the given dbx.DB instance.
func New(db *dbx.DB) *DB {
	return &DB{db}
}

// Initialize db connection and return DB connection
func Init(uri string, driver string) *DB {
	db, err := dbx.MustOpen(driver, uri)

	if err != nil {
		log.Fatal("ERR - Open DB failed: ", err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Fatal("ERR - Ping DB failed: ", err)
	}

	return New(db)
}
