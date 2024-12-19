package user_application

import (
	"context"
	"errors"
	user_domain "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/domain"
	"github.com/thebranchcrafter/go-kit/pkg/bus"
	"github.com/thebranchcrafter/go-kit/pkg/bus/event"
)

type CreateUserCommand struct {
	ID                string
	Name              string
	Surname           string
	Username          string
	PlainPassword     string
	Email             string
	Role              string
	ProfilePictureUrl string
	IsFormSocialAuth  bool
}

func (c CreateUserCommand) Id() string {
	return "create-user-command"
}

type CreateUserCommandHandler struct {
	r  user_domain.UserRepository
	eb event.EventBus
}

func NewCreateUserCommandHandler(r user_domain.UserRepository, eb event.EventBus) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{r: r, eb: eb}
}

func (cuch CreateUserCommandHandler) Handle(ctx context.Context, c bus.Command) error {
	cuc, ok := c.(*CreateUserCommand)
	if !ok {
		return errors.New("invalid command")
	}

	password, err := user_domain.GenerateHashedPassword(cuc.IsFormSocialAuth, cuc.PlainPassword)
	if err != nil {
		return errors.New("failed to generate hashed password")
	}

	user := user_domain.CreateUser(
		cuc.ID,
		cuc.Username,
		cuc.Email,
		password,
		cuc.Name,
		cuc.Surname,
		cuc.Role,
		cuc.ProfilePictureUrl,
	)

	if err := cuch.r.Save(ctx, user); err != nil {
		return err
	}

	// Publish events recorded in the User
	for _, event := range user.Events() {
		if err := cuch.eb.Publish(ctx, event); err != nil {
			return err
		}
	}

	return nil
}
