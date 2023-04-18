package main

import (
	"fmt"
	// Uncomment this block to pass the first stage!
	"os"
	"os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage!

	command := os.Args[3]
	// fmt.Println("Command:", command)
	args := os.Args[4:len(os.Args)]
	// fmt.Println("Args:", args)

	cmd := exec.Command(command, args...)
	// fmt.Println("CMD:", cmd)

	if args[0] == "echo_stderr" {
		processSTDERR(cmd)
	}

	if args[0] == "echo" {
		processEcho(cmd)
	}

}

func processSTDERR(cmd *exec.Cmd) {
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println("Error creating stderr pipe:", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	// Read from stderr
	buf := make([]byte, 1024)
	n, err := stderr.Read(buf)
	if err != nil {
		fmt.Println("Error reading from stderr:", err)
		return
	}

	// fmt.Println("STDERR:", string(buf[:n]))
	fmt.Println(string(buf[:n]))

	if err := cmd.Wait(); err != nil {
		fmt.Println("Command finished with error:", err)
		return
	}
}

func processEcho(cmd *exec.Cmd) {
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Err: %v", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
	// fmt.Println("Output: ", string(output))
}
