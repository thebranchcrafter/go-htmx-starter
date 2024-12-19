package user_domain

import (
	"github.com/thebranchcrafter/go-kit/pkg/bus/event"
	"time"
)

type UserId string

type User struct {
	id                UserId
	username          string
	email             string
	hashedPassword    string
	name              string
	surname           string
	role              string
	profilePictureUrl string
	createdAt         time.Time
	updatedAt         time.Time

	events []event.Event
}

type LoggedUser struct {
	ID    string
	Email string
}

func (t *User) Id() UserId {
	return t.id
}

func (t *User) Username() string {
	return t.username
}

func (t *User) Email() string {
	return t.email
}

func (t *User) HashedPassword() string {
	return t.hashedPassword
}

func (t *User) Name() string {
	return t.name
}

func (t *User) Surname() string {
	return t.surname
}

func (t *User) Role() string {
	return t.role
}

func (t *User) ProfilePictureUrl() string {
	return t.profilePictureUrl
}

func (t *User) CreatedAt() time.Time {
	return t.createdAt
}

func (t *User) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *User) AddEvent(event event.Event) {
	t.events = append(t.events, event)
}

func (t *User) Events() []event.Event {
	return t.events
}

func CreateUser(id, username, email, password, name, surname, role, profilePictureUrl string) *User {
	user := &User{
		id:                UserId(id),
		username:          username,
		email:             email,
		hashedPassword:    password,
		name:              name,
		surname:           surname,
		role:              role,
		profilePictureUrl: profilePictureUrl,
		createdAt:         time.Now(),
		updatedAt:         time.Now(),
	}
	e := NewUserCreatedEvent(
		user.id,
		user.username,
		user.email,
		user.name,
		user.surname,
		user.role,
		user.profilePictureUrl,
		time.Now(),
	)
	user.AddEvent(e)
	return user
}

func FromPrimitives(
	id string,
	username string,
	email string,
	hashedPassword string,
	name string,
	surname string,
	role string,
	profilePictureUrl string,
	createdAt, updatedAt time.Time,
) *User {
	return &User{
		id:                UserId(id),
		username:          username,
		email:             email,
		hashedPassword:    hashedPassword,
		name:              name,
		surname:           surname,
		role:              role,
		profilePictureUrl: profilePictureUrl,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}
}

// UserCreatedEvent represents an immutable domain event
type UserCreatedEvent struct {
	ID                UserId    `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	Surname           string    `json:"surname"`
	Role              string    `json:"role"`
	ProfilePictureUrl string    `json:"profile_picture_url"`
	OccurredAt        time.Time `json:"occurred_at"`
}

func (e UserCreatedEvent) EventName() string {
	return "user_created"
}

// NewUserCreatedEvent creates an immutable UserCreatedEvent
func NewUserCreatedEvent(
	id UserId,
	username, email, name, surname, role, profilePictureUrl string,
	occurredAt time.Time,
) UserCreatedEvent {
	return UserCreatedEvent{
		ID:                id,
		Username:          username,
		Email:             email,
		Name:              name,
		Surname:           surname,
		Role:              role,
		ProfilePictureUrl: profilePictureUrl,
		OccurredAt:        occurredAt,
	}
}
