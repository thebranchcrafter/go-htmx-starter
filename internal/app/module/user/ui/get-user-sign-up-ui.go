package user_ui

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/ui/views"
)

func HandleUserSignUpIndex(g *gin.Context) {
	templ.Handler(views.SignUp()).ServeHTTP(g.Writer, g.Request)
}
