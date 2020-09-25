package main

import (
	"fmt"
	"log"
	"net"
	"time"
)


/**
slave启动时向master register
 */
func register() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8088")
	if(err!=nil){
		log.Fatal(err)
	}
	//bytes := make([]byte,1)
	//bytes[0]=1
	//conn.Write(bytes)
	return conn
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
			n, err := conn.Write(bytes)
			if(err!=nil){
				fmt.Print(err)
			}
			fmt.Print("n=",n)

			//time.Sleep(6 * time.Second)

		}
	}

}

func main(){
	conn := register()
	Heartbeat(conn)

}
