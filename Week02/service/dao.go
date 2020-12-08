package service

import (
	"database/sql"

	"github.com/pkg/errors"
)

type UserDao struct{}

func (d *UserDao) GetUser(id int) (User, error) {
	var u User
	if id >= 1 && id <= 10 {
		u = User{accountID: id, firstName: "Hello", lastName: "world"}
		return u, nil
	}
	u = User{}
	return u, errors.Wrapf(sql.ErrNoRows, "dao: user not found. (id:%d)", id)
}
