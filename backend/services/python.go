package services

import (
	"log"
	"os/exec"
)

// ExecutePythonCode executes python code using a LOCAL python interpreter
func ExecutePythonCode(code string) string {
	cmd := exec.Command("/usr/bin/python3.8", "-c", code)

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
