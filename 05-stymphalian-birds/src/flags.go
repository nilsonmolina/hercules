package main

var project parameters

// Project : project template
type parameters struct {
	name     string
	language string
	flags    map[string]interface{}
}

func setProject(language string, name string) {
	if language == "c" {
		setCProject(name)
	} else if language == "go" {
		setGoProject(name)
	} else {
		showUsage()
	}
}

func setCProject(name string) {
	project = parameters{
		name:     name,
		language: "c",
		flags: map[string]interface{}{
			"-libft":  false,
			"-author": "nmolina",
		},
	}
}

func setGoProject(name string) {
	project = parameters{
		name:     name,
		language: "go",
		flags: map[string]interface{}{
			"-src": false,
		},
	}
}
