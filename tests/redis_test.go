package test

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {

	bm, error := cache.NewCache("redis", `{"conn":"","key":"collectionName","dbNum":"0","password":""}`)
	if error != nil {
		fmt.Println("redis error:", error)
	}
	bm.Put("test", "hello", time.Second*100)
	v := bm.Get("test")
	fmt.Println("value:", string(v.([]byte)))
}
