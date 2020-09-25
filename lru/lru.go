package main

import (
	"container/list"
)
type Value interface {
	Len() int
}
type Cache struct {
	maxSpance int64
	usedSpance int64
	ll *list.List
	cacheMap map[string]*list.Element
	CallBackMethod func(key string,value Value)
}

type entry struct {
	key string
	value Value
}

func newCache(size int64,callBack func(string, Value)) *Cache{
     return &Cache{maxSpance:size,ll:list.New(),cacheMap: map[string]*list.Element{},CallBackMethod:callBack}
}

func(c *Cache) addCache(key string,value Value ){
	if ele ,ok := c.cacheMap[key]; ok{
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.usedSpance+=int64(value.Len())-int64(kv.value.Len())
		kv.value=value
	}else {
		ele:=c.ll.PushFront(&entry{key:key,value:value})
		c.cacheMap[key]=ele
		c.usedSpance+=int64(len(key)+value.Len())
	}
	for c.maxSpance!=0 &&c.usedSpance>c.maxSpance{
		c.removeOldele()
	}
}

func (c *Cache)get(key string)  (value Value,ok bool)  {
	if ele, ok:= c.cacheMap[key];ok{
		kv:=ele.Value.(*entry)
		c.ll.MoveToFront(ele)
		return kv.value ,ok
	}
	return
}
func (c *Cache)removeOldele(){
	ele:=c.ll.Back();
	if ele!=nil{
		c.ll.Remove(ele)
		kv:=ele.Value.(*entry)
		delete(c.cacheMap, kv.key)
		c.usedSpance-=int64(len(kv.key))+int64(kv.value.Len())
		if c.CallBackMethod!=nil{
			c.CallBackMethod(kv.key,kv.value)
		}
	}
}

func (c *Cache) len() int  {
	return c.ll.Len()
}

