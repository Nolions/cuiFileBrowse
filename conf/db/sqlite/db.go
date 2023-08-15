package sqlite

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type Config struct {
	Database string `mapstructure:"database"`
}

func New(conf Config) (*xorm.Engine, error) {
	fmt.Println("conf.Database")
	fmt.Println(conf.Database)
	db, err := xorm.NewEngine("sqlite3", conf.Database)

	if err != nil {
		return nil, err
	}

	return db, nil
}
