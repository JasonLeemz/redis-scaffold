package test

import (
	"context"
	"fmt"
	"github.com/JasonLeemz/redis-scaffold/gedis"
	"testing"
)

func TestStringOperation(t *testing.T) {
	ctx := context.Background()
	k1 := "cache:equip:召唤师峡谷:0:1001"
	k2 := "cache:equip:召唤师峡谷:0:1002"
	k3 := "cache:equip:召唤师峡谷:0:1003"
	k4 := "cache:equip:召唤师峡谷:0:1004"
	//
	//ret := gedis.
	//	NewStringOperation(ctx).
	//	Get(k1).
	//	Unwrap_OR("1002")
	//fmt.Println(ret)

	it := gedis.NewStringOperation(ctx).MGet(k1, k2, k3, k4).Iterator()

	for it.HasNext() {
		fmt.Println(it.Next())
	}
	//
	//ret3 := gedis.NewStringOperation(ctx).MGet(k1, k2, k3, k4).Unwrap()
	//fmt.Println(ret3)
}
