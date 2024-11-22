//go:build sqlitego || session_sqlitego

package sqlite3

import (
	_ "github.com/glebarez/go-sqlite"        // SQLite3 driver. sqlite://
	_ "github.com/glebarez/go-sqlite/compat" // SQLite3 driver. sqlite3://
)