package repoInterface

import (
	"context"

	commonModel "github.com/luandnh/go-notify/common/model"
	"github.com/luandnh/go-notify/repository/model"
)

type IApplicationRepository interface {
	GetApplications(ctx context.Context, filter commonModel.GeneralFilter) (*[]model.Application, []byte, error)
	FindByApplicationName(ctx context.Context, applicationName string) (*model.Application, error)
	Insert(ctx context.Context, app ...*model.Application) error
	FindByApplicationToken(ctx context.Context, token string) (*model.Application, error)
}
