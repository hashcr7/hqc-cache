package main

import (
	"fmt"
	"testing"
)


type String string

func (d String) Len() int {
	return len(d)
}

func TestName(t *testing.T) {
k1,k2,k3 := "key1","key2","key3"
v1,v2,v3 := "value1","value2","value3"
	lru:=newCache(int64(len(k1)+ len(k2)+ len(v1)+ len(v2)),nil)
	lru.addCache(k1,String(v1))
	lru.addCache(k2,String(v2))
	lru.addCache(k3,String(v3))
	fmt.Print(lru.get("key1"))
	fmt.Print(lru.get("key2"))
	fmt.Print(lru.get("key3"))
	fmt.Print( lru.len())
	if _, ok := lru.get("key2"); !ok  {
		t.Fatalf("Removeoldest key1 failed")
	}
}

func main() {

}
