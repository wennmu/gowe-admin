package model

type Admin struct {
}

func (a Admin) UserInfo(username, password string) int64 {
	var uid int64
	uid = 1
	return uid
}
