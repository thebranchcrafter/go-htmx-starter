package user_ui

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleLogout(g *gin.Context) {
	// Clear the user session or authentication cookies
	session, err := g.Cookie("session") // Assuming session is stored in a cookie named "session"
	if err == nil && session != "" {
		// Invalidate the session
		g.SetCookie("session", "", -1, "/", "", false, true)
	}

	// Redirect to the home page
	g.Redirect(http.StatusFound, "/")
}
