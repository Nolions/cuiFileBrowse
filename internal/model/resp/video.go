package resp

type Video struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Topic  Topic   `json:"topic"`
	Actors []Actor `json:"actors"`
}

type Topic struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Actor struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
