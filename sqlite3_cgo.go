//go:build !sqlitego && !session_sqlitego

package sqlite3

import (
	_ "github.com/mattn/go-sqlite3"
)
