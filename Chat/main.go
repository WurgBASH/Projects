package main

import(
	"fmt"
	"net"
	"os"
	"bufio"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha1"
	"strconv"
	"math/big"
)

const(
	tcpProtocol = "tcp4"
	keySize = 1024
	readWriterSize = keySize/8
)

type remoteConn struct {
	c *net.TCPConn
	pubK *rsa.PublicKey	
}

func checkErr(err error){ 
	if err != nil {
		fmt.Println(err) 
		os.Exit(1) 
	}
}

var listenAddr = &net.TCPAddr{IP: net.IPv4(192,168,0,101), Port: 8080}

func getRemoteConn(conn *net.TCPConn) *remoteConn{
	return &remoteConn{c: conn, pubK: waitPubKey(bufio.NewReader(conn))}
}


func waitPubKey(buf *bufio.Reader) (*rsa.PublicKey) {
	line, _, err := buf.ReadLine(); checkErr(err)
	if string(line) == "CONNECT" {
		line, _, err := buf.ReadLine(); checkErr(err)
		pubKey := rsa.PublicKey{N: big.NewInt(0)} 
		pubKey.N.SetString(string(line), 10)
		line, _, err = buf.ReadLine(); checkErr(err)
		pubKey.E, err = strconv.Atoi(string(line)); checkErr(err)
		return &pubKey
		
	} else {
		fmt.Println("Error: unkown command ", string(line)) 
		os.Exit(1) 
	}
	return nil
}

func (rConn *remoteConn) sendCommand(comm string) {
	eComm, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, rConn.pubK, []byte(comm), nil)
	checkErr(err)
	rConn.c.Write(eComm)
}


func listen() {
	l, err := net.ListenTCP(tcpProtocol, listenAddr); checkErr(err)
	fmt.Println("Listen port: ", l.Addr().(*net.TCPAddr).Port)
	c, err := l.AcceptTCP(); checkErr(err)
	fmt.Println("Connect from:", c.RemoteAddr())
	rConn := getRemoteConn(c)
	rConn.sendCommand("Go Language Server v0.1 for learning")
	rConn.sendCommand("Привет!")
	rConn.sendCommand("Привіт!")
	rConn.sendCommand("Прывітанне!")
	rConn.sendCommand("Hello!")
	rConn.sendCommand("Salut!")
	rConn.sendCommand("ハイ!")
	rConn.sendCommand("您好!")
	rConn.sendCommand("안녕!")
	rConn.sendCommand("Hej!")
}

func main() {
	listen()
}