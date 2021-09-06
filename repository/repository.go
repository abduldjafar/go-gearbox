package repository

import "go-gearbox/entity"

type User interface {
	CreateUser(entity.User) error
	GetUserByID(id string) entity.User
	GetUser(params ...interface{}) []entity.User
	UpdateUser(entity.User) error
	DeleteUser(id string) error
}
