package path

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/spf13/viper"
)

type Path struct {
	Dir string
	Input string
	Output bool
}

type Paths struct {
	Orca_path string
	Paths []Path
}

func NewPaths() *Paths {
	paths, err := getPaths()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	return &Paths{
		Orca_path: viper.GetString("path.orca_path"),
		Paths: *paths,
	}
}

func getPaths() (*[]Path, error) {

	var paths []Path

	basePath := viper.GetString("path.work_path")

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	var countDir int

	for _, file := range files {
		if file.IsDir() {
			countDir++
			path, err := getPath(basePath + "/" + file.Name())
			if err != nil {
				fmt.Println(err)
				log.Println(err)
			} else {
				paths = append(paths, *path)
			}
		}
	}

	if countDir == 0 {
		path, err := getPath(basePath)
		if err != nil {
			return nil, err
		}
		paths = append(paths, *path)
	}

	if len(paths) == 0 {
		return nil, errors.New("no *.inp files found")
	}

	return &paths, nil
}

func getPath(dir string) (*Path, error) {
	var path Path
	path.Dir = dir

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	inpReg, err := regexp.Compile(`.*.inp`)
	if err != nil {
		log.Fatal(err)
	}
	outReg, err := regexp.Compile(`.*.out`)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {
			if inpReg.MatchString(file.Name()) { path.Input = file.Name() }
			if outReg.MatchString(file.Name()) { path.Output = true }
		}
	}

	if path.Input == "" {
		return nil, errors.New("there is no .inp file in the directory: " + dir)
	}

	return &path, nil
}
