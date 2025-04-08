package component

import "time"

type user struct {
	ID        uint32
	UName     string
	UPassword string
	UNickname string
	UKey      string
	ULv       string
	LastIP    string
	CreatedAt time.Time
}

type Shop struct {
	UUID    string
	Name    string
	Mobile  string
	Actived uint
}
