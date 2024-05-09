package registry

import (
	"back_end_svc/adapter/controller"
	"back_end_svc/adapter/repository"
	"back_end_svc/usecase/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type Registry struct {
	db *mongo.Database
}

func NewRegistry(db *mongo.Database) Registry {
	return Registry{
		db: db,
	}
}

func (r *Registry) NewTaskController() controller.TaskController {
	db := repository.NewDatabaseRepo(r.db.Collection("user"), r.db.Collection("task"), r.db.Collection("user_info"))
	return controller.NewTaskController(
		*service.NewTaskService(db),
	)
}

func (r *Registry) NewUserController() controller.UserController {
	db := repository.NewDatabaseRepo(r.db.Collection("user"), r.db.Collection("task"), r.db.Collection("user_info"))
	return controller.NewUserController(
		*service.NewUserService(db),
	)
}
