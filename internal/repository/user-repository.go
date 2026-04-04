package repository

import (
	"errors"
)

type UserRep struct {
	users map[string]*User
}

func NewUserRep() *UserRep {
	return &UserRep{
		users: make(map[string]*User),
	}
}

func (r *UserRep) Create(user *User) error {

	if _, ok := r.users[user.Username]; ok {
		return errors.New("User already exist")
	}

	r.users[user.Username] = user
	return nil
}

func (r *UserRep) ExistByUserName(name string) bool {
	_, exist := r.users[name]
	return exist
}
