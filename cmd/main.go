package main
    
import (
	"os"
  "log"
	"github.com/adzsx/etwas/internal/utils"
  "github.com/adzsx/etwas/internal/store"
)

var(
  help string = `
etwas usage:
  etwas [mode] [flags]

Modes:
  install   Install from filepath
  config    config the tool
  set       set custom attributes

Flags:
  global flags:
    --debug     enable debug mode




  install flags:
    -r      restart after installing everything
    
  
  config flags:
    -l            list available config fields
    -i              start interactive TUI
    -c  [key:value]      set custom values for individual fields
  


  `
)

func main() {
	args := os.Args
  
  input := utils.Format(args)
  log.Println(input)

  store.Setup("~/test-etas.txt")
}
