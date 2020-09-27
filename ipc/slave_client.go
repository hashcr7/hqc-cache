package ipc

import (
	"fmt"
	"log"
	"net"
	"time"
)


/**
slave启动时向master register
 */
func Register() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8088")
	if(err!=nil){
		log.Fatal(err)
	}
	//bytes := make([]byte,1)
	//bytes[0]=1
	//conn.Write(bytes)
	go readMsterInfo(conn)
	return conn
}
/**
salva  接收  master 发来的消息
 */
func readMsterInfo(conn net.Conn){
	bytes := make([]byte, 1024)
	for {
		n, err := conn.Read(bytes)
		if(err!=nil){
			fmt.Print(err)
		}
		Data:=bytes[:n]
		fmt.Print("slave接收master发来的消息=",string(Data))

	}
}
/**
向master 发送心跳包
 */
func Heartbeat(conn net.Conn){
	ticker := time.NewTicker(5 * time.Second)
	for{
		select {
		case t := <-ticker.C:
			fmt.Print("时间t={]",t)
			bytes := make([]byte,1)
			bytes[0]=1
			_, err := conn.Write(bytes)
			if(err!=nil){
				fmt.Print(err)
			}
			//time.Sleep(6 * time.Second)
		}
	}

}


