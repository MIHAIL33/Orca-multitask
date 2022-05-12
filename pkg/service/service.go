package service

import (
	"fmt"

	"github.com/MIHAIL33/Orca-multitask/pkg/path"
	"github.com/MIHAIL33/Orca-multitask/pkg/runner"
)

type Service struct {
	orcaRun runner.OrcaRunner
	path path.Path
}

func NewService() *Service {
	return &Service{
		orcaRun: runner.OrcaRunner{},
		path: path.Path{},
	}
}

func (s *Service) Run() error {

	paths := path.NewPaths()

	fmt.Println(paths)

	return nil
}