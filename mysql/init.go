package mysql

import (
	"github.com/go-sql-driver/mysql"
	"github.com/Heer0/sql"
)

func init() {
	sql.Register("mysql", &mysql.MySQLDriver{}, sql.WithDSNParser(ParseDSN))
}
