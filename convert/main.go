package main

import (
	"gotour/convert/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}
