package runner

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/MIHAIL33/Orca-multitask/pkg/path"
)

type OrcaRunner struct {
	orca_path string
	path path.Path
}

func NewOrcaRunner(orca_path string, path path.Path) *OrcaRunner {
	return &OrcaRunner{
		orca_path: orca_path,
		path: path,
	}
}

func (o *OrcaRunner) StartOrca() error {
	cmd := exec.Command(o.orca_path, o.path.Input)
	cmd.Dir = o.path.Dir

	outNameFile := getNameOutputFile(o.path.Dir + "/" + o.path.Input)
	outfile, err :=  os.Create(outNameFile)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	cmd.Stdout = outfile

	err = cmd.Start()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for calculation: %s", cmd.Dir)
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
	return nil
}

func getNameOutputFile(input string) string {
	strs := strings.Split(input, ".")
	strs[len(strs) - 1] = "out"
	return strings.Join(strs, ".")
}