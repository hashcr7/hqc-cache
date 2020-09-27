package ipc


/**
消息结构体:
 */
type SendInfo struct {
	data []byte`json:"data"`

}
/**1
主节点给从节点发送消息
 */
func (sendInfo *SendInfo)Send(){
	for addr, conn:=range ConnMap{
		if(addr==""){

		}
		conn.Write(sendInfo.data)
	}
}

