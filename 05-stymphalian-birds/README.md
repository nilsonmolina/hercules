# Labour 05: Stymphalian Bird

Goal: This labour is made to make you more efficient when starting a new project by automating all of the repetitive settings you usually have to do.

> Write a script that creates your project with at least a gitignore, if it’s a C project it has to add a Makefile and src/lib.
You’re script has to take options like the language of your project and the extend of the options (like if you’re including your libft or so).  

## How To Run
To run this script on a macOS system: 

1. Run the script
```
./automate <PROJECT_NAME> <OPTIONS> <SUBOPTIONS>
```

**_*Note:_** *To run this script on another OS, you will need to cross-compile the files in src.*

## Compile to a smaller size
```
go build -ldflags "-s -w" -o <NAME>
```

<!-- Task List -->
## Features to add:
* [x] c   - create c project
* [x] go  - create a go project
* [x] all - create a .gitignore file
* [x] c   - allow custom author
* [x] go  - allow src folder 
* [x] go  - create a main.go file
* [x] c   - create a main.go, main.h, and directories
* [x] c   - create a Makefile
* [x] all - run using a wizard
* [x] all - run using flags
* [x] all - clear terminal between questions on wizard
* [x] c   - allow git clone of libft 
* [x] all - accept a path for project directory
* [ ] web - create an html5 project
* [ ] ORGANIZE CODE STRUCTURE

