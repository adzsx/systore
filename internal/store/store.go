package store

import(
  "os"
  "log"
)

func Setup(location string){
  _, err := os.Create(location)
  
  log.Println(err)
  log.Println("Created file")
}
