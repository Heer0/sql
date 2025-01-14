package sqlite3

import (
	sqlite3 "github.com/mattn/go-sqlite3"
	"github.com/Heer0/sql"
)

func init() {
	sql.Register("sqlite3", &sqlite3.SQLiteDriver{}, sql.WithDSNParser(ParseDSN))
}
