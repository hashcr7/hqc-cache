package consistency

import (
	"encoding/json"
	"fmt"
	"hqc/ipc"
)
type  Value interface {
	Len() int
}
//type mylist list.List
//
//func ( mylist *mylist) PushFront(list_ele List_ele)  {
//	mylist.PushFront(list_ele)
//}

/**
master slave 保证数据一致性 网络传输的数据结构
 */
type Cache_data_serialization struct {
	//存放 master上修改的数据 定时同步slave
	Cache_list []interface{} `json:"cacheList"`
	S int`json:"s"`
    master_client ipc.SendInfo
}


func New_Cache_data_serialization(cache_list []interface{}) *Cache_data_serialization{
	return &Cache_data_serialization{Cache_list:cache_list,S:3}
}

/**
master 缓存数据同步slave
 */
func (c *Cache_data_serialization)Write_slave( ){
	//获取所有slave
	connMap := ipc.GetConnMap()
	//给所有slave发送 同步数据
	for _,conn:=range connMap{
		bytes, err := json.Marshal(c)
		if(err!=nil){
			fmt.Print(err)
		}
		fmt.Print("反序列化数据=",string(bytes))
		conn.Write(bytes)
	}
	c.master_client.Send()
}


