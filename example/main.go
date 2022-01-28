package main

import (
	"log"
	"time"

	"github.com/kylin-ops/loadbalance"
	"github.com/kylin-ops/loadbalance/balance"
)

func main(){
	instances1 := []*balance.Instance{
		{Host: "1.1.1.1", Port: 80, Meta: map[string]string{"scheme": "http"}},
		{Host: "1.1.1.1", Port: 81, Meta: map[string]string{"scheme": "http"}},
		{Host: "1.1.1.1", Port: 82, Meta: map[string]string{"scheme": "http"}},
	}

	instances2 := []*balance.Instance{
		{Host: "1.1.1.2", Port: 80, Meta: map[string]string{"scheme": "http"}},
		{Host: "1.1.1.2", Port: 81, Meta: map[string]string{"scheme": "http"}},
		{Host: "1.1.1.2", Port: 82, Meta: map[string]string{"scheme": "http"}},
	}
	for {
		log.Println(loadbalance.DoBalance(balance.BalanceTypeRoundRobin, instances1))
		log.Println(loadbalance.DoBalance(balance.BalanceTypeRoundRobin, instances2))
		time.Sleep(time.Second)
	}
}