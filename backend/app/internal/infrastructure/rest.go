package infrastructure

import "github.com/gofiber/fiber/v2"

type RestConfig struct {
	Key  string
	Port string
}

type RestServer struct {
	App     *fiber.App
	Config  RestConfig
	Service *AppService
}

func NewRestServer(app *fiber.App) *RestServer {
	return &RestServer{
		App: app,
	}
}

func (server *RestServer) Start() error {
	var err error
	// global middle ware
	if err = initRestMiddleware(server.App, server.Config.Key); err != nil {
		return err
	}

	var port string
	if server.Config.Port != "" {
		port = ":" + server.Config.Port
	} else {
		port = ":9000"
	}
	return server.App.Listen(port)
}
