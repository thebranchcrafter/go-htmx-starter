package user_ui

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	components "github.com/thebranchcrafter/go-htmx-starter/components/dashboard"
	user_domain "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/domain"
	"log"
	"net/http"
)

func HandleUserPasswordSign(g *gin.Context) {
	u, exists := g.Get("user")
	if !exists {
		log.Println("UserInfo not found in context")
		g.JSON(http.StatusInternalServerError, gin.H{"error": "User info not found"})
		return
	}

	user, ok := u.(*user_domain.User)
	if !ok {
		panic("user not found")
	}

	g.Writer.Header().Set("Content-Type", "text/html")
	log.Println(user.Name())
	templ.Handler(components.DashboardLayout(user)).ServeHTTP(g.Writer, g.Request)
}
