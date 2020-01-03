package pq

import (
	pq "github.com/lib/pq"
	"github.com/Heer0/sql"
)

func init() {
	sql.Register("postgres", &pq.Driver{}, sql.WithDSNParser(ParseDSN))
}
