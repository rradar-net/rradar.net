package env

import (
	"context"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/rradar-net/rradar.net/ent"
	"github.com/rradar-net/rradar.net/internal/users"
	"github.com/rs/zerolog/log"
)

type Env struct {
	Ctx            context.Context
	UserRepository users.Repository
	Cfg            config
}

type privateConfig struct {
	DbHost     string `default:"postgres" split_words:"true"`
	DbUser     string `default:"postgres" split_words:"true"`
	DbPassword string `default:"superuser" split_words:"true"`
	DbName     string `default:"rradar" split_words:"true"`
	DbPort     int    `default:"5432" split_words:"true"`
}

type config struct {
	Port int `default:"8080"`
}

func Init() Env {
	var privateCfg privateConfig
	err := envconfig.Process("rradar", &privateCfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	var cfg config
	err = envconfig.Process("rradar", &cfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			privateCfg.DbHost, privateCfg.DbPort, privateCfg.DbUser, privateCfg.DbName, privateCfg.DbPassword))
	if err != nil {
		log.Fatal().Msgf("failed opening connection to postgres: %v", err)
	}

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal().Msgf("failed creating schema resources: %v", err)
	}

	return Env{
		Ctx:            ctx,
		UserRepository: users.NewRepository(client),
		Cfg:            cfg,
	}
}
