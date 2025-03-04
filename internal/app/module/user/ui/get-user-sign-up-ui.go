package user_ui

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	components "github.com/thebranchcrafter/go-htmx-starter/components/landing"
)

func HandleUserSignUpIndex(g *gin.Context) {
	templ.Handler(components.SignUp()).ServeHTTP(g.Writer, g.Request)
}
