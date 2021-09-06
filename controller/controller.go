package controller

type UserController interface {
	FindById() interface{}
	FindAll() interface{}
	Delete() interface{}
	Update() interface{}
	Create() interface{}
}

type FileController interface {
	Create() interface{}
}
