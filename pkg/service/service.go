package service

import (
	"fmt"
	"os"

	"github.com/MIHAIL33/Orca-multitask/pkg/path"
	"github.com/MIHAIL33/Orca-multitask/pkg/runner"
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
		orcaRun.StartOrca()
	}



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