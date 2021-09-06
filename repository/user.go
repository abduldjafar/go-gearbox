package repository

import "go-gearbox/entity"

type user struct{}

func (*user) CreateUser(entity.User) error {
	return nil
}
func (*user) GetUserByID(id string) entity.User {
	return entity.User{}
}
func (*user) GetUser(params ...interface{}) []entity.User {
	return []entity.User{}
}
func (*user) UpdateUser(entity.User) error {
	return nil
}
func (*user) DeleteUser(id string) error {
	return nil
}

func ImplUserRepository() User {
	return &user{}
}
