package app

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	user_domain "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/domain"
	"github.com/thebranchcrafter/go-htmx-starter/pkg/infrastructure"
	"log"
	"path/filepath"

	"github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user"
	"github.com/thebranchcrafter/go-kit/pkg/bus/command"
	"github.com/thebranchcrafter/go-kit/pkg/bus/query"
	"github.com/thebranchcrafter/go-kit/pkg/infrastructure/app"
	http_response "github.com/thebranchcrafter/go-kit/pkg/infrastructure/http/response"
	"github.com/thebranchcrafter/go-kit/pkg/infrastructure/logger"
)

func Bootstrap() (*app.Kernel, error) {
	// Initialize kernel with configuration
	gob.Register(user_domain.LoggedUser{})
	engine := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	engine.Use(sessions.Sessions("session", store))

	absPath, err := filepath.Abs("./assets")
	engine.Static("/assets", absPath)

	l := logger.NewZerologAdapter()

	eventBus, err := infrastructure.NewNATSEventBus("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Error initializing NATS EventBus: %v\n", err)
	}

	k := app.NewKernel(
		app.WithRouter(engine),
		app.WithEventBus(eventBus),
		app.WithLogger(l),
		app.WithCommandBus(command.InitCommandBus(l)),
		app.WithQueryBus(query.InitQueryBus(l)),
		app.WithJsonResponseWriter(http_response.NewJsonResponseWriter()),
	)

	um := user.InitUserModule(k.CommonDependencies)

	err = k.AddModule(um)
	if err != nil {
		return nil, err
	}

	k.RegisterRoutes()

	return k, nil
}
