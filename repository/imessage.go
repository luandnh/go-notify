package repository

import (
	repoInterface "github.com/luandnh/go-notify/repository/interface"
	"github.com/luandnh/go-notify/repository/scylla"
)

var MessageRepo repoInterface.IMessageRepository

func NewMessageRepo() repoInterface.IMessageRepository {
	if RepoType == SCYLLA {
		return scylla.NewMessageRepo()
	} else if RepoType == POSTGRESQL {
		// TODO: implement
	}
	return nil
}
