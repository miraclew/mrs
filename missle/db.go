package missle

import (
	"fmt"
)

type Profile struct {
	Id       int64
	Nickname string
	Avatar   string
}

func GetProfile(userId int64) *Profile {
	name := fmt.Sprintf("Name %d", userId)
	p := &Profile{userId, name, "https://angularjs.org/img/AngularJS-large.png"}

	return p
}
