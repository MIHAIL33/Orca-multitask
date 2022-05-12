package runner

type OrcaRunner struct {
	dir string
	orca_path string
}

func NewOrcaRunner(dir string, orca_path string) *OrcaRunner {
	return &OrcaRunner{
		dir: dir,
		orca_path: orca_path,
	}
}

func (o *OrcaRunner) Start() error {
	return nil
}