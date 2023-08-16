package model

import (
	"time"
)

type Actor struct {
	Id       int64
	Name     string    `xorm:"name"`
	CreateAt time.Time `xorm:"create_at"`
	UpdateAt time.Time `xorm:"update_at"`
}

func (Actor) TableName() string {
	return "actors"
}
