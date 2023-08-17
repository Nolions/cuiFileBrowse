package model

import (
	"time"
)

type Topic struct {
	Id       int64
	Name     string    `xorm:"name"`
	CreateAt time.Time `xorm:"create_at created"`
	UpdateAt time.Time `xorm:"update_at updated"`
}

func (Topic) TableName() string {
	return "topics"
}
