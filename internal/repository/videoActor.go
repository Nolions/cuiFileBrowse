package repository

import "github.com/Nolions/cuiFileBrowser/internal/model"

// GetActorsByVideo 查詢影片的演員
func (repo *Repository) GetActorsByVideo(id int64) ([]model.VideoActorRelationship, error) {
	var data = make([]model.VideoActorRelationship, 0)
	err := repo.db.Where("video_id=?", id).
		Join("INNER", "actors", "actors.id=videoActors.actor_id").
		Find(&data)

	if err != nil {
		return nil, err
	}

	return data, err
}

// InsertVideoActors 寫入影片關聯演員資料
func (repo *Repository) InsertVideoActors(m *model.VideoActor) (int64, error) {
	return repo.db.Insert(m)
}

// BatchInsertVideoActors 批次寫入影片關聯演員資料
func (repo *Repository) BatchInsertVideoActors(m *[]model.VideoActor) (int64, error) {
	return repo.db.Insert(m)
}

// DeleteVideoActors 刪除影片關聯的演員資料
func (repo *Repository) DeleteVideoActors(videoId int64) (int64, error) {
	m := new(model.VideoActor)
	return repo.db.Where("video_id=?", videoId).Delete(m)
}
