package main

import (
	"log"
	"os"

	"github.com/MIHAIL33/Orca-multitask/pkg/service"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	os.Mkdir("logs", 0755)

	logfile, err := os.OpenFile("logs/log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error creating or opening log file: %v", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	service := service.NewService()
	service.Run()

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
