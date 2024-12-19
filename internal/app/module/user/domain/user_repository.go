package user_domain

import "context"

type UserRepository interface {
	// Save a new User
	Save(ctx context.Context, u *User) error

	// GetByID a User by its ID
	GetByID(ctx context.Context, id UserId) (*User, error)

	// GetByEmail a User by its ID
	GetByEmail(ctx context.Context, email string) (*User, error)

	// GetAll User with optional filters (if necessary)
	GetAll(ctx context.Context, filters map[string]interface{}) ([]*User, error)

	// Update an existing User
	Update(ctx context.Context, u *User) error

	// Delete a User by its ID
	Delete(ctx context.Context, id UserId) error
}
