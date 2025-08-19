package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Running Flag Lint")
	var fail bool
	flag.BoolVar(&fail, "fix", false, "test val for if the init should fail")
	flag.Parse()
	if fail {
		os.Exit(1)
	}
	os.Exit(0)
}
