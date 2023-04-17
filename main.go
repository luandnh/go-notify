package main

import (
	/// THIRD PARTY PACKAGE

	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/luandnh/go-notify/internal/scylla"
	"github.com/luandnh/go-notify/middleware/auth"
	"github.com/luandnh/go-notify/repository"
	repoScylla "github.com/luandnh/go-notify/repository/scylla"
	"github.com/luandnh/go-notify/service"

	"github.com/luandnh/go-notify/api"
	apiV1 "github.com/luandnh/go-notify/api/v1"

	_ "time/tzdata"

	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Dir      string `env:"CONFIG_DIR" envDefault:"config/config.json"`
	Port     string
	LogType  string
	LogLevel string
	LogFile  string
}

var config Config

func init() {
	if err := env.Parse(&config); err != nil {
		log.Fatal(err)
	}

	viper.SetConfigFile(config.Dir)

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
		panic(err)
	}

	cfg := Config{
		Dir:      config.Dir,
		Port:     viper.GetString(`main.port`),
		LogType:  viper.GetString(`main.log_type`),
		LogLevel: viper.GetString(`main.log_level`),
		LogFile:  viper.GetString(`main.log_file`),
	}

	repoScylla.RepoClient = scylla.NewGocqlXClient(scylla.Config{
		Hosts:    viper.GetStringSlice(`scylla.hosts`),
		Port:     viper.GetInt(`scylla.port`),
		Username: viper.GetString(`scylla.username`),
		Password: viper.GetString(`scylla.password`),
		Keyspace: viper.GetString(`scylla.keyspace`),
		Timeout:  time.Second * 30,
		Retry:    2,
	})
	if err := repoScylla.RepoClient.Connect(); err != nil {
		panic(err)
	} else {
		repoScylla.InitRepo()
	}
	repository.RepoType = repository.SCYLLA
	repository.InitRepo()
	config = cfg
}

func main() {
	_ = os.Mkdir(filepath.Dir(config.LogFile), 0755)
	file, _ := os.OpenFile(config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	setAppLogger(config, file)

	server := api.NewServer()
	auth.AuthMdw = auth.NewLocalAuthMiddleware()
	defaultAdminPassword := viper.GetString(`auth.default_admin_password`)
	defaultAdminToken := viper.GetString(`auth.default_admin_token`)
	apiV1.NewApplicationAPI(server.Engine, service.NewApplication())
	service.UserSvr = service.NewUser(defaultAdminPassword, defaultAdminToken)
	service.ApplicationSvr = service.NewApplication()
	server.Start(config.Port)
}

func setAppLogger(cfg Config, file *os.File) {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	switch cfg.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	switch cfg.LogType {
	case "DEFAULT":
		log.SetOutput(os.Stdout)
	case "FILE":
		if file != nil {
			log.SetOutput(io.MultiWriter(os.Stdout, file))
		} else {
			log.SetOutput(os.Stdout)
		}
	default:
		log.SetOutput(os.Stdout)
	}
}
