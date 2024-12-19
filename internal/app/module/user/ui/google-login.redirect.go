package user_ui

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
)

type HandleGoogleRedirect struct {
	googleOauthConfig *oauth2.Config
}

func NewHandleGoogleRedirect(googleOauthConfig *oauth2.Config) *HandleGoogleRedirect {
	return &HandleGoogleRedirect{googleOauthConfig: googleOauthConfig}
}

func (h *HandleGoogleRedirect) HandleGoogleLoginRedirect(g *gin.Context) {
	url := h.googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	g.Redirect(http.StatusFound, url)

}
