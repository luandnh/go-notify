package repository

import (
	repoInterface "github.com/luandnh/go-notify/repository/interface"
	"github.com/luandnh/go-notify/repository/scylla"
)

var ClientRepo repoInterface.IClientRepository

func NewClientRepo() repoInterface.IClientRepository {
	if RepoType == SCYLLA {
		return scylla.NewClientRepo()
	} else if RepoType == POSTGRESQL {
		// TODO: implement
	}
	return nil
}
