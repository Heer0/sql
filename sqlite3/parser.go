package sqlite3

import (
	"github.com/Heer0/sql"
	"strings"
)

// ParseDSN parses the sqlite3 datasource name.
func ParseDSN(name string) sql.DSNInfo {
	if pos := strings.IndexRune(name, '?'); pos >= 0 {
		name = name[:pos]
	}
	return sql.DSNInfo{
		Database: name,
	}
}
