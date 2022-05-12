package http

import (
	"github.com/sirupsen/logrus"

	a "chi-test/pkg/api"
	"chi-test/pkg/config"
	"chi-test/pkg/storage/postgres"
)

type api struct {
	Storage postgres.Postgres
	Config  *config.Config
	Log     *logrus.Logger
}

func New(postgres postgres.Postgres, conf *config.Config) a.API {
	return &api{
		Storage: postgres,
		Config:  conf,
		Log:     logrus.New(),
	}
}

func (a *api) Run() {
	router := a.NewRouter()
	a.Log.Fatal(router.Start(a.Config.Port))
}
