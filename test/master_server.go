package main

import (
	"hqc/consistency"
	"hqc/ipc"
	"time"
)
type String string

func (d String) Len() int {
	return len(d)
}
func main() {
	go master_client()
	ipc.Accept()

}
/**
master  同步数据到slave  测试
 */
func master_client(){

	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case  <-ticker.C:
			//var list list.List
			//cache_map := make(map[string]lru.Value)
			//cache_map["key1"]=String("value1")
			var identifier []interface{}
			bytes := append(identifier, "key1", "value1")

			//list.PushBack("key1")
			//list.PushBack("value1")
			//[5]consistency.List_ele{ele}
			//fmt.Print("来了===========",list)
			//marshal, _ := json.Marshal(bytes)
			serialization := consistency.New_Cache_data_serialization(bytes)

			serialization.Write_slave()
		}
	}
}
