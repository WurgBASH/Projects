package main

import (
	"fmt"
	"net"
	"bufio"
	"container/list"


)
var clients *list.List

func HandleClient(socket net.Conn){
	for {
		buffer, err := bufio.NewReader(socket).ReadString('\n')
		if err !=nil{
			fmt.Println("Error: ", err.Error())
			socket.Close()
			return
		}
		for i := clients.Front(); i != nil; i = i.Next() {
			fmt.Println(i.Value.(net.Conn), buffer)
			//bufio.NewWriter(i.Value.(net.Conn)).WriteString(buffer)
		}
	}
}

func main(){
	fmt.Println("Server status: online")
	clients = list.New()
	server,err := net.Listen("tcp",":8080")
	if err != nil{
		fmt.Println("Error: ", err.Error())
		return
	}
	for{
		client, err := server.Accept()
		if err != nil{
			fmt.Println("Error: ", err.Error())
			return
		}
		fmt.Println("New client is connected: ", client.RemoteAddr())
		clients.PushBack(client)
		go HandleClient(client)
	}
}