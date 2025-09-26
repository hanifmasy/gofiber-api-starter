package observability

import (
	"log"
	"os"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/getsentry/sentry-go"
	sentryfiber "github.com/getsentry/sentry-go/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

// InitSentry initializes Sentry with the provided DSN and environment
func InitSentry(dsn, environment string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         dsn,
		Environment: environment,
		// Set TracesSampleRate to 1.0 to capture 100% of the transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Printf("Sentry initialization failed: %v", err)
	}
}

func Setup(app *fiber.App) {
	// Request ID
	app.Use(requestid.New())

	// Logger
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	// Recovery
	app.Use(recover.New())

	// Prometheus
	prom := fiberprometheus.New("golang_fiber_api")
	prom.RegisterAt(app, "/metrics")
	app.Use(prom.Middleware)

	// Sentry
	InitSentry(os.Getenv("SENTRY_DSN"), os.Getenv("APP_ENV"))
	app.Use(sentryfiber.New(sentryfiber.Options{}))
}
