package main

import (
    "net"
    "fmt"
    "os"
    "strings"
)

func main() {
    //if len(os.Args) != 4 {
    //    fmt.Println("Not enough argument")
    //}
    conn, err := net.Dial("tcp", "127.0.0.1:8001")
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Fprintln(conn, strings.Join(os.Args[1:], " "))
}
