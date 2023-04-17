package scylla

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type (
	IGocqlXClient interface {
		Connect() error
		GetSession() *gocqlx.Session
	}
	GocqlXClient struct {
		Config
		session *gocqlx.Session
	}
)

type Config struct {
	Hosts        []string      // hosts
	Port         int           // timeout
	Keyspace     string        // keyspace name
	ProtoVersion int           // protcol version
	CQLVersion   string        // CQL version
	Retry        int           // number of times to retry queries
	Timeout      time.Duration // timeout

	Username string
	Password string
}

func NewGocqlXClient(cfg Config) IGocqlXClient {
	return &GocqlXClient{
		Config: cfg,
	}
}

func (c *GocqlXClient) Connect() error {
	cluster := gocql.NewCluster(c.Hosts...)
	cluster.ProtoVersion = c.ProtoVersion
	cluster.CQLVersion = "3.0.0"
	cluster.Timeout = c.Timeout
	cluster.ConnectTimeout = c.Timeout
	cluster.Port = c.Port
	cluster.NumConns = 2
	cluster.Consistency = gocql.Quorum
	cluster.MaxPreparedStmts = 1000
	cluster.MaxRoutingKeyInfo = 1000
	cluster.PageSize = 5000
	cluster.DefaultTimestamp = true
	cluster.MaxWaitSchemaAgreement = 60 * time.Second
	cluster.ReconnectInterval = 60 * time.Second
	cluster.ConvictionPolicy = &gocql.SimpleConvictionPolicy{}
	cluster.ReconnectionPolicy = &gocql.ConstantReconnectionPolicy{
		MaxRetries: 3,
		Interval:   1 * time.Second}
	cluster.WriteCoalesceWaitTime = 200 * time.Microsecond
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: c.Username,
		Password: c.Password,
	}
	if c.Retry > 0 {
		cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: c.Retry}
	}
	cluster.Keyspace = c.Keyspace
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return err
	}
	c.session = &session
	return nil
}

func (c *GocqlXClient) GetSession() *gocqlx.Session {
	return c.session
}
