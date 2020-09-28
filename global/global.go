package global
/**
master 发送缓存数据到这个通道    定时消费通道，同步到slave
 */
var CH =make(chan []interface{},200)

type String string

func (d String) Len() int {
	return len(d)
}
