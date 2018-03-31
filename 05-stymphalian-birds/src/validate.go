package main

import (
	"fmt"
	"os"
	"strconv"
)

func validateFlags() {
	// check required -lang flag
	if os.Args[2] != "-lang" || len(os.Args) < 4 || len(os.Args)%2 != 0 {
		showUsage()
	}
	// sets project defaults based on the language
	setProject(os.Args[3], os.Args[1])
	// check remaining flags
	arg := 4
	for arg < len(os.Args) {
		flag := os.Args[arg]
		value := os.Args[arg+1]
		if v, ok := project.flags[flag]; !ok {
			showUsage()
		} else {
			switch v.(type) {
			case string:
				project.flags[flag] = value
			case bool:
				project.flags[flag], err = strconv.ParseBool(value)
				handleErrorMessage(err, fmt.Sprintf("The '%v' flag requires a value of type bool", flag))
			case int:
				project.flags[flag], err = strconv.Atoi(value)
				handleErrorMessage(err, fmt.Sprintf("The '%v' flag requires a value of type int", flag))
			}
		}
		arg = arg + 2 // '+ 2' because flags come in {-flag value} pairs (ex. -src true)
	}
}

// // Instead of the switch statement, the following could also be used: (type assertion)
// if _, ok := flag.(string); ok {
// 	project.flags[os.Args[arg]] = os.Args[arg+1]
// } else if _, ok := flag.(bool); ok {
// 	project.flags[os.Args[arg]], err = strconv.ParseBool(os.Args[arg+1])
// 	handleError(err)
// } else if _, ok := flag.(int); ok {
// 	project.flags[os.Args[arg]], err = strconv.Atoi(os.Args[arg+1])
// 	handleError(err)
// }
