package auth

import (
	"context"
	"log/slog"
	"sso/internal/domain/models"
	"time"
)

type Auth struct {
	userSaver    UserSaver
	userProvider AppProvider
	log          *slog.Logger
	appProvider  AppProvider
	tokenTTL     time.Duration
}
type UserSaver interface {
	SaveUser(ctx context.Context, email string, passHash []byte) (uid int64, err error)
}
type UserProvider interface {
	User(ctx context.Context, email string) (models.User, error)
	IsAdmin(ctx context.Context, userID int64) (bool, error)
}
type AppProvider interface {
	App(ctx context.Context, appID int) (models.App, error)
}

// New returns a new instance of the Auth service.
func New(log *slog.Logger, userSaver UserSaver, userPovider UserProvider, appProvider AppProvider, tokenTTL time.Duration) *Auth {
	return &Auth{
		userSaver:    userSaver,
		userProvider: appProvider,
		log:          log,
		appProvider:  appProvider,
		tokenTTL:     tokenTTL,
	}
}

// Login checks if user with given credentials exist in the system
//
// If user exist, but password is incorrect,return error.
// If user doesn't exist , return error.
func (a *Auth) Login(ctx context.Context, email string, password string, appID int) (string, error) {
	panic("not implemented")
}

// RefisterNewUser reqisters new user in the system and returns user ID.
// If user with given username already exists, return error.
func (a *Auth) RegisterNewUser(ctx context.Context, email string, pass string) (int64, error) {
	panic("not implemented")
}

// IsAdmin checks if user is admin
func (a *Auth) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	panic("not implemented")
}
