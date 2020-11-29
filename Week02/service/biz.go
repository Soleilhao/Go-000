package service

//type biz.go
type UserBiz struct{}

func (u *UserBiz) GetUser(id int) (User, error) {
	d := UserDao{}
	return d.GetUser(id)
}
