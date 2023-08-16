package repository

import (
	"log"
	"xorm.io/xorm"
)

type Repository struct {
	idx int
	db  *xorm.Engine
}

func New(db *xorm.Engine) Repository {
	return Repository{
		db: db,
	}
}

// Close 關閉與DB的連線
func (repo *Repository) Close() (err error) {
	if repo.db != nil {
		if err = repo.db.Close(); err != nil {
			// TODO
			log.Printf("repository::Close, Repository(%d) failed to close database connection, err = %v\n", repo.idx, err)
		}
		// TODO
		log.Printf("Repository(%d) closed the db connection.\n", repo.idx)
	}

	return
}
