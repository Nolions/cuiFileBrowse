package model

import (
	"database/sql"
	"time"
)

type Video struct {
	Id       int64
	TopicId  sql.NullInt64 `xorm:"topic_id"`
	Name     string        `xorm:"name"`
	CreateAt time.Time     `xorm:"create_at"`
	UpdateAt time.Time     `xorm:"update_at"`
}

func (Video) TableName() string {
	return "videos"
}

// VideoTopic videos table relational topics table
type VideoTopic struct {
	Id       int64
	Name     string        `xorm:"name"`
	CreateAt time.Time     `xorm:"create_at"`
	UpdateAt time.Time     `xorm:"update_at"`
	TopicId  sql.NullInt64 `xorm:"topic_id"`
	Topic    `xorm:"extends"`
}

func (VideoTopic) TableName() string {
	return "videos"
}
