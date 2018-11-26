package main

// import "strUtils"
import "fmt"
import "ggo/strUtils"

func main() {
    a := "abc"
    b := "def"
    c := strUtils.ConnectStrs(a, b)
    fmt.Println(c)
}

