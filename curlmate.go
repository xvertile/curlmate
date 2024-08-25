package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: curlmate <proxy-string>")
		os.Exit(1)
	}
	proxyString := os.Args[1]
	components := strings.Split(proxyString, ":")
	if len(components) != 4 {
		fmt.Println("Invalid proxy string format")
		os.Exit(1)
	}
	usernamePassword := components[2] + ":" + components[3]
	proxyURL := components[0] + ":" + components[1]
	curlCmd := fmt.Sprintf("curl --proxy %s@%s https://ipinfo.io", usernamePassword, proxyURL)
	fmt.Println(curlCmd)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", curlCmd)
	} else {
		cmd = exec.Command("sh", "-c", curlCmd)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing curl command:", err)
		os.Exit(1)
	}
}
