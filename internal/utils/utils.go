package utils

import (
	"log"
	"os"
)

var (
	verbose    int
	ConfigPath string
)

func Verbose(level int, args ...any) {
	if level >= verbose {
		log.Println(args...)
	}
}

func Home() {
	var dir string
	dir, err := os.UserHomeDir()
	if err != nil {
		Verbose(1, "No home directory found. Using /root instead")
		dir = "/root"
	}
	ConfigPath = dir + "/.systore/"
}
