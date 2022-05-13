package runner

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

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
		fmt.Println(err)
		log.Fatal(err)
	}

	log.Printf("waiting for calculation: %s", cmd.Dir)
	fmt.Printf("waiting for calculation: %s\n", cmd.Dir)
	err = cmd.Wait()
	if err != nil {
		log.Printf("calculation %s finished with error: %v", cmd.Dir, err)
		fmt.Printf("calculation %s finished with error: %v\n", cmd.Dir, err)
	}
	log.Printf("calculation %s finished without errors", cmd.Dir)
	fmt.Printf("calculation %s finished without errors\n", cmd.Dir)
	return nil
}

func getNameOutputFile(input string) string {
	strs := strings.Split(input, ".")
	strs[len(strs) - 1] = "out"
	nameOut := strings.Join(strs, ".")

	if _, err := os.Stat(nameOut); err != nil {
		if os.IsNotExist(err) {
			return nameOut
		}
	} 

	strs[len(strs) - 2] = strs[len(strs) - 2] + "_" + time.Now().Format("20060102150405")
	nameOut = strings.Join(strs, ".")
	return nameOut
}