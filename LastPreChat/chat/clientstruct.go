package chat

import 	"net"

type Client struct{
	id int
	tc *net.TCPConn
}