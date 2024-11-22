//go:build !sqlitego && !session_sqlitego

package driver

import (
	_ "github.com/coscms/session-sqlitestore/driver/cgo"
)
