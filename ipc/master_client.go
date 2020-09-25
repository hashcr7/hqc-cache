package main


/**
消息结构体:
 */
type SendInfo struct {
	data []byte`json:"data"`
}
/**
主节点给从节点发送消息
 */
func (sendInfo *SendInfo)send(){
	for addr, conn:=range connMap{
		if(addr==""){

		}
		conn.Write(sendInfo.data)
	}
}
func main() {
	
}
