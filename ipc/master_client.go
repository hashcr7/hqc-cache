package ipc

import (
	"encoding/json"
	"fmt"
)

/**
消息结构体:
 */
type SendInfo struct {
	Data []interface{}`json:"data"`

}
func NewSendInfo(data []interface{})*SendInfo{
	return &SendInfo{Data:data}
}
/**1
主节点给从节点发送缓存数据  保证数据一致性
 */
func (s *SendInfo)SendInfoSlave(){
	for addr, conn:=range ConnMap{
		if(addr==""){

		}
		//fmt.Print("nani===",string(s.Data))
		bytes, _ := json.Marshal(s)
		fmt.Print("nani===",string(bytes))
		conn.Write(bytes)
	}
}

