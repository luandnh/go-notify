package scylla

import (
	repoInterface "github.com/luandnh/go-notify/repository/interface"
	"github.com/scylladb/gocqlx/v2/table"
)

type (
	MessageRepository struct {
		table *table.Table
	}
)

func NewMessageRepo() repoInterface.IMessageRepository {
	tblMeta := table.Metadata{
		Name: "message",
		Columns: []string{
			"application_id",
			"user_id",
			"message_id",
			"title",
			"content",
			"priority",
			"extras",
			"is_active",
			"is_deleted",
			"created_at",
			"updated_at",
		},
		PartKey: []string{"message_id", "user_id"},
		SortKey: []string{},
	}
	tbl := table.New(tblMeta)
	repo := &MessageRepository{
		table: tbl,
	}
	if err := repo.createTable(); err != nil {
		panic(err)
	}
	return repo
}

func (repo *MessageRepository) createTable() error {
	if err := RepoClient.GetSession().ExecStmt(`CREATE TABLE IF NOT EXISTS message (
		application_id uuid,
		user_id uuid,
		message_id uuid PRIMARY KEY,
		title text,
		content text,
		priority int,
		extras map<text,frozen<extra_value>>,
		is_active boolean,
		is_deleted boolean,
		created_at timestamp,
		updated_at timestamp
	)`); err != nil {
		return err
	}
	return nil
}

// func (repo *MessageRepository) Insert(message *Message) error {
// 	if err := repo.table.Insert(message); err != nil {
// 		return err
// 	}
// 	return nil
// }
