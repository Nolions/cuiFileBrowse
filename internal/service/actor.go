package service

import (
	"errors"
	"github.com/Nolions/cuiFileBrowser/internal/model"
	"log"
)

// CreateActor 建立演員資料
func (s *Serv) CreateActor(model *model.Actor) error {
	_, err := s.Repo.InsertActor(model)
	if err != nil {
		// TODO
		log.Printf("create actor fail, error: %v\n", err.Error())
		return err
	}
	return nil
}

// DeleteActor 刪除演員資料
func (s *Serv) DeleteActor(id int64) error {
	m, err := s.GetActor(id)
	if err != nil {
		// TODO
		log.Printf("edit actor fail, get actor error: %v\n", err.Error())
		return err
	}

	_, err = s.Repo.DeleteActor(m.Id)
	if err != nil {
		// TODO
		log.Printf("delete actor fail, error: %v\n", err.Error())
		return err
	}
	return nil
}

// EditActor 編輯演員資料
func (s *Serv) EditActor(id int64, name string) error {
	m, err := s.GetActor(id)
	if err != nil {
		// TODO
		log.Printf("edit actor fail, get actor error: %v\n", err.Error())
		return err
	}

	var newModel model.Actor
	newModel = *m
	newModel.Name = name

	_, err = s.Repo.UpdateActor(&newModel)
	if err != nil {
		// TODO
		log.Printf("edit actor fail, error: %v\n", err.Error())
		return err
	}
	return nil
}

// GetActor 取得演員資料
func (s *Serv) GetActor(id int64) (*model.Actor, error) {
	m, err := s.Repo.FindActor(id)
	if err != nil {
		// TODO
		log.Printf("get actor fail, error: %v\n", err.Error())
		return nil, err
	}

	if m == nil {
		// TODO
		e := errors.New("no data found")
		log.Printf("edit actor fail, error: %v\n", e.Error())
		return nil, errors.New("no data")
	}

	return m, nil
}
