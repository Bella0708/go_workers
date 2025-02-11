package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Path to the PID file
	pidFile := "/run/goserver.pid"

	// Fork the actual server process
	cmd := exec.Command("./server", port) 
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start server process: %v", err)
	}
	fmt.Printf("Server process started with PID: %d\n", cmd.Process.Pid)

	pid := cmd.Process.Pid
	pidStr := strconv.Itoa(pid)
	err = ioutil.WriteFile(pidFile, []byte(pidStr), 0644)
	if err != nil {
		log.Fatalf("Failed to write PID to file: %v", err)
	}
	fmt.Printf("PID written to %s\n", pidFile)

	os.Exit(0)
}
