package main
import (
        "fmt"
        "os"
        "strconv"
)
func main()  {
       i := os.Args[1]
       n, err :=strconv.Atoi(i)
       if err != nil  {
              fmt.Println("not a number",i)
              os.Exit(1)
} 

if n%2 == 0 {
              fmt.Println("even number", n)
} else {
              fmt.Println("odd number", n)
       }
  }
