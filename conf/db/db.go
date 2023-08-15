package db

import "github.com/Nolions/cuiFileBrowser/conf/db/sqlite"

type Conf struct {
	SQLite sqlite.Config `mapstructure:"sqlite"`
}
