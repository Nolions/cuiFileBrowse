package model

type VideoActor struct {
	ActorId int64 `xorm:"actor_id"`
	VideoId int64 `xorm:"video_id"`
}

func (VideoActor) TableName() string {
	return "video_actors"
}

type VideoActorRelationship struct {
	ActorId int64 `xorm:"actor_id"`
	VideoId int64 `xorm:"video_id"`
	Actor   `xorm:"extends"`
	Video   `xorm:"extends"`
}

func (VideoActorRelationship) TableName() string {
	return "video_actors"
}
