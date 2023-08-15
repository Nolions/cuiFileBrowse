package service

import (
	"github.com/Nolions/cuiFileBrowser/internal/repository"
	"xorm.io/xorm"
)

type Serv struct {
	Repo *repository.Repository
}

func New(db *xorm.Engine) *Serv {
	repo := repository.New(db)

	return &Serv{
		Repo: &repo,
	}
}
