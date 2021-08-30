package main

import (
	"fmt"
	"os"
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
	var commandToExec []string = os.Args[2:]
	fileToWatchInfos, err := os.Stat(fileToWatch) // find a way to use var
	var fileToWatchLastModifyOriginal time.Time = fileToWatchInfos.ModTime()
	if err != nil {
		fmt.Println("Problem this file to watch:", err)
		os.Exit(1)
	}
	for {
		time.Sleep(time.Second)
		var fileToWatchLastModify time.Time = fileToWatchInfos.ModTime()
		if fileToWatchLastModify != fileToWatchLastModifyOriginal {

			fmt.Println(commandToExec)
			fileToWatchLastModifyOriginal = fileToWatchLastModify
		}
	}
}
