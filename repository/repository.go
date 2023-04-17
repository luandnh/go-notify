package repository

const SCYLLA = "scylla"
const POSTGRESQL = "postgresql"

var RepoType string

func InitRepo() {
	ApplicationRepo = NewApplicationRepo()
	UserRepo = NewUserRepo()
	ClientRepo = NewClientRepo()
	MessageRepo = NewMessageRepo()
}
