package main

import (
	"gotour/time/cmd"
	"log"
)

func main() {
	err := cmd.Execue()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
