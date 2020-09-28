package ipc

import (
	"encoding/json"
	"fmt"
	"hqc/lru"
	"log"
	"net"
	"time"
)
type String string

func (d String) Len() int {
	return len(d)
}
type Client_json struct {
	Data []interface{}
	//S int
}

/**
slave启动时向master register
 */
func Register() net.Conn {
	conn, err := net.Dial("tcp", "localhost:8088")
	if(err!=nil){
		log.Fatal(err)
	}
	go readMsterInfo(conn)
	return conn
}
/**
salva  接收  master 发来的消息
 */
func readMsterInfo(conn net.Conn){
	bytes := make([]byte, 1024)
	lru:=lru.NewCache(100,nil)
	for {
		n, err := conn.Read(bytes)
		if(err!=nil){
			fmt.Print(err)
		}
		Data:=bytes[:n]
		client_json := &Client_json{}
		json.Unmarshal(Data, client_json)
		//fmt.Print("slave接收master发来的消息=",string(Data))
		for i:=0;i< len(client_json.Data);i+=2{
			receved_data:= client_json.Data[i : i+2]
			i1 := receved_data[0]
			i2:= receved_data[1]
			fmt.Print("ele1=",i1,"ele2=",i2)
			lru.AddCache(i1.(string),i2)
		}
		fmt.Print("slave接收master发来的消息=",client_json.Data)
		fmt.Print("slave最新缓存数据=",lru)
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


