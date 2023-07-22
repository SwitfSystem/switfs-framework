package daemon

import (
	"fmt"
	"os"
	"os/exec"
)

var isDaemon = false

func IsDaemonMode() bool {
	return isDaemon
}

func init() {
	argc := len(os.Args)
	if argc == 1 {
		// redirectStderr()
	} else if argc == 2 {
		if os.Args[1] == "--daemon=true" {
			os.Args[1] = "--daemon"
			cmd := exec.Command(os.Args[0], os.Args[1:]...)
			if err := cmd.Start(); err != nil {
				fmt.Printf("start error :  %s", err.Error())
			}
			fmt.Printf("Server running in daemon . PID [%v]", cmd.Process.Pid)

			os.Exit(0)
		} else if os.Args[1] == "--daemon" {
			if err := os.Stdin.Close(); err != nil {
				fmt.Printf("start error :  %s", err.Error())
			}
			if err := os.Stdout.Close(); err != nil {
				fmt.Printf("start error :  %s", err.Error())
			}
			isDaemon = true

			i := 1
			os.Args = append(os.Args[:i], os.Args[i+1:]...) // 恢复参数位置
			fmt.Println("Daemon Server Initializing...")
		} else {
		}
	} else {
		fmt.Println("The number of arguments incorrect .")
		os.Exit(-1)
	}

}
