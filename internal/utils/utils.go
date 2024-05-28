package utils

import(
  "log"
)

var(
  verbose int
)

func Verbose(level int, args ...any){
  if level >= verbose{
    log.Println(args)
  } 
}
