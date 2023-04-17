package scylla

import (
	"context"

	repoInterface "github.com/luandnh/go-notify/repository/interface"
	"github.com/luandnh/go-notify/repository/model"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type (
	ClientRepository struct {
		table     *table.Table
		tableName string
	}
)

func NewClientRepo() repoInterface.IClientRepository {
	tblMeta := table.Metadata{
		Name: "client",
		Columns: []string{
			"application_id",
			"user_id",
			"client_id",
			"name",
			"client_token",
			"extras",
			"is_active",
			"is_deleted",
			"created_at",
			"updated_at",
			"expired_at",
		},
		PartKey: []string{"client_id", "user_id"},
		SortKey: []string{},
	}
	tbl := table.New(tblMeta)
	repo := &ClientRepository{
		table:     tbl,
		tableName: "go_notify.client",
	}
	if err := repo.createTable(); err != nil {
		panic(err)
	}
	return repo
}

func (repo *ClientRepository) createTable() error {
	err := RepoClient.GetSession().ExecStmt(`CREATE TABLE IF NOT EXISTS client (
		application_id uuid,
		user_id uuid,
		client_id uuid PRIMARY KEY,
		name text,
		client_token text,
		extras extra_value,
		is_active boolean,
		is_deleted boolean,
		created_at timestamp,
		updated_at timestamp,
		expired_at timestamp
	)`)
	return err
}

func (repo *ClientRepository) GetClientByClientToken(ctx context.Context, clientToken string) (*model.Client, error) {
	client := new(model.Client)
	q := qb.Select(repo.tableName).Where(qb.EqLit("client_token", clientToken)).Query(*RepoClient.GetSession())
	err := q.GetRelease(client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (repo *ClientRepository) GetClientByUserId(ctx context.Context, userId string) (*model.Client, error) {
	client := new(model.Client)
	q := qb.Select(repo.tableName).Where(qb.EqLit("user_id", userId)).Query(*RepoClient.GetSession())
	err := q.GetRelease(client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (repo *ClientRepository) GetClientByClientId(ctx context.Context, clientId string) (*model.Client, error) {
	client := new(model.Client)
	q := qb.Select(repo.tableName).Where(qb.EqLit("client_id", clientId)).Query(*RepoClient.GetSession())
	err := q.GetRelease(client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (repo *ClientRepository) InsertClient(ctx context.Context, client *model.Client) error {
	q := qb.Insert(repo.tableName).Columns(
		"application_id",
		"user_id",
		"client_id",
		"name",
		"client_token",
		"extras",
		"is_active",
		"is_deleted",
		"created_at",
		"updated_at",
		"expired_at",
	).Query(*RepoClient.GetSession()).BindStruct(client).WithContext(ctx)
	if err := q.ExecRelease(); err != nil {
		return err
	}
	return nil
}
