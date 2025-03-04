package user_ui

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/thebranchcrafter/go-htmx-starter/components/landing"
)

func HandleUserSignInIndex(g *gin.Context) {
	props := landing.SignInProps{
		GoogleAuthURL: templ.SafeURL("/auth/google"),
	}

	templ.Handler(landing.SignIn(props)).ServeHTTP(g.Writer, g.Request)
}
