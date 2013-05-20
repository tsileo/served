package main

import (
    "fmt"
    "net/http"
    "net"
    "log"
    "strings"
    "bufio"
    "strconv"
)

type Server struct {
    path string
    port int
    channel chan bool
}

var data = map[string]Server{}

func serve2(data Server) {
    server := &http.Server{Handler: http.FileServer(http.Dir(data.path))}
    fmt.Println(strconv.FormatInt(int64(data.port), 10))
    ln, err := net.Listen("tcp", ":" + strconv.FormatInt(int64(data.port), 10))
    if err != nil { panic(err.Error()) }
    go func() {
        <-data.channel
        ln.Close()
        fmt.Println("closed")
        fmt.Println(data)
    }()
    server.Serve(ln)
    fmt.Println("finish ahah")
}

func main() {
    l, err := net.Listen("tcp", ":8001")
    if err != nil {
        log.Fatal(err)
    }
    defer l.Close()
    for {
        // Wait for a connection.
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go func(c net.Conn) {
            reader := bufio.NewReader(conn)
            cmd, _ := reader.ReadString('\n')
            out := strings.Split(strings.Trim(string(cmd), "\n"), " ")
            fmt.Println(out)
            if out[0] == "start" {
                c := make(chan bool)
                port, _ := strconv.Atoi(out[2])
                data[out[1]] = Server{out[1], port, c}
                fmt.Println(data)
                serve2(data[out[1]])

            }
            if out[0] == "stop" {
                data[out[1]].channel <- true
                delete(data, out[1])
                fmt.Println(data)
            }
            c.Close()
        }(conn)
    }
}
