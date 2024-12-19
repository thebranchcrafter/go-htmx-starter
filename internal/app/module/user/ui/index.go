package user_ui

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/ui/views"
)

func HandleIndex(g *gin.Context) {
	templ.Handler(views.Index()).ServeHTTP(g.Writer, g.Request)
}
