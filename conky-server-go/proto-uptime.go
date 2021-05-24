package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Printf("Error : %v", err)
	}
	log.Println(hostname)

	output, errcmd := exec.Command("uptime").CombinedOutput()
	if errcmd != nil {
		os.Stderr.WriteString(errcmd.Error())
	}
	result := string(output)
	log.Print(result)
}
