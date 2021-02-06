package main

import (
	"fmt"
	"os"
)

func main() {
  names:= os.Args[1:]
  tmp:=""
  if len(os.Args) >1{
    for i:=0;i<len(names);i++{
      tmp+=names[i]
      tmp+=" "
    }
    fmt.Println("Hello "+tmp +", Welcome to the jungle") 
  }else{
    fmt.Println("error")
  }
}
