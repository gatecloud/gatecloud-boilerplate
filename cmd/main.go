package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Param struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

var (
	param       Param
	version     = "0.1"
	targetFiles = []string{
		"api",
		"libraries",
		"proxy",
	}
)

func main() {
	flag.StringVar(&(param.Name), "name", "", "Enter your project name")
	flag.StringVar(&(param.Version), "version", version, "Get the version of this project")
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	root := dir[0:strings.LastIndex(dir, string(os.PathSeparator))]

	for i := range targetFiles {
		targetFiles[i] = "(\\\\" + targetFiles[i] + "\\\\)"
	}
	regStr := strings.Join(targetFiles, "|")
	match := regexp.MustCompile(regStr).MatchString
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if match(path) {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(root)

}
