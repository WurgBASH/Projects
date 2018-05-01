package main

import(
	"fmt"
	"net"
	"os"
	"bufio"
	"crypto/rsa"
	"crypto/rand"
	"strconv"
)

const(
	tcpProtocol	= "tcp"
	keySize = 1024
	readWriterSize = keySize/8
)

func checkErr(err error){
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var connectAddr = &net.TCPAddr{IP: net.IPv4(192,168,0,101), Port: 8080}


// Считываем с командной строки нужный нам порт и пытаемся соединится с сервером
func connectTo() *net.TCPConn{
	// Выводим текст "Enter port:" без перехода но новую строку
	fmt.Print("Enter port:")
	
	// Считываем число с консоли в десятичном формате "%d"
	fmt.Scanf("%d", &connectAddr.Port)
	// Scanf не возвращает значение зато замечательно работает если передать туда ссылку
	
	fmt.Println("Connect to", connectAddr)
	
	// Создаём соединение с сервером
	c ,err := net.DialTCP(tcpProtocol, nil, connectAddr); checkErr(err)
	return c
}

// Функция в определённом порядке отправляет PublicKey
func sendKey(c *net.TCPConn, k *rsa.PrivateKey) {
	
	// Говорим серверу что сейчас будет передан PublicKey
	c.Write([]byte("CONNECT\n"))
	
	// передаём N типа *big.Int
	c.Write([]byte(k.PublicKey.N.String() + "\n"))
	// String() конвертирует *big.Int в string
	
	// передаём E типа int
	c.Write([]byte(strconv.Itoa(k.PublicKey.E) + "\n"))
	// strconv.Itoa() конвертирует int в string
	
	// []byte() конвертирует "строку" в срез байт
}

func sendMessage(c *net.TCPConn, msg string){
	c.Write([]byte("SEND: "))
	c.Write([]byte(msg+"\n"))
}

// Читает и освобождает определённый кусок буфера
// Вернёт срез байт
func getBytes(buf *bufio.Reader, n int) []byte {
	// Читаем n байт
	bytes, err:= buf.Peek(n); checkErr(err)
	// Освобождаем n байт
	skipBytes(buf, n)
	return bytes
}

// Освобождает, пропускает определённое количество байт
func skipBytes(buf *bufio.Reader, skipCount int){
	for i:=0; i<skipCount; i++ {
		buf.ReadByte()
	}
}

/*func SendTo(c *net.TCPConn){
	var message string
	fmt.Println("Сообщение: ")
	fmt.Scanf("%s", &message)
	sendMessage(c,message) 
}*/

func main() {
	conn := connectTo()
	//var RecMsg []byte
	var message string
	k, err := rsa.GenerateKey(rand.Reader, keySize); checkErr(err)
	sendKey(conn, k)
	fmt.Println("Сообщения: ")
	for {
		fmt.Scanf("%s", &message)
		go sendMessage(conn,message) 
		}
}