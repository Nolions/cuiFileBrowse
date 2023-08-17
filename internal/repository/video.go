package repository

import (
	"github.com/Nolions/cuiFileBrowser/internal/model"
)

// InsertVideo 寫入影片主題
func (repo *Repository) InsertVideo(m *model.Video) (int64, error) {
	return repo.db.Insert(m)
}

// DeleteVideo 刪除主題資料
func (repo *Repository) DeleteVideo(id int64) (int64, error) {
	m := new(model.Video)
	return repo.db.Where("id=?", id).Delete(m)
}

// FindVideo 取得影片
func (repo *Repository) FindVideo(id int64) (*model.Video, error) {
	var video model.Video

	b, err := repo.db.Where("id=?", id).Get(&video)

	if err != nil {
		return nil, err
	} else if !b {
		return nil, nil // data not found
	} else {
		return &video, nil
	}
}

// FindVideoTopic 取得主題資料
func (repo *Repository) FindVideoTopic(id int64) (*model.VideoTopic, error) {
	var video model.VideoTopic
	b, err := repo.db.Where("videos.id=?", id).
		Join("LEFT OUTER", "topics", "topics.id=videos.topic_id").
		Get(&video)

	if err != nil {
		return nil, err
	} else if !b {
		return nil, nil // data not found
	} else {
		return &video, nil
	}
}

// UpdateVideo 更新主題資料
func (repo *Repository) UpdateVideo(model *model.Video) (int64, error) {
	return repo.db.Where("id=?", model.Id).Update(model)
}
