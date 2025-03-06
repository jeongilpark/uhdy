package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sagepulse.ai/uhdy/user-service/db"
)

var (
	ErrNoRecord = errors.New("no matching record found")
)

func ConnectDatabase(user string, pw string, host string, port int, dbname string) *pgxpool.Pool {
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, pw, host, port, dbname)
	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalf("Unable to parse database url: %v\n", err)
	}
	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect database: %v\n", err)
	}
	return dbpool
}

type UserRepository interface {
	CreateUser(ctx context.Context, username string, password string) error
	GetUser(ctx context.Context, username string) (db.User, error)
}

type UserPostgresRepository struct {
	queries *db.Queries
}

func NewUserPostgresRepository(user string, pw string, host string, port int, dbname string) UserRepository {
	dbpool := ConnectDatabase(user, pw, host, port, dbname)
	return &UserPostgresRepository{
		queries: db.New(dbpool),
	}
}

func (r *UserPostgresRepository) CreateUser(ctx context.Context, username string, password string) error {
	_, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Username:     username,
		PasswordHash: password,
	})
	return err
}

func (r *UserPostgresRepository) GetUser(ctx context.Context, username string) (db.User, error) {
	user, err := r.queries.GetUserByUsername(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, ErrNoRecord
		}
		return user, err
	}
	return user, nil
}
