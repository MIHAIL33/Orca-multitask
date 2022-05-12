package main

import (
	"log"

	"github.com/MIHAIL33/Orca-multitask/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	service := service.NewService()
	service.Run()

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// func start() {
// 	cmd := exec.Command("/home/mihail/orca303/orca", "metanol_opt.inp")
// 	cmd.Dir = "/home/mihail/orca_tasks/test/"
// 	fmt.Println(cmd)
// 	err := cmd.Start()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("Waiting for command to finish...")
// 	err = cmd.Wait()
// 	log.Printf("Command finished with error: %v", err)
// }