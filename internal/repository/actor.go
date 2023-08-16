package repository

import (
	"github.com/Nolions/cuiFileBrowser/internal/model"
)

// InsertActor 寫入演員資料
func (repo *Repository) InsertActor(model *model.Actor) (int64, error) {
	return repo.db.Insert(model)
}

// DeleteActor 刪除演員資料
func (repo *Repository) DeleteActor(id int64) (int64, error) {
	actor := new(model.Actor)
	return repo.db.Where("id=?", id).Delete(actor)
}

// UpdateActor 更新演員資料
func (repo *Repository) UpdateActor(model *model.Actor) (int64, error) {
	return repo.db.Where("id=?", model.Id).Update(model)
}

// FindActor 取得演員資料
func (repo *Repository) FindActor(id int64) (*model.Actor, error) {
	var actor model.Actor
	b, err := repo.db.Where("id=?", id).Get(&actor)

	if err != nil {
		return nil, err
	} else if !b {
		return nil, nil // data not found
	} else {
		return &actor, nil
	}
}
