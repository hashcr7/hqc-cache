package ipc

import (
	"fmt"
	"log"
	"net"
	"time"
)
var(
	ConnMap=make(map[string]net.Conn)
	//conn_ch =make(chan map[string]net.Conn)
)
func GetConnMap() map[string]net.Conn{
	return ConnMap
}
/**
服务器阻塞接收客户端信息
 */
func Accept(){
	listener, err := net.Listen("tcp", "localhost:8088")
	if err!=nil{
		log.Fatal(err)
	}
	for{
		conn, err := listener.Accept()
		if(err!=nil){
			fmt.Print(err)
			log.Fatal(err)
		}
		//设置短连接(10秒) 10秒内server没有响应 则该conn超时结束
		conn.SetReadDeadline(time.Now().Add(time.Duration(15)*time.Second))
		//处理conn
		go handleConn(conn)
	}
}

/**
保存从节点的conn信息
 */
func handleConn(conn net.Conn){
	addr := conn.RemoteAddr().String()
	if(ConnMap[addr]==nil){
		ConnMap[addr]=conn
	}
	fmt.Print("connMap={}",ConnMap)

	buffer := make([]byte, 1024)
	for{
		n, err := conn.Read(buffer)
		if(err!=nil){
			fmt.Print(err)
			fmt.Print("conn dead now")
			delete(ConnMap,conn.RemoteAddr().String())
			return
		}
		fmt.Print("nnnn=",n)
		receive_msg_ch := make(chan byte)
		Data := buffer[:n]
		go GravelChannel(Data,receive_msg_ch)
		go HeartBeating(conn,6,receive_msg_ch)
	}
}

func GravelChannel(bytes []byte, mess chan byte) {
	for _, v := range bytes{
		fmt.Print("检测=",v," ")
		mess <- v
	}
	close(mess)
}
/**
心跳监测，有消息进来 ，延长链接时间
 */
func HeartBeating(conn net.Conn,timeout int,receive_msg_ch chan byte){
		select {
		//心跳包有信息，则延长conn时间
		case fk:=<-receive_msg_ch:
			fmt.Print(conn.RemoteAddr().String(), "心跳:第", string(fk), "times")
			conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
			break
		}
}

func main(){
	Accept()
}
