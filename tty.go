// Package vhs tty.go spawns the ttyd process.
// It runs on the specified port and is generally meant to run in the background
// so that other processes (go-rod) can connect to the tty.
//
// xterm.js is used for rendering the terminal and can be adjusted using the Set command.
//
// Set FontFamily "DejaVu Sans Mono"
// Set FontSize 12
// Set Padding 50
package main

import (
	"fmt"
	"net"
	"os/exec"
)

// randomPort returns a random port number that is not in use.
func randomPort() int {
	addr, _ := net.Listen("tcp", ":0") //nolint:gosec
	_ = addr.Close()
	return addr.Addr().(*net.TCPAddr).Port
}

// buildTtyCmd builds the ttyd exec.Command on the given port.
func buildTtyCmd(port int, shell Shell) *exec.Cmd {
	args := []string{
		fmt.Sprintf("--port=%d", port),
		"--interface", "127.0.0.1",
	}

	args = append(args, shell.Command...)

	//nolint:gosec
	// We use our custom fork of ttyd here which has WebGL rendering
	// enabled.
	cmd := exec.Command("./ttyd", args...)
	if shell.Env != nil {
		cmd.Env = append(shell.Env, cmd.Env...)
	}
	return cmd
}
