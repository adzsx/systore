package utils

func InSlice(arr []string, str string) bool {

  for _, v := range arr{
    if v == str{
      return true
    }
  }
  
  return false
}


func Err(err error){
  if err != nil{
    panic(err)
  }
}
