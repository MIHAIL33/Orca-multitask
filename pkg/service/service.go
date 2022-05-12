package service

import (
	"fmt"

	"github.com/MIHAIL33/Orca-multitask/pkg/path"
	"github.com/MIHAIL33/Orca-multitask/pkg/runner"
)

type Service struct {}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Run() error {

	paths := path.NewPaths()

	fmt.Println(paths)

	for _, path := range paths.Paths {
		orcaRun := runner.NewOrcaRunner(paths.Orca_path, path)
		orcaRun.StartOrca()
	}



	return nil
}