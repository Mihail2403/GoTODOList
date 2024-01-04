package app

import (
	"context"
	"go_todo_list/internal/repository"
	"go_todo_list/internal/server"
	"go_todo_list/internal/service"
	http_handler "go_todo_list/internal/transport/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func App() {
	// логер
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// конфиг файл
	err := initConfig()
	if err != nil {
		logrus.Fatalf("error init config: %s", err)
	}
	if err = godotenv.Load(); err != nil {
		logrus.Fatalf("err on load env file: %s", err)
	}

	// инициализация бд postgres
	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalln(err)
	}

	// внедрение зависимостей
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := http_handler.NewHandler(services)

	// создание и запуск сервера
	srv := new(server.Server)
	go func() {
		err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
		if err != nil {
			logrus.Fatalf("error on running server: %s", err)
		}
	}()
	logrus.Println("App Started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Println("App Stopped...")
	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error closing connection to database: %s", err.Error())
	}

}
