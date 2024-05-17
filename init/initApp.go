package init

import (
	"back_end_svc/adapter/connection"
	"back_end_svc/registry"

	"github.com/labstack/echo/v4"
)

func InitApp() {
	mon := connection.Connn()
	r := registry.NewRegistry(mon)
	tc := r.NewTaskController()
	uc := r.NewUserController()
	e := echo.New()
	NewRouter(e, tc, uc)
	//go http.ListenAndServe(":8081", nil)
	go e.Logger.Fatal(e.Start(":8080"))
}
