package main

var project Project

// Project : project template
type Project struct {
	name     string
	language string
	flags    map[string]string
}

func setProject(lang string, name string) {
	if lang == "c" {
		setCProject(name)
	} else if lang == "go" {
		setGoProject(name)
	}
}

func setCProject(name string) {
	project = Project{
		name:     name,
		language: "c",
		flags: map[string]string{
			"-libft":  "false",
			"-author": "nmolina",
		},
	}
}

func setGoProject(name string) {
	project = Project{
		name:     name,
		language: "go",
		flags: map[string]string{
			"-web": "false",
		},
	}
}
