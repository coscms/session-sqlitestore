package sqlite3

import (
	"os"
	"testing"
	"time"

	//_ "github.com/coscms/session-sqlitestore/driver"
	//_ "github.com/coscms/session-sqlitestore/driver/cgo"
	_ "github.com/coscms/session-sqlitestore/driver/go"
	sqlstore "github.com/coscms/session-sqlstore"
	"github.com/stretchr/testify/require"
)

func TestXxx(t *testing.T) {
	store, err := NewSQLiteStore(&Options{
		Path:    `./db.test`,
		Options: sqlstore.Options{},
	})
	require.NoError(t, err)
	quit, done := store.Cleanup(time.Second)
	time.Sleep(time.Second * 5)
	store.StopCleanup(quit, done)
	store.Close()
	os.Remove(`./db.test`)
}
