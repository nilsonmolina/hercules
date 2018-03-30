package main

import (
	"fmt"
	"io/ioutil"
	"os"

	git "gopkg.in/src-d/go-git.v4"
)

var err error

func createProject() {
	// create project directory
	err = os.Mkdir(project.name, os.ModePerm)
	handleError(err)
	// cd into project directory
	err = os.Chdir(project.name)
	handleError(err)
	// create files depending on language
	if project.language == "c" {
		createCProject()
	} else if project.language == "go" {
		createGoProject()
	}
	//
	fmt.Println(" - Project created!")
}

func createGoProject() {
	// create files
	err = ioutil.WriteFile(".gitignore", []byte(".*\n\n!/.gitignore"), os.ModePerm)
	if project.flags["-src"] == true {
		err = os.Mkdir("src", os.ModePerm)
		err = os.Chdir("src")
	}
	err = ioutil.WriteFile("main.go", []byte(goMain), os.ModePerm)

	if err != nil {
		fmt.Println(err)
	}
}

func createCProject() {
	author := fmt.Sprint(project.flags["-author"])

	// create root files
	err := ioutil.WriteFile("Makefile", []byte(cMakefile), os.ModePerm)
	err = ioutil.WriteFile("author", []byte(author), os.ModePerm)
	err = ioutil.WriteFile(".gitignore", []byte(".*\n\n!/.gitignore"), os.ModePerm)
	// create src folder and files
	err = os.Mkdir("src", os.ModePerm)
	err = ioutil.WriteFile("src/main.c", []byte(cMain), os.ModePerm)
	// create includes folder and files
	err = os.Mkdir("includes", os.ModePerm)
	err = ioutil.WriteFile("includes/main.h", []byte(cHeaderFile), os.ModePerm)
	// create libs folder and files
	err = os.Mkdir("libs", os.ModePerm)
	if project.flags["-libft"] == true {
		downloadLibft()
	}

	if err != nil {
		fmt.Println(err)
	}
}

func downloadLibft() {
	_, err = git.PlainClone("libs/libft", false, &git.CloneOptions{
		URL:      "https://github.com/nilsonmolina/libft",
		Progress: os.Stdout,
	})
}
