package service

import (
	"context"
	"time"

	"github.com/luandnh/go-notify/repository"
	"github.com/luandnh/go-notify/repository/model"
)

type (
	IUser interface {
		FindByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error)
		FindByUserToken(ctx context.Context, userToken string) (*model.User, error)
	}
	User struct {
		DefaultAdminPassword string
		DefaultAdminToken    string
	}
)

func NewUser(defaultAdminPassword, defaultAdminToken string) IUser {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := repository.UserRepo.Insert(ctx, &model.User{
		ApplicationId: "00000000-0000-0000-0000-000000000000",
		UserId:        "00000000-0000-0000-0000-000000000000",
		RefUserId:     "00000000-0000-0000-0000-000000000000",
		Username:      "admin",
		Password:      defaultAdminPassword,
		UserToken:     defaultAdminToken,
		Level:         ADMIN,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}); err != nil {
		panic(err)
	}
	return &User{
		DefaultAdminPassword: defaultAdminPassword,
		DefaultAdminToken:    defaultAdminToken,
	}
}

func (s *User) FindByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error) {
	return repository.UserRepo.FindByUsernameAndPassword(ctx, username, password)
}

func (s *User) FindByUserToken(ctx context.Context, userToken string) (*model.User, error) {
	return repository.UserRepo.FindByUserToken(ctx, userToken)
}
