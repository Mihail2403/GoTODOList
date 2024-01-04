package app

import (
	"go_todo_list/internal/repository"
	"go_todo_list/internal/server"
	"go_todo_list/internal/service"
	http_handler "go_todo_list/internal/transport/http"
	"os"

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
	err = srv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		logrus.Fatalf("error on running server: %s", err)
	}
}
