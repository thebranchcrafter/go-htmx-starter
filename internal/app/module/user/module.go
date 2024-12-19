package user

import (
	user_application "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/application"
	user_domain "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/domain"
	user_infrastructure "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/infrastructure"
	user_ui "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/ui"
	"github.com/thebranchcrafter/go-htmx-starter/pkg/http/gin/middleware"
	"github.com/thebranchcrafter/go-kit/pkg/infrastructure/app"
	"github.com/thebranchcrafter/go-kit/pkg/infrastructure/router"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"net/http"
	"os"
)

type UserModule struct {
	*app.BaseModule
	UserRepository        user_domain.UserRepository
	GoogleRedirectHandler *user_ui.HandleGoogleRedirect
	GoogleCallbackHandler *user_ui.GoogleCallbackHandler
}

func (u *UserModule) Routes() []app.Route {
	authMiddleware := middleware.NewAuthSessionMiddleware(u.QueryBus)

	signUpHandler := user_ui.NewUserSignUpHandler(u.CommandBus, u.QueryBus)
	return []app.Route{
		{
			Method:      http.MethodGet,
			Path:        "/sign-in",
			Handler:     user_ui.HandleUserSignInIndex,
			Middlewares: nil,
		},
		{
			Method:      http.MethodPost,
			Path:        "/sign-in",
			Handler:     user_ui.HandleUserPasswordSign,
			Middlewares: nil,
		},
		{
			Method:      http.MethodPost,
			Path:        "/sign-up",
			Handler:     signUpHandler.HandleUserSignUp,
			Middlewares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "sign-up",
			Handler:     user_ui.HandleUserSignUpIndex,
			Middlewares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/auth/google",
			Handler:     u.GoogleRedirectHandler.HandleGoogleLoginRedirect,
			Middlewares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/auth/google/callback",
			Handler:     u.GoogleCallbackHandler.HandleGoogleCallback,
			Middlewares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/dashboard",
			Handler:     user_ui.HandleUserDashboard,
			Middlewares: []router.Middleware{authMiddleware.AuthMiddleware()},
		},
		{
			Method:      http.MethodGet,
			Path:        "/logout",
			Handler:     user_ui.HandleLogout,
			Middlewares: []router.Middleware{authMiddleware.AuthMiddleware()},
		},
		{
			Method:      http.MethodGet,
			Path:        "/",
			Handler:     user_ui.HandleIndex,
			Middlewares: []router.Middleware{},
		},
	}
}

func InitUserModule(d app.CommonDependencies) *UserModule {
	poolConfig := "postgres://myuser:mypassword@localhost:5432/mydb"
	ur, err := user_infrastructure.NewPostgresUserRepository(poolConfig, "users")
	if err != nil {
		log.Fatalf("[InitUserModule] create user repository fail: %v", err)
	}

	var googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_CLIENT_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	um := &UserModule{
		BaseModule:     &app.BaseModule{CommonDependencies: d},
		UserRepository: ur,

		GoogleRedirectHandler: user_ui.NewHandleGoogleRedirect(googleOauthConfig),
		GoogleCallbackHandler: user_ui.NewGoogleCallbackHandler(d.CommandBus, d.QueryBus, googleOauthConfig),
	}

	um.AddCommand(&user_application.CreateUserCommand{}, user_application.NewCreateUserCommandHandler(ur, d.EventBus))
	um.AddQuery(&user_application.FindUserQuery{}, user_application.NewFindUserQueryHandler(ur))

	return um
}

func (u *UserModule) Name() string {
	return "user_module"
}
