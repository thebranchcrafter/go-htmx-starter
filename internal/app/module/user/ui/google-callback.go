package user_ui

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mik3lon/starter-template/pkg/router"
	user_application "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/application"
	user_domain "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/domain"
	"github.com/thebranchcrafter/go-kit/pkg/bus/command"
	"github.com/thebranchcrafter/go-kit/pkg/bus/query"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

type GoogleCallbackHandler struct {
	cb                command.Bus
	qb                query.Bus
	googleOauthConfig *oauth2.Config
}

func NewGoogleCallbackHandler(cb command.Bus, qb query.Bus, googleOauthConfig *oauth2.Config) *GoogleCallbackHandler {
	return &GoogleCallbackHandler{cb: cb, qb: qb, googleOauthConfig: googleOauthConfig}
}

func (h *GoogleCallbackHandler) HandleGoogleCallback(g *gin.Context) {
	code := g.Query("code")
	if code == "" {
		http.Error(g.Writer, `{"error": "Code not found"}`, http.StatusBadRequest)
		return
	}

	// Exchange the authorization code for an access token
	token, err := h.googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Failed to exchange token: %v\n", err)
		http.Error(g.Writer, `{"error": "Failed to exchange token"}`, http.StatusInternalServerError)
		return
	}

	// Use the access token to get user info
	client := h.googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Failed to get user info: %v\n", err)
		http.Error(g.Writer, `{"error": "Failed to get user info"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Display user info as JSON (for demonstration purposes)
	var userInfo router.GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Printf("Failed to parse user info: %v\n", err)
		http.Error(g.Writer, `{"error": "Failed to parse user info"}`, http.StatusInternalServerError)
		return
	}

	cuc := &user_application.CreateUserCommand{
		ID:                uuid.NewString(),
		Name:              userInfo.GivenName,
		Surname:           userInfo.FamilyName,
		Username:          userInfo.Name,
		PlainPassword:     "",
		Email:             userInfo.Email,
		Role:              "ROL_USER",
		ProfilePictureUrl: userInfo.Picture,
		IsFormSocialAuth:  true,
	}

	err = h.cb.Dispatch(g, cuc)
	var userAlreadyExists *user_domain.UserAlreadyExists
	switch {
	case err == nil, errors.As(err, &userAlreadyExists):
		user, err := h.qb.Ask(g, &user_application.FindUserQuery{Email: userInfo.Email})
		if err != nil {
			panic(err)
		}

		createdUser := user.(*user_domain.User)
		session := sessions.Default(g)
		session.Set(
			"user",
			user_domain.LoggedUser{
				ID:    string(createdUser.Id()),
				Email: createdUser.Email(),
			},
		)
		if err := session.Save(); err != nil {
			log.Printf("Failed to save session: %v\n", err)
			http.Error(g.Writer, `{"error": "Failed to save session"}`, http.StatusInternalServerError)
			return
		}

		http.Redirect(g.Writer, g.Request, "/dashboard", http.StatusFound)
	default:
		http.Error(g.Writer, `{"error": "Failed to get user info"}`, http.StatusInternalServerError)
	}
}
