package service

import (
	"database/sql"
	"errors"
	"github.com/Nolions/cuiFileBrowser/internal/model"
	"github.com/Nolions/cuiFileBrowser/internal/model/resp"
	"log"
)

func (s *Serv) CreateVideo(videoName string, topicId int64, actorIds []int64) error {
	var topic *model.Topic
	var err error
	
	// 寫入影片資料
	video := model.Video{
		Name: videoName,
	}

	// 檢查topic是否存在
	if topicId > 0 {
		topic, err = s.GetTopic(topicId)
		if err != nil { // no found topic
			return err
		}
	}
	if topic != nil {
		video.TopicId = sql.NullInt64{Int64: 1, Valid: true}
	}

	// insert data to videos
	_, err = s.Repo.InsertVideo(&video)
	if err != nil {
		// TODO
		log.Printf("create video fail, error: %v\n", err.Error())
		return err
	}

	//
	// 寫入影片關聯的演員資料
	//
	videoActors := make([]model.VideoActor, 0)
	for _, id := range actorIds {
		// 檢查演員是否存在
		_, err := s.GetActor(id)
		if err != nil {
			// TODO
			log.Printf("create video fail, insert video fail, check actor id fail, error:%v\n", err.Error())
			return errors.New("insert video fail, actor id no found")
		}
		videoActors = append(videoActors, model.VideoActor{
			ActorId: id,
			VideoId: video.Id,
		})
	}

	if len(videoActors) > 0 {
		// insert data to video_actors
		_, err = s.Repo.BatchInsertVideoActors(&videoActors)
		if err != nil {
			log.Printf("create video fail, insert video fail, error:%v\n", err.Error())
			return err
		}
	}

	return nil
}

// DeleteVideo 刪除影片資料
func (s *Serv) DeleteVideo(id int64) error {
	m, err := s.Repo.FindVideo(id)
	if err != nil {
		// TODO
		log.Printf("delete video fail, get video error: %v\n", err.Error())
		return err
	}

	_, err = s.Repo.DeleteVideoActors(m.Id)
	if err != nil {
		// TODO
		log.Printf("delete video fail, delete  video relactionship actor fail, error: %v\n", err.Error())
		return err
	}

	if m == nil {
		// TODO
		e := errors.New("no data found")
		log.Printf("delete video fail, get video error: %v\n", e.Error())
		return e
	}

	_, err = s.Repo.DeleteVideo(m.Id)
	if err != nil {
		// TODO
		log.Printf("delete video fail, error: %v\n", err.Error())
		return err
	}
	return nil
}

// GetVideo 取得影片完整資料(含主題與演員列表)
func (s *Serv) GetVideo(id int64) (*resp.Video, error) {
	m, err := s.Repo.FindVideoTopic(id)
	if err != nil {
		// TODO
		log.Printf("get video fail, error: %v\n", err.Error())
		return nil, err
	}

	if m == nil {
		// TODO
		e := errors.New("no data found")
		log.Printf("get video fail, error: %v\n", e.Error())
		return nil, errors.New("no data")
	}

	video := resp.Video{
		Id:   m.Id,
		Name: m.Name,
	}

	// topic
	if m.TopicId.Valid {
		video.Topic = resp.Topic{
			Id:   m.Topic.Id,
			Name: m.Topic.Name,
		}
	}

	// actors of video
	actors, err := s.Repo.GetActorsByVideo(m.Id)
	if err != nil {
		log.Printf("get video fail, get actor for vidoe error: %v\n", err.Error())
	}

	actorSlice := make([]resp.Actor, 0)
	for _, actor := range actors {
		actorSlice = append(actorSlice, resp.Actor{
			Id:   actor.Actor.Id,
			Name: actor.Actor.Name,
		})
	}
	video.Actors = actorSlice

	return &video, err
}
