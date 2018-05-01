package main

import (
	"fmt"
	"net"
	"io"

)

func main(){
	conn, err := net.Dial("tcp",":8080")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	buf := make([]byte, 0, 4096) 
    tmp := make([]byte, 256)     
	for{
		n, err := conn.Read(tmp)
        if err != nil {
            if err != io.EOF {
                fmt.Println("read error:", err)
            }
            break
        }
        //fmt.Println("got", n, "bytes.")
        buf = append(buf, tmp[:n]...)

    }
    fmt.Println("total size:", len(buf))
    //fmt.Println(string(buf))
	/*
	conn, err := net.Dial("tcp", "google.com:80")
    if err != nil {
        fmt.Println("dial error:", err)
        return
    }
    defer conn.Close()
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

    buf := make([]byte, 0, 4096) // big buffer
    tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
    for {
        n, err := conn.Read(tmp)
        if err != nil {
            if err != io.EOF {
                fmt.Println("read error:", err)
            }
            break
        }
        //fmt.Println("got", n, "bytes.")
        buf = append(buf, tmp[:n]...)

    }
    fmt.Println("total size:", len(buf))
    //fmt.Println(string(buf))
	*/
}