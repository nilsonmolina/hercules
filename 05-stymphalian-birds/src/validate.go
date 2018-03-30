package main

import (
	"os"
	"strconv"
)

func validateFlags() {
	// check required -lang flag
	if os.Args[2] != "-lang" || len(os.Args) < 4 {
		showUsage()
	}
	// sets project defaults based on the language
	setProject(os.Args[3], os.Args[1])
	// check remaining flags
	arg := 4
	var err error
	for arg < (len(os.Args)) {
		if flag, ok := project.flags[os.Args[arg]]; !ok {
			showUsage()
		} else {
			switch flag.(type) {
			case string:
				project.flags[os.Args[arg]] = os.Args[arg+1]
			case bool:
				project.flags[os.Args[arg]], err = strconv.ParseBool(os.Args[arg+1])
				handleError(err)
			case int:
				project.flags[os.Args[arg]], err = strconv.Atoi(os.Args[arg+1])
				handleError(err)
			}
		}
		arg = arg + 2
	}
}

// // Instead of the switch statement, the following could also be used:
// if _, ok := flag.(string); ok {
// 	project.flags[os.Args[arg]] = os.Args[arg+1]
// } else if _, ok := flag.(bool); ok {
// 	project.flags[os.Args[arg]], err = strconv.ParseBool(os.Args[arg+1])
// 	handleError(err)
// } else if _, ok := flag.(int); ok {
// 	project.flags[os.Args[arg]], err = strconv.Atoi(os.Args[arg+1])
// 	handleError(err)
// }
