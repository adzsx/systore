package utils

import(
  "log"
  "os"
  "strings"
)


type Input struct{
  Mode string
  Restart bool
  TUI bool
  Custom map[string]string
}

func Format(args []string) Input {
  
  if len(args) < 2 && !InSlice(args, "--help"){
    log.Println("Enter --help for help")
    os.Exit(0)
  } 

  input := Input{}
  input.Custom = make(map[string]string) 

  for index, element := range args{
    switch element[1:]{
      case "r":
        input.Restart = true
      case "i":
        input.TUI = true
      case "c":
        data := strings.Split(args[index+1], ":")
        input.Custom[data[0]] = data[1]
      }
  }
  return input
}
