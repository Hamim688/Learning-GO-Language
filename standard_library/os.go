package standardlibrary

import (
	"fmt"
	"os"
)

func OsFunction() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error hostname", err.Error())
	}
}