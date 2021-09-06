package api

import (
	"go-gearbox/controller"
	"go-gearbox/middleware"

	"github.com/gogearbox/gearbox"
)

type GearboxEndpoint struct {
	Router gearbox.Gearbox
}

var (
	gearboxRouter                            = gearbox.New()
	userController controller.UserController = controller.ImplUserController()
	fileController controller.FileController = controller.ImplFileController()
	newMiddleware  middleware.Middleware     = middleware.ImplMiddleware()
)

func (g *GearboxEndpoint) ALL() {

	gearboxRouter.Use(newMiddleware.Logger)

	gearboxRouter.Group("/user", []*gearbox.Route{
		gearboxRouter.Get("/:id", userController.FindById().(func(ctx gearbox.Context))),
		gearboxRouter.Delete("/:id", userController.Delete().(func(ctx gearbox.Context))),
		gearboxRouter.Post("/profile/upload", fileController.Create().(func(ctx gearbox.Context))),
	})

	g.Router = gearboxRouter

}
func (g *GearboxEndpoint) Api() gearbox.Gearbox {
	return g.Router
}
