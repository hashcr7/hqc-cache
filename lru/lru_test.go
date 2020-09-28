package lru

import (
	"fmt"
	"testing"
)




func TestName(t *testing.T) {
k1,k2,k3 := "key1","key2","key3"
v1,v2,v3 := "value1","value2","value3"
	lru:=NewCache(int64(len(k1)+ len(k2)+ len(v1)+ len(v2)),nil)
	lru.AddCache(k1,v1)
	lru.AddCache(k2,v2)
	lru.AddCache(k3,v3)
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
