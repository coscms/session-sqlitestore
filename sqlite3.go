package sqlite3

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

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
	name := `sqlite`
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
// path - path for Set-Cookie header
func NewSQLiteStore(cfg *Options) (*SQLiteStore, error) {
	var uri string
	if len(cfg.Path) == 0 {
		d := fmt.Sprintf("%d-session", time.Now().Unix())
		uri = filepath.Join(os.TempDir(), d, "sessions.db")
		os.MkdirAll(filepath.Dir(uri), 0755)
		uri += `?tmp=true`
	} else {
		os.MkdirAll(filepath.Dir(cfg.Path), 0755)
		uri = cfg.Path
	}
	db, err := sql.Open("sqlite3", "file:"+uri)
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
