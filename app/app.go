package app

import (
	"challenge-SmartMEI/controller"
	"challenge-SmartMEI/database"
	"challenge-SmartMEI/handlers/addBook"
	"challenge-SmartMEI/handlers/createUser"

	httping "github.com/ednailson/httping-go"
	log "github.com/sirupsen/logrus"
)

type App struct {
	server    httping.IServer
	closeFunc httping.ServerCloseFunc
}

func LoadApp(cfg Config) (*App, error) {
	var app App
	db, err := database.NewDatabase(cfg.Database.Config)
	if err != nil {
		return nil, err
	}
	userColl, err := db.Collection(cfg.Database.UserCollection)
	if err != nil {
		return nil, err
	}
	controller := controller.NewController(userColl)
	app.server = loadServer(controller)
	return &app, nil
}

func (a *App) Run() <-chan error {
	closeFunc, chErr := a.server.RunServer()
	a.closeFunc = closeFunc
	return chErr
}

func (a *App) Close() {
	err := a.closeFunc()
	if err != nil {
		log.WithField("error", err.Error()).Errorf("failed to close func")
	}
}

func loadServer(ctrl *controller.Controller) httping.IServer {
	server := httping.NewHttpServer("", 8082)
	createUserHandler := createUser.NewHandler(*ctrl)
	server.NewRoute(nil, "/createUser").POST(createUserHandler.Handle)
	addBookHandler := addBook.NewHandler(*ctrl)
	server.NewRoute(nil, "/addBook/:userid").POST(addBookHandler.Handle)
	return server
}
