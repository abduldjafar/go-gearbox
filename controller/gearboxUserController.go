package controller

import (
	"github.com/gogearbox/gearbox"
)

type user struct{}

func (*user) FindById() interface{} {
	return func(ctx gearbox.Context) {
		id := ctx.Param("id")
		ctx.SendJSON(map[string]interface{}{"id": id})
	}
}
func (*user) FindAll() interface{} {
	return func(ctx gearbox.Context) {

	}

}
func (*user) Delete() interface{} {
	return func(ctx gearbox.Context) {

	}

}
func (*user) Update() interface{} {
	return func(ctx gearbox.Context) {

	}

}
func (*user) Create() interface{} {
	return func(ctx gearbox.Context) {

	}

}
func ImplUserController() UserController {
	return &user{}
}
