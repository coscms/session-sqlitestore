package sqlite3

import (
	"database/sql"
	"log"

	"github.com/admpub/sessions"
	sqlstore "github.com/coscms/session-sqlstore"
	ss "github.com/webx-top/echo/middleware/session/engine"
	"github.com/webx-top/echo/middleware/session/engine/file"
)

func New(cfg *Options) sessions.Store {
	eng, err := NewSQLiteStore(cfg)
	if err != nil {
		log.Println("sessions: Operation SQLite failed:", err)
		return file.NewFilesystemStore(&file.FileOptions{
			SavePath:      ``,
			KeyPairs:      cfg.KeyPairs,
			CheckInterval: cfg.CheckInterval,
		})
	}
	return eng
}

func Reg(store sessions.Store, args ...string) {
	name := `sqlite3`
	if len(args) > 0 {
		name = args[0]
	}
	ss.Reg(name, store)
}

func RegWithOptions(opts *Options, args ...string) sessions.Store {
	store := New(opts)
	Reg(store, args...)
	return store
}

type Options struct {
	Path string `json:"path"`
	sqlstore.Options
}

type SQLiteStore struct {
	*sqlstore.SQLStore
}

const DDL = "CREATE TABLE IF NOT EXISTS %s (" +
	"	`id` char(64) PRIMARY KEY," +
	"	`data` longblob NOT NULL," +
	"	`created` int NOT NULL DEFAULT '0'," +
	"	`modified` int NOT NULL DEFAULT '0'," +
	"	`expires` int NOT NULL DEFAULT '0');"

// NewSQLiteStore takes the following paramaters
// endpoint - A sql.Open style endpoint
// tableName - table where sessions are to be saved. Required fields are created automatically if the table doesnot exist.
// path - path for Set-Cookie header
// maxAge
// codecs
func NewSQLiteStore(cfg *Options) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", cfg.Path)
	if err != nil {
		return nil, err
	}

	return NewSQLiteStoreFromConnection(db, cfg)
}

// NewSQLiteStoreFromConnection .
func NewSQLiteStoreFromConnection(db *sql.DB, cfg *Options) (*SQLiteStore, error) {
	cfg.Options.SetDDL(DDL)
	base, err := sqlstore.New(db, &cfg.Options)
	if err != nil {
		return nil, err
	}
	s := &SQLiteStore{
		SQLStore: base,
	}
	return s, nil
}
