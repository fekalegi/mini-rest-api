package api

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"mini-rest-api/api/controller"
	"mini-rest-api/api/handler"
	"mini-rest-api/common/helper"
	external "mini-rest-api/infra"
	"mini-rest-api/repository/postgres"
	"mini-rest-api/usecase/note"
	"mini-rest-api/usecase/user"
	"net/http"
	"time"
)

// Server Struct
type Server struct {
	Route *echo.Echo
}

// NewServer : Create Server Instance
func NewServer(e *echo.Echo) *Server {
	return &Server{
		e,
	}
}

func (server *Server) InitializeServer() {
	newDB := external.NewGormDB()

	helperInterface := helper.NewHelperFunction()
	repoUser := postgres.NewUserRepository(newDB)
	userUC := user.NewUserImplementation(repoUser, helperInterface)
	userController := controller.NewUserController(userUC)

	repoNote := postgres.NewNoteRepository(newDB)
	noteUC := note.NewNoteImplementation(repoNote)
	noteController := controller.NewNoteController(noteUC)

	apiGroup := server.Route.Group("/api")

	userController.Route(apiGroup)
	noteController.Route(apiGroup)

	server.Route.GET("/swagger/*", echoSwagger.WrapHandler)

	handler.UseCustomValidatorHandler(server.Route)

	serverConfiguration := &http.Server{
		Addr:         ":5000",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	server.Route.Logger.Fatal(server.Route.StartServer(serverConfiguration))

}
