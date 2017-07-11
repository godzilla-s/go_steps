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
	err := os.Mkdir(dir, 0755)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func deleteProjDir(dir string) error {
	return nil
}

func deleteAllProjDir(proj string) error {
	err := os.RemoveAll(proj)
	return err
}

func getGopathEnv() (string, error) {
	return os.Getenv(GoPath), nil
}

func isFileExist(proj string) bool {
	_, err := os.Stat(proj)
	return err == nil || os.IsExist(err)
}

func newFile(filename string) bool {
	var err error
	var f	*os.File
	if isFileExist(filename) {
		f, err = os.OpenFile(filename, os.O_RDWR | os.O_APPEND, 0644)
	} else {
		f, err = os.Create(filename)
	}
	
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString("Hello")
	return true
}

func main() {
	dir := getWorkDir()
	fmt.Println(dir)
	
	env, _ := getGopathEnv()
	fmt.Println(env)

	if isFileExist("sms77.go") {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	//deleteAllProjDir("test")
/*
	err := createProjDir("test")
	if err != nil {
		log.Fatal(err)
	}	

	err = createProjDir("test/src")
	if err != nil {
		log.Fatal(err)
	}

	err = createProjDir("test/pkg")
	if err != nil {
        log.Fatal(err)
    }

	err = createProjDir("test/bin")
	if err != nil {
        log.Fatal(err)
    }
*/

}
