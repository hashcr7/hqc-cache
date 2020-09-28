package consistency

import (
	"fmt"
	"hqc/global"
	"hqc/ipc"
	"time"
)
type  Value interface {
	Len() int
}

/**
master slave 保证数据一致性 网络传输的数据结构
 */
type Cache_data_serialization struct {
	//存放 master上修改的数据 定时同步slave
	Cache_list []interface{} `json:"CacheList"`
   // master_client ipc.SendInfo
}

func New_Cache_data_serialization(cache_list []interface{}) *Cache_data_serialization{
	return &Cache_data_serialization{Cache_list:cache_list}
}

func(c *Cache_data_serialization) GetCacheList () []interface{}{
	return c.Cache_list
}
/**
master  定时同步数据到  所有slave
 */
func (c *Cache_data_serialization)TimerSendInfoToSlave(){
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case  <-ticker.C:
			//serialization := New_Cache_data_serialization(c.Cache_list)
			data,ok:=<-global.CH
			if(ok){
				fmt.Print("通道数据==",data)
				c.Cache_list=data
				c.Write_slave()
				var ii []interface{}
				c.Cache_list=ii
			}
		}
	}
}

/**
write 数据到对应的slave
 */
func (c *Cache_data_serialization)Write_slave( ){
	//获取所有slave
	//connMap := ipc.GetConnMap()
	////给所有slave发送 同步数据
	//for _,conn:=range connMap{
	//	bytes, _ := json.Marshal(c.Cache_list)
	//	if(err!=nil){
	//		fmt.Print(err)
	//	}
	//	fmt.Print("反序列化数据=",string(bytes))
	//	//conn.Write(bytes)
	//}

	sendInfo := ipc.NewSendInfo(c.Cache_list)
	sendInfo.SendInfoSlave()
}





