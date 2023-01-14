package fiber

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	_ "github.com/kyzykyky/softwarearch/bookservice/internal/controller/server/fiber/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kyzykyky/softwarearch/bookservice/pkg/domain"
)

type Server struct {
	App     *fiber.App
	Service domain.BookService
	Conf
}
type Conf struct {
	Host string
}

//	@title			Bookservice Fiber API
//	@version		1.0.0
//	@description	Layering lab bookservice
//	@schemes		http
func NewServer(svc domain.BookService, conf Conf) Server {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "Bookservice",
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: RecoverStackTraceHandler,
	}))
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(fiberLogger.New(GetLoggerConfig()))
	app.Get("/docs/*", swagger.New())

	server := Server{
		App:     app,
		Service: svc,
		Conf:    conf,
	}
	server.SetupBookserviceRoutes()

	return server
}

func (app *Server) Listen() error {
	return app.App.Listen(app.Host)
}
