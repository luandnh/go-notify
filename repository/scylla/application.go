package scylla

import (
	"context"

	"github.com/gocql/gocql"
	commonModel "github.com/luandnh/go-notify/common/model"
	repoInterface "github.com/luandnh/go-notify/repository/interface"
	"github.com/luandnh/go-notify/repository/model"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type (
	ApplicationRepository struct {
		tableName string
		table     *table.Table
	}
)

func NewApplicationRepo() repoInterface.IApplicationRepository {
	tblMeta := table.Metadata{
		Name: "application",
		Columns: []string{
			"application_id",
			"application_name",
			"application_token",
			"description",
			"is_active",
			"is_deleted",
			"created_at",
			"updated_at",
			"expired_at",
		},
		PartKey: []string{"application_id"},
		SortKey: []string{},
	}
	tbl := table.New(tblMeta)
	repo := &ApplicationRepository{
		tableName: "go_notify.application",
		table:     tbl,
	}
	if err := repo.createTable(); err != nil {
		panic(err)
	}
	return repo
}

func (repo *ApplicationRepository) createTable() error {
	err := RepoClient.GetSession().ExecStmt(`CREATE TABLE IF NOT EXISTS application (
		application_id uuid PRIMARY KEY,
		application_name text,
		application_token text,
		description text,
		is_active boolean,
		is_deleted boolean,
		created_at timestamp,
		updated_at timestamp,
		expired_at timestamp
	)`)
	return err
}

func (repo *ApplicationRepository) GetApplications(ctx context.Context, filter commonModel.GeneralFilter) (*[]model.Application, []byte, error) {
	applications := new([]model.Application)
	q := qb.Select(repo.tableName).Query(*RepoClient.GetSession())
	iter := q.PageSize(filter.PageSize).PageState(filter.Page).Iter()
	return applications, iter.PageState(), iter.Select(applications)
}

func (repo *ApplicationRepository) FindByApplicationName(ctx context.Context, applicationName string) (*model.Application, error) {
	application := new(model.Application)
	q := qb.Select(repo.tableName).Columns("*").Where(qb.Eq("application_name")).
		Limit(1).
		AllowFiltering().
		Query(*RepoClient.GetSession()).
		BindMap(qb.M{
			"application_name": applicationName,
		})
	qa := gocqlx.Query(q.Query, nil)
	err := qa.GetRelease(application)
	if err == gocql.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return application, nil
}

func (repo *ApplicationRepository) Insert(ctx context.Context, app ...*model.Application) error {
	q := qb.Insert(repo.tableName).Columns(
		"application_id",
		"application_name",
		"application_token",
		"description",
		"is_active",
		"is_deleted",
		"created_at",
		"updated_at",
		"expired_at",
	).Query(*RepoClient.GetSession()).WithContext(ctx)
	for _, u := range app {
		if err := q.BindStruct(u).Exec(); err != nil {
			return err
		}
	}
	q.Release()
	return nil
}

func (repo *ApplicationRepository) FindByApplicationToken(ctx context.Context, token string) (*model.Application, error) {
	app := new(model.Application)
	q := qb.Select(repo.tableName).Columns("*").Where(qb.Eq("application_token")).
		Limit(1).
		AllowFiltering().
		Query(*RepoClient.GetSession()).
		BindMap(qb.M{
			"application_token": token,
		})
	qa := gocqlx.Query(q.Query, nil)
	err := qa.GetRelease(app)
	if err != nil {
		return nil, err
	}
	return app, nil
}
