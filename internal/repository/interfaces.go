package repository

type UserRepository interface {
	Create(user *User)
	ExistByUserName(name string)
}