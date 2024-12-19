package user_ui

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	user_application "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/application"
	user_domain "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/domain"
	"github.com/thebranchcrafter/go-kit/pkg/bus/command"
	"github.com/thebranchcrafter/go-kit/pkg/bus/query"
	"net/http"
)

type UserSignUpHandler struct {
	cb command.Bus
	qb query.Bus
}

func NewUserSignUpHandler(cb command.Bus, qb query.Bus) *UserSignUpHandler {
	return &UserSignUpHandler{cb: cb, qb: qb}
}

func (h *UserSignUpHandler) HandleUserSignUp(g *gin.Context) {
	id := uuid.NewString()
	username := g.PostForm("username")
	plainPassword := g.PostForm("password")
	email := g.PostForm("email")

	err := h.cb.Dispatch(g, &user_application.CreateUserCommand{
		ID:                id,
		Name:              "",
		Surname:           "",
		Username:          username,
		PlainPassword:     plainPassword,
		Email:             email,
		Role:              "ROL_USER",
		ProfilePictureUrl: "",
		IsFormSocialAuth:  false,
	})
	switch err.(type) {
	case nil:
		g.Writer.Header().Set("HX-Trigger", `{"userSignUpSuccess": "Your account has been created. Please check your email to activate your account."}`)
		g.Writer.WriteHeader(http.StatusNoContent)
		return
	case *user_domain.UserAlreadyExists:
		g.Writer.Header().Set("Content-Type", "application/json")
		g.Writer.Header().Set("HX-Trigger", `{"userSignUpError": "Unable to complete the sign-up process. Please try again later."}`)
		g.Writer.WriteHeader(http.StatusBadRequest)
		g.Writer.Write([]byte(`{"error": "Unable to complete the sign-up process. Please try again later."}`))
		return
	default:
		http.Error(g.Writer, `{"error": "Failed to get user info"}`, http.StatusInternalServerError)
	}
}
