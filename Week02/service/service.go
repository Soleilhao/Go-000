package service

type User struct {
	accountID int
	firstName string
	lastName  string
}

//service.go

type UserService struct{}

func (u *UserService) GetUser(id int) (User, error) {
	b := UserBiz{}
	user, err := b.GetUser(id)
	return user, err
}
