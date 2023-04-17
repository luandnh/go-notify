package repoInterface

import (
	"context"

	"github.com/luandnh/go-notify/repository/model"
)

type IUserRepository interface {
	Insert(ctx context.Context, user ...*model.User) error
	FindByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error)
	FindByUserToken(ctx context.Context, userToken string) (*model.User, error)
}
