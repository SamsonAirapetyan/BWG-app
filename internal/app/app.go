package app

import (
	"context"
	"fmt"
	"github.com/SamsonAirapetyan/BWG-test"
	"github.com/SamsonAirapetyan/BWG-test/internal/handler"
	"github.com/SamsonAirapetyan/BWG-test/internal/repository"
	"github.com/SamsonAirapetyan/BWG-test/internal/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func Run() {
	if err := initConfig(); err != nil {
		log.Println("error init config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Println("Error with loading password %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	fmt.Println("After postgreess")
	if err != nil {
		log.Println("failed conection with BD %s", err.Error())
	}
	err = StartDataBase(db)
	if err != nil {
		log.Println("failed inserting to BD %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// handlers := new(handler.Handler)
	srv := new(BWG_test.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()
	log.Print("BWG-app running...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("BWG-app Stop")
	if err = srv.Shutdown(context.Background()); err != nil {
		log.Println("Error with shutting down: %s", err.Error())
	}
	if err = db.Close(); err != nil {
		log.Println("Error with Data Base Close: %s", err.Error())
	}
}

// служит для получения данных с конфиг, для этого нужна библиотека viper
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func StartDataBase(db *sqlx.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS Wallet(wallet_num bigint not null unique,USDT float default 0,RUB float default 0,EUR float default 0)")
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Transaction(id serial not null unique,wallet_num bigint not null,currency char(100) not null,sum float,status char(100) default 'created')")
	_, err = db.Exec("INSERT INTO Wallet (wallet_num) VALUES (1234567812345678)")
	_, err = db.Exec("INSERT INTO Wallet (wallet_num) VALUES (8765567887655678)")
	return err
}
