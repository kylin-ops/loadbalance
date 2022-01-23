# 1. 使用范例
```go
package main

import (
	"log"
	"time"

	"github.com/kylin-ops/loadbalance"
	"github.com/kylin-ops/loadbalance/balance"
)

func main(){
	instances1 := []*balance.Instance{
		{Host: "1.1.1.1", Port: 80},
		{Host: "1.1.1.1", Port: 81},
		{Host: "1.1.1.1", Port: 82},
	}

	instances2 := []*balance.Instance{
		{Host: "1.1.1.2", Port: 80},
		{Host: "1.1.1.2", Port: 81},
		{Host: "1.1.1.2", Port: 82},
	}
	for {
		log.Println(loadbalance.DoBalance(balance.BalanceTypeRoundRobin, instances1))
		log.Println(loadbalance.DoBalance(balance.BalanceTypeRoundRobin, instances2))
		time.Sleep(time.Second)
	}
}
```

