package env

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Env struct {
	Db  *gorm.DB
	Cfg config
}

type privateConfig struct {
	DbHost     string `default:"localhost" split_words:"true"`
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
		log.Fatalln(err.Error())
	}

	var cfg config
	err = envconfig.Process("rradar", &cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Europe/Warsaw",
		privateCfg.DbHost, privateCfg.DbUser, privateCfg.DbPassword, privateCfg.DbName, privateCfg.DbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to the database")
	}

	return Env{
		Db:  db,
		Cfg: cfg,
	}
}
