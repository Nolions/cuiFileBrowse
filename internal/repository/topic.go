package repository

import "github.com/Nolions/cuiFileBrowser/internal/model"

// InsertTopic 寫入影片主題
func (repo *Repository) InsertTopic(m *model.Topic) (int64, error) {
	return repo.db.Insert(m)
}

// DeleteTopic 刪除主題資料
func (repo *Repository) DeleteTopic(id int64) (int64, error) {
	m := new(model.Topic)
	return repo.db.Where("id=?", id).Delete(m)
}

// FindTopic 取得主題資料
func (repo *Repository) FindTopic(id int64) (*model.Topic, error) {
	var topic model.Topic
	b, err := repo.db.Where("id=?", id).Get(&topic)

	if err != nil {
		return nil, err
	} else if !b {
		return nil, nil // data not found
	} else {
		return &topic, nil
	}
}

// UpdateTopic 更新主題資料
func (repo *Repository) UpdateTopic(model *model.Topic) (int64, error) {
	return repo.db.Where("id=?", model.Id).Update(model)
}
