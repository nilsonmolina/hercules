package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func createProject() {
	// create project directory
	handleErrorMessage(os.Mkdir(project.name, os.ModePerm), "directory already exists. Please remove directory and try again.")
	// cd into project directory
	handleError(os.Chdir(project.name))
	// create files depending on language
	if project.language == "c" {
		createCProject()
	} else if project.language == "go" {
		createGoProject()
	}
	fmt.Println(" - Project created!")
}

func createGoProject() {
	// create files
	handleError(ioutil.WriteFile(".gitignore", []byte(".*\n\n!/.gitignore"), os.ModePerm))
	if project.flags["-src"] == true {
		handleError(os.Mkdir("src", os.ModePerm))
		handleError(os.Chdir("src"))
	}
	handleError(ioutil.WriteFile("main.go", []byte(goMain), os.ModePerm))
}

func createCProject() {
	author := fmt.Sprintf("%v\n", project.flags["-author"])

	// create root files
	handleError(ioutil.WriteFile("Makefile", []byte(cMakefile), os.ModePerm))
	handleError(ioutil.WriteFile("author", []byte(author), os.ModePerm))
	handleError(ioutil.WriteFile(".gitignore", []byte(".*\n\n!/.gitignore"), os.ModePerm))
	// create src folder and files
	handleError(os.Mkdir("src", os.ModePerm))
	handleError(ioutil.WriteFile("src/main.c", []byte(cMain), os.ModePerm))
	// create includes folder and files
	handleError(os.Mkdir("includes", os.ModePerm))
	handleError(ioutil.WriteFile("includes/main.h", []byte(cHeaderFile), os.ModePerm))
	// create libs folder and files
	handleError(os.Mkdir("libs", os.ModePerm))
	if project.flags["-libft"] == true {
		gitClone("libs/libft", "https://github.com/nilsonmolina/libft")
	}
}

func gitClone(path string, url string) {
	cmd := exec.Command("git", "clone", url, path)
	cmd.Stdout = os.Stdout
	handleError(cmd.Run())
}
