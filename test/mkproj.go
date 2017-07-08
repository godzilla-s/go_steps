package main

import (
	"fmt"
	"log"
	"os"
	"flag"
)

var (
	GoPath = "GOPATH"
)

func init() {
	flag.String("c", "temp", "Create New Project")
	flag.String("d", "temp", "Delete Project")
}

func getWorkDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func createProjDir(dir string) error {
	err := os.Mkdir(dir, 0644)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func getGopathEnv() (string, error) {
	return os.Getenv(GoPath), nil
}

func main() {
	dir := getWorkDir()
	fmt.Println(dir)
	
	env, _ := getGopathEnv()
	fmt.Println(env)

	os.Setenv("uuid", "12354353")
}
