package main

import (
	"os"
)

func validateFlags() {
	// check required -lang flag
	if os.Args[2] != "-lang" || len(os.Args) < 4 {
		showUsage()
	}
	if os.Args[3] != "c" && os.Args[3] != "go" {
		showUsage()
	}
	setProject(os.Args[3], os.Args[1])
	// check remaining flags
	arg := 4
	for arg < (len(os.Args)) {
		if _, ok := project.flags[os.Args[arg]]; ok {
			project.flags[os.Args[arg]] = os.Args[arg+1]
		} else {
			showUsage()
		}
		arg = arg + 2
	}
}
