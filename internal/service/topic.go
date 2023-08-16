package service

import (
	"errors"
	"github.com/Nolions/cuiFileBrowser/internal/model"
	"log"
	"time"
)

// CreateTopic 建立主題資料
func (s *Serv) CreateTopic(model *model.Topic) error {
	_, err := s.Repo.InsertTopic(model)

	if err != nil {
		// TODO
		log.Printf("create topic fail, error: %v\n", err.Error())
		return err
	}
	return nil
}

// GetTopic 取得主題資料
func (s *Serv) GetTopic(id int64) (*model.Topic, error) {
	m, err := s.Repo.FindTopic(id)
	if err != nil {
		// TODO
		log.Printf("get topic fail, error: %v\n", err.Error())
		return nil, err
	}

	if m == nil {
		// TODO
		e := errors.New("no data found")
		log.Printf("get topic fail, get topic error: %v\n", e.Error())
		return nil, e
	}

	return m, nil
}

// DeleteTopic 刪除主題
func (s *Serv) DeleteTopic(id int64) error {
	m, err := s.GetTopic(id)
	if err != nil {
		// TODO
		log.Printf("edit topic fail, get topic error: %v\n", err.Error())
		return err
	}

	_, err = s.Repo.DeleteTopic(m.Id)
	if err != nil {
		// TODO
		log.Printf("delete topic fail, error: %v\n", err.Error())
		return err
	}
	return nil
}

// EditTopic 編輯主題資料
func (s *Serv) EditTopic(id int64, name string) error {
	m, err := s.GetTopic(id)
	if err != nil {
		// TODO
		log.Printf("edit topic fail, get topicl error: %v\n", err.Error())
		return err
	}

	var newModel model.Topic
	newModel = *m
	newModel.Name = name
	newModel.UpdateAt = time.Now()

	_, err = s.Repo.UpdateTopic(&newModel)
	if err != nil {
		// TODO
		log.Printf("edit topic fail, error: %v\n", err.Error())
		return err
	}
	return nil
}
