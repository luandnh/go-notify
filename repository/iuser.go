package repository

import (
	repoInterface "github.com/luandnh/go-notify/repository/interface"
	"github.com/luandnh/go-notify/repository/scylla"
)

var UserRepo repoInterface.IUserRepository

func NewUserRepo() repoInterface.IUserRepository {
	if RepoType == SCYLLA {
		return scylla.NewUserRepo()
	} else if RepoType == POSTGRESQL {
		// TODO: implement
	}
	return nil
}
