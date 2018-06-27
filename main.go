package main

import (
	"fmt"
	"runtime"
	"os/exec"
	"syscall"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		os.Stderr.WriteString(fmt.Sprintf("Error: port argument missing.\n"))
		os.Exit(1);
	}

	port := os.Args[1]

	if _, err := strconv.Atoi(port); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Error: port argument is not a number.\n"))
		os.Exit(1);
	}

	if runtime.GOOS == "windows" {
		command := fmt.Sprintf("(Get-NetTCPConnection -LocalPort %s).OwningProcess -Force", port)
		exec_cmd(exec.Command("Stop-Process", "-Id", command))
	} else {
		command := fmt.Sprintf("lsof -i tcp:%s | grep LISTEN | awk '{print $2}' | xargs kill -9", port)
		exec_cmd(exec.Command("bash", "-c", command))
	}
}

// Execute command and return exited code.
func exec_cmd(cmd *exec.Cmd) {
	var waitStatus syscall.WaitStatus
	if err := cmd.Run(); err != nil {
		if err != nil {
			os.Stderr.WriteString(fmt.Sprintf("Error: %s\n", err.Error()))
		}
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			fmt.Printf("Output: %s\n", []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		fmt.Printf("Output: %s\n", []byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}
}
