package repository

import (
	repoInterface "github.com/luandnh/go-notify/repository/interface"
	"github.com/luandnh/go-notify/repository/scylla"
)

var ApplicationRepo repoInterface.IApplicationRepository

func NewApplicationRepo() repoInterface.IApplicationRepository {
	if RepoType == SCYLLA {
		return scylla.NewApplicationRepo()
	} else if RepoType == POSTGRESQL {
		// TODO: implement
	}
	return nil
}
