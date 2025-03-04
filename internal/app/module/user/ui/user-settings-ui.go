package user_ui

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	components "github.com/thebranchcrafter/go-htmx-starter/components/dashboard"
	user_domain "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/domain"
	"log"
	"net/http"
)

func HandleUserSettings(g *gin.Context) {
	u, exists := g.Get("user")
	if !exists {
		log.Println("UserInfo not found in context")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "User info not found"})
		return
	}

	_, ok := u.(*user_domain.User)
	if !ok {
		panic("user not found")
	}

	isHTMX := g.GetHeader("HX-Request") == "true"
	g.Writer.Header().Set("Content-Type", "text/html")
	if isHTMX {
		templ.Handler(components.UserSettings()).ServeHTTP(g.Writer, g.Request)
	} else {
		templ.Handler(components.UserSettingsLayout()).ServeHTTP(g.Writer, g.Request)
	}
}
