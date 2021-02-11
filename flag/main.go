package main

import (
	"flag"
	"log"
)

func main() {
	var name string

	flag.Parse()
	args := flag.Args()
	if len(args) <= 0 {
		return
	}
	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "go语言", "帮助信息")
		goCmd.Parse(args[1:])
	case "php":
		goCmd := flag.NewFlagSet("php", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "php语言", "帮助信息")
		goCmd.Parse(args[1:])
	}
	log.Printf("name= %s", name)
}
