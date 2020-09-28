package main

import (
	"hqc/consistency"
	"hqc/global"
	"hqc/ipc"
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
	var identifier []interface{}

	//global.CH<-identifier
	serialization := consistency.New_Cache_data_serialization(identifier)
	go tesr(*serialization)
	 serialization.TimerSendInfoToSlave()
	//list2 := serialization.GetCacheList()

	//serialization.TimerSendInfoToSlave()
}

func tesr(serialization consistency.Cache_data_serialization){
	var identifier2 []interface{}
	identifier2 = append(identifier2, "key4", "value4")
	global.CH<-identifier2
	//serialization.Cache_list=identifier2
	var identifier3 []interface{}
	identifier3 = append(identifier3, "key5", "value5")
	global.CH<-identifier3
	//identifier2 = append(identifier2, "key6", "value6")
	//global.CH<-identifier2

}

