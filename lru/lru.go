package lru

import (
	"container/list"
)
//type  Value interface {
//	Len() int
//}
type Cache struct {
	maxSpance int64
	usedSpance int64
	ll *list.List
	cacheMap map[string]*list.Element
	CallBackMethod func(key string,value interface{})
}

type Entry struct {
	key string
	value interface{}
}

func NewCache(size int64,callBack func(string, interface{})) *Cache{
	return &Cache{maxSpance:size,ll:list.New(),cacheMap: map[string]*list.Element{},CallBackMethod:callBack}
}

func(c *Cache) AddCache(key string,value interface{} ){
	if ele ,ok := c.cacheMap[key]; ok{
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*Entry)
		c.usedSpance+=int64(len(value.(string)))-int64(len(kv.value.(string)))
		kv.value=value
	}else {
		ele:=c.ll.PushFront(&Entry{key:key,value:value})
		c.cacheMap[key]=ele
		c.usedSpance+=int64(len(key)+ len(value.(string)))
	}
	for c.maxSpance!=0 &&c.usedSpance>c.maxSpance{
		c.removeOldele()
	}
}

func (c *Cache)get(key string)  (value interface{},ok bool)  {
	if ele, ok:= c.cacheMap[key];ok{
		kv:=ele.Value.(*Entry)
		c.ll.MoveToFront(ele)
		return kv.value ,ok
	}
	return
}
func (c *Cache)removeOldele(){
	ele:=c.ll.Back();
	if ele!=nil{
		c.ll.Remove(ele)
		kv:=ele.Value.(*Entry)
		delete(c.cacheMap, kv.key)
		c.usedSpance-=int64(len(kv.key))+int64(len(kv.value.(string)))
		if c.CallBackMethod!=nil{
			c.CallBackMethod(kv.key,kv.value)
		}
	}
}

func (c *Cache) len() int  {
	return c.ll.Len()
}
