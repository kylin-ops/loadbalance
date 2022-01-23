package loadbalance

import (
	"fmt"

	"github.com/kylin-ops/loadbalance/balance"
)
// balanceType: random,weight_roundrobin,roundrobin,hash, shuffle, shuffle 6种类型
func DoBalance(balanceType balance.BalanceType, insts []*balance.Instance) (*balance.Instance, error) {
	Balancer, ok := balance.GetBalance(balanceType)
	if !ok {
		fmt.Printf("Not found %s balancer", balanceType)
		return nil, fmt.Errorf("not fund balancer")
	}
	return Balancer.DoBalance(insts)
}
