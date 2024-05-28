package utils

import(
  "log"
  "os"
  "strings"
)


type Input struct{
  Mode string
  Custom map[string]string
  Flags []string
}

func Format(args []string) Input {
  
  if len(args) < 2 && !InSlice(args, "--help"){
    log.Println("Enter --help for help")
    os.Exit(0)
  } 

  input := Input{}
  input.Mode = args[1]
  input.Custom = make(map[string]string) 

  for index, element := range args{
    switch element[1:]{
      case "c":
        data := strings.Split(args[index+1], ":")
        input.Custom[data[0]] = data[1]
    default:
      if string(element[0]) == "-"{
        input.Flags = append(input.Flags, element)
      }
    }
  }
  return input

}
