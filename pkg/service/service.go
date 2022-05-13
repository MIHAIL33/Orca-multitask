package service

import (
	"fmt"
	"log"
	"os"

	"github.com/MIHAIL33/Orca-multitask/pkg/mailer"
	"github.com/MIHAIL33/Orca-multitask/pkg/path"
	"github.com/MIHAIL33/Orca-multitask/pkg/runner"
	"github.com/spf13/viper"
)

type Service struct {}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Run() error {

	paths := path.NewPaths()

	CheckOutputFile(*paths)

	for _, path := range paths.Paths {
		orcaRun := runner.NewOrcaRunner(paths.Orca_path, path)
		err := orcaRun.StartOrca()

		if viper.GetBool("email.mail_successed") {

			if err != nil {
				failedMail := mailer.NewMailer(path.Dir, mailer.Failed)
				err = failedMail.SendMail()
				if err != nil {
					log.Println(err)
				}
				log.Printf("send failed mail: %s", path.Dir)
			} else {
				if len(paths.Paths) > 1 {
					successedMail := mailer.NewMailer(path.Dir, mailer.Successed)
					err = successedMail.SendMail()
					if err != nil {
						log.Println(err)
					}
					log.Printf("send successed mail: %s", path.Dir)
				}
			}
		}
	}

	finishMail := mailer.NewMailer(viper.GetString("path.work_path"), mailer.Finished)
	err := finishMail.SendMail()
	if err != nil {
		log.Println(err)
	}
	log.Printf("send finished mail: %s", viper.GetString("path.work_path"))
	fmt.Printf("send finished mail: %s\n", viper.GetString("path.work_path"))

	return nil
}

func CheckOutputFile(paths path.Paths) {
	flag := false
	for _, path := range paths.Paths {
		if path.Output {
			flag = true
			fmt.Println("directory already has *.out file: " + path.Dir)
		}
	}

	if flag {
		answer := GetAnswer()
		for {
			switch answer {
			case "y":
				return
			case "n":
				os.Exit(0)
			default:
				answer = GetAnswer()
			}
		}
	}
}

func GetAnswer() string {
	fmt.Println("new *.out files will be created, if you wish to continue? (y/n)")
	var answer string
	_, err := fmt.Fscan(os.Stdin, &answer)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return answer
}