package main

// import "strUtils"
import "fmt"
import "GoUtils/strUtils"

func main() {
    a := "abc"
    b := "def"
    c := strUtils.ConnectStrs(a, b)
    fmt.Println(c)
}

