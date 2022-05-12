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
	dir string
	input string
	output bool
}

type Paths struct {
	orca_path string
	paths []Path
}

func NewPaths() *Paths {
	paths, err := GetPaths()
	if err != nil {
		log.Fatal(err)
	}
	return &Paths{
		orca_path: viper.GetString("path.orca_path"),
		paths: *paths,
	}
}

func GetPaths() (*[]Path, error) {

	var paths []Path

	basePath := viper.GetString("path.work_path")

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
	}

	var countDir int

	for _, file := range files {
		if file.IsDir() {
			countDir++
			path, err := getPath(basePath + "/" + file.Name())
			if err != nil {
				//log.Panic(err)
				fmt.Println(err)
			} else {
				paths = append(paths, *path)
			}
		}
	}

	if countDir == 0 {
		path, err := getPath(basePath)
		if err != nil {
			//log.Panic(err)
			fmt.Println(err)
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
	path.dir = dir

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
			if inpReg.MatchString(file.Name()) { path.input = file.Name() }
			if outReg.MatchString(file.Name()) { path.output = true }
		}
	}

	if path.input == "" {
		return nil, errors.New("there is no .inp file in the directory: " + dir)
	}

	return &path, nil
}
