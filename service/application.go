package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/luandnh/go-notify/common/log"
	commonModel "github.com/luandnh/go-notify/common/model"
	"github.com/luandnh/go-notify/common/response"
	"github.com/luandnh/go-notify/common/util"
	"github.com/luandnh/go-notify/repository"
	"github.com/luandnh/go-notify/repository/model"
)

type (
	IApplication interface {
		GetApplications(ctx context.Context, filter commonModel.GeneralFilter) (int, interface{})
		PostApplication(ctx context.Context, body map[string]any) (int, interface{})
		FindByApplicationToken(ctx context.Context, applicationToken string) (*model.Application, error)
	}
	Application struct{}
)

func NewApplication() IApplication {
	return &Application{}
}

func (s *Application) GetApplications(ctx context.Context, filter commonModel.GeneralFilter) (int, interface{}) {
	applications, nextPage, err := repository.ApplicationRepo.GetApplications(ctx, filter)
	if err != nil {
		log.Error(err)
		return response.Paging(response.EmptyData(), nil, nil)
	}
	if len(nextPage) == 0 {
		return response.Paging(applications, nil, nil)
	}
	return response.Paging(applications, nil, nextPage)
}

func (s *Application) PostApplication(ctx context.Context, body map[string]any) (int, interface{}) {
	appName, _ := body["application_name"].(string)
	applicationExisted, err := repository.ApplicationRepo.FindByApplicationName(ctx, appName)
	if err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("Failed to create application")
	} else if applicationExisted != nil {
		return response.BadRequestMsg("Application is existed")
	}
	app := model.Application{}
	app.ApplicationId = uuid.NewString()
	app.ApplicationName = appName
	app.ApplicationToken = util.GenRandomString(32)
	app.Description, _ = body["description"].(string)
	app.IsActive = true
	app.IsDeleted = false
	app.CreatedAt = time.Now()
	app.UpdatedAt = time.Now()
	// TODO: implement expired
	if err := repository.ApplicationRepo.Insert(ctx, &app); err != nil {
		log.Error(err)
		return response.ServiceUnavailableMsg("Failed to create application")
	}
	return response.Created(map[string]any{
		"application_id":    app.ApplicationId,
		"application_name":  app.ApplicationName,
		"application_token": app.ApplicationToken,
	})
}

func (s *Application) FindByApplicationToken(ctx context.Context, applicationToken string) (*model.Application, error) {
	return repository.ApplicationRepo.FindByApplicationToken(ctx, applicationToken)
}
