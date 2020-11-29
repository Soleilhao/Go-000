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
	// if err != nil {
	// 	return user, fmt.Errorf("service: error occured. \r\n", err)
	// }
	// fmt.Printf("Found user: %#v\r\n", user)
	return user, err
}
