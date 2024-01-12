package main

import (
	bilimsearch "BilimSearch"
	"BilimSearch/package/handler"
	"BilimSearch/package/repository"
	"BilimSearch/package/service"
	"log"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type accounts struct {
	Id       int `db:"user_id"`
	Username string `db:"username"`
}
type Testing struct {
	db *sqlx.DB
}

func newTesting(db *sqlx.DB) *Testing {
	return &Testing{db: db}
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("some error with initializiing: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})

	if err != nil {
		log.Fatalf("error with data base: %s", err.Error())
	}
	
	if err != nil {
		log.Fatalf("error with data base: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(bilimsearch.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error with run server: %s", err.Error())
	}
	
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
