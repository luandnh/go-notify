package scylla

import (
	"context"

	repoInterface "github.com/luandnh/go-notify/repository/interface"
	"github.com/luandnh/go-notify/repository/model"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type (
	UserRepository struct {
		table     *table.Table
		tableName string
	}
)

func NewUserRepo() repoInterface.IUserRepository {
	tblMeta := table.Metadata{
		Name: "user",
		Columns: []string{
			"application_id",
			"user_id",
			"ref_user_id",
			"username",
			"password",
			"user_token",
			"level",
			"extras",
			"is_active",
			"is_deleted",
			"created_at",
			"updated_at",
			"expired_at",
		},
		PartKey: []string{"user_id"},
		SortKey: []string{},
	}
	tbl := table.New(tblMeta)
	repo := &UserRepository{
		table:     tbl,
		tableName: "go_notify.user",
	}
	if err := repo.createTable(); err != nil {
		panic(err)
	}
	return repo
}

func (repo *UserRepository) createTable() error {
	if err := RepoClient.GetSession().ExecStmt(`CREATE TABLE IF NOT EXISTS user (
		application_id uuid,
		user_id uuid PRIMARY KEY,
		ref_user_id text,
		username text,
		password text,
		user_token text,
		level text,
		extras extra_value,
		is_active boolean,
		is_deleted boolean,
		created_at timestamp,
		updated_at timestamp,
		expired_at timestamp
	)`); err != nil {
		return err
	}
	if err := RepoClient.GetSession().ExecStmt(`CREATE INDEX IF NOT EXISTS idx_ref_user_id ON user (ref_user_id)`); err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Insert(ctx context.Context, user ...*model.User) error {
	q := qb.Insert(repo.tableName).Columns(
		"application_id",
		"user_id",
		"ref_user_id",
		"username",
		"password",
		"user_token",
		"level",
		"extras",
		"is_active",
		"is_deleted",
		"created_at",
		"updated_at",
		"expired_at",
	).Query(*RepoClient.GetSession()).WithContext(ctx)
	for _, u := range user {
		if err := q.BindStruct(u).Exec(); err != nil {
			return err
		}
	}
	q.Release()
	return nil
}

func (repo *UserRepository) FindByUsernameAndPassword(ctx context.Context, username, password string) (*model.User, error) {
	user := new(model.User)
	q := qb.Select(repo.tableName).Columns("*").Where(qb.Eq("username")).Where(qb.Eq("password")).
		Limit(1).
		AllowFiltering().
		Query(*RepoClient.GetSession()).
		BindMap(qb.M{
			"username": username,
			"password": password,
		})
	qa := gocqlx.Query(q.Query, nil)
	err := qa.GetRelease(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) FindByUserToken(ctx context.Context, token string) (*model.User, error) {
	user := new(model.User)
	q := qb.Select(repo.tableName).Columns("*").Where(qb.Eq("user_token")).
		Limit(1).
		AllowFiltering().
		Query(*RepoClient.GetSession()).
		BindMap(qb.M{
			"user_token": token,
		})
	qa := gocqlx.Query(q.Query, nil)
	err := qa.GetRelease(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
