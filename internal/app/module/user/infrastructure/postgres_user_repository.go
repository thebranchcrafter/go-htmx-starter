package user_infrastructure

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	user_domain "github.com/thebranchcrafter/go-htmx-starter/internal/app/module/user/domain"
	"time"
)

var columns = []string{
	"id",
	"username",
	"email",
	"hashed_password",
	"name",
	"surname",
	"role",
	"profile_picture_url",
	"created_at",
	"updated_at",
}

// PostgresUserRepository is a Postgres implementation of UserRepository using Squirrel and pgxpool.
type PostgresUserRepository struct {
	tableName string
	DB        *pgxpool.Pool
}

// NewPostgresUserRepository initializes a new Postgres User repository with a connection pool.
func NewPostgresUserRepository(poolConfig, tableName string) (*PostgresUserRepository, error) {
	pool, err := pgxpool.Connect(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return &PostgresUserRepository{
		DB:        pool,
		tableName: tableName,
	}, nil
}

// Save creates or updates a User record.
func (r *PostgresUserRepository) Save(ctx context.Context, u *user_domain.User) error {
	id, err := uuid.Parse(string(u.Id()))
	if err != nil {
		return err
	}

	query := squirrel.Insert(r.tableName).
		Columns(columns...).
		Values(
			id,
			u.Username(),
			u.Email(),
			u.HashedPassword(),
			u.Name(),
			u.Surname(),
			u.Role(),
			u.ProfilePictureUrl(),
			u.CreatedAt(),
			u.UpdatedAt(),
		).
		Suffix("ON CONFLICT (id) DO UPDATE SET updated_at = EXCLUDED.updated_at").
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(ctx, sqlQuery, args...)
	if err != nil {
		// Check if the error is a unique constraint violation
		var pgErr *pgconn.PgError
		errors.As(err, &pgErr)
		if pgErr.Code == "23505" {
			return user_domain.NewUserAlreadyExists(u.Email())
		}
		return err
	}
	return nil
}

// GetByID retrieves a User by its ID.
func (r *PostgresUserRepository) GetByID(ctx context.Context, id user_domain.UserId) (*user_domain.User, error) {
	queryBuilder := squirrel.Select(columns...).
		From(r.tableName).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.DB.QueryRow(ctx, sqlQuery, args...)
	return scanRow(row)
}

// GetByEmail retrieves a User by its ID.
func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*user_domain.User, error) {
	queryBuilder := squirrel.Select(columns...).
		From(r.tableName).
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	row := r.DB.QueryRow(ctx, sqlQuery, args...)
	return scanRow(row)
}

// GetAll retrieves all User records, with optional filters.
func (r *PostgresUserRepository) GetAll(ctx context.Context, filters map[string]interface{}) ([]*user_domain.User, error) {
	query := squirrel.Select(columns...).
		From(r.tableName).
		PlaceholderFormat(squirrel.Dollar)

	for key, value := range filters {
		query = query.Where(squirrel.Eq{key: value})
	}

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.DB.Query(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*user_domain.User
	for rows.Next() {
		user, err := scanRow(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

// Update modifies an existing User record.
func (r *PostgresUserRepository) Update(ctx context.Context, u *user_domain.User) error {
	query := squirrel.Update(r.tableName).
		Set("updated_at", u.UpdatedAt()).
		Where(squirrel.Eq{"id": u.Id()}).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(ctx, sqlQuery, args...)
	return err
}

// Delete removes a User by its ID.
func (r *PostgresUserRepository) Delete(ctx context.Context, id user_domain.UserId) error {
	query := squirrel.Delete(r.tableName).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(ctx, sqlQuery, args...)
	return err
}

// scanRow scans a row into a User domain object.
func scanRow(row squirrel.RowScanner) (*user_domain.User, error) {
	var (
		id                 string
		username           string
		email              string
		hashedPwd          string
		name               string
		surname            string
		role               string
		profile_pictureUrl string
		createdAt          time.Time
		updatedAt          time.Time
	)

	err := row.Scan(
		&id,
		&username,
		&email,
		&hashedPwd,
		&name,
		&surname,
		&role,
		&profile_pictureUrl,
		&createdAt,
		&updatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user_domain.FromPrimitives(
		id,
		username,
		email,
		hashedPwd,
		name,
		surname,
		role,
		profile_pictureUrl,
		createdAt,
		updatedAt,
	), nil
}
