package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// UsesScratchImage opens the given Dockerfile and returns true if it uses "FROM scratch".
func UsesScratchImage(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Check if line starts with FROM and includes "scratch"
		if strings.HasPrefix(strings.ToUpper(line), "FROM") && strings.Contains(line, "scratch") {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}
	return false, nil
}

func main() {
	fmt.Println("Running Flag Lint")
	var fail bool
	flag.BoolVar(&fail, "fix", false, "test val for if the init should fail")
	flag.Parse()
	scratch, err := UsesScratchImage("Dockerfile")
	if err != nil {
		fmt.Println("ERROR")
	}
	if scratch {
		os.Exit(1)
	}
	os.Exit(0)
}
