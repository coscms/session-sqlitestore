//go:build !sqlitecgo

package driver

import (
	_ "github.com/coscms/session-sqlitestore/driver/go"
)
