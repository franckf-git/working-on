package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// build a remplacement for entr utility - it will not be has good than
// the original but it is for eductional
func main() {
	if len(os.Args) <= 2 {
		fmt.Println("usage: entr file-towatch.ext [command to exec]")
		os.Exit(1)
		return
	}
	var fileToWatch string = os.Args[1]
	var command []string = os.Args[2:]
	for {
		var fileToWatchInfos, err = os.Stat(fileToWatch)
		if err != nil {
			fmt.Println("Problem this file to watch:", err)
			os.Exit(1)
		}
		var fileToWatchLastModifyOriginal time.Time = fileToWatchInfos.ModTime()

		time.Sleep(time.Second)

		var fileToWatchLastInfos, _ = os.Stat(fileToWatch)
		var fileToWatchLastModify time.Time = fileToWatchLastInfos.ModTime()
		if fileToWatchLastModify != fileToWatchLastModifyOriginal {
			var commandToExec = []string{"-c"}
			commandToExec = append(commandToExec, command...)
			var cmd *exec.Cmd = exec.Command("bash", commandToExec...)
			var stdout, errcmd = cmd.CombinedOutput()
			if errcmd != nil {
				fmt.Println("Problem with this command:", errcmd)
				os.Exit(1)
			}
			fmt.Println(string(stdout))
		}
	}
}
