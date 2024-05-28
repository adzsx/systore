package main

import (
	"log"
	"os"

	"github.com/adzsx/etwas/internal/store"
	"github.com/adzsx/etwas/internal/utils"
)

var (
	help string = `
etwas usage:
  etwas [mode] [flags]

Modes:
  install   Install from filepath
  config    config the tool
  set       set custom attributes

Flags:
  global flags:
    -v      --verbose   enable verbose mode
    --debug             enable debug mode




  install flags:
    -r      restart after installing everything
    
  
  config flags:
    -l            list available config fields
    -i              start interactive TUI
    -c  [key:value]      set custom values for individual fields
  


  `
)

func main() {
	log.SetFlags(0)
	args := os.Args

	input := utils.Format(args)
	log.Println(input)
	utils.Home()
	store.Setup()
}
