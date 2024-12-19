package middleware

import (
	user_application "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/application"
	user_domain "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/domain"
	"github.com/thebranchcrafter/go-kit/pkg/bus/query"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthSessionMiddleware struct {
	qb query.Bus
}

func NewAuthSessionMiddleware(qb query.Bus) *AuthSessionMiddleware {
	return &AuthSessionMiddleware{qb: qb}
}

// AuthMiddleware ensures that the user is authenticated
func (m *AuthSessionMiddleware) AuthMiddleware() func(gin.HandlerFunc) gin.HandlerFunc {
	return func(next gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			session := sessions.Default(c)
			user := session.Get("user")

			if user == nil {
				// If user is not authenticated, respond with an error
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}

			lu, ok := user.(user_domain.LoggedUser)
			if !ok {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid session data"})
				c.Abort()
				return
			}

			u, e := m.qb.Ask(c, &user_application.FindUserQuery{
				Email: lu.Email,
			})
			if e != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": e.Error()})
			}

			// Store user in the context for downstream handlers
			c.Set("user", u)

			// Call the next handler
			next(c)
		}
	}
}
