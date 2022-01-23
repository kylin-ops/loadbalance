// 各类balance注册到 allBalance中
package balance

type BalanceType string

const(
	BalanceTypeHash BalanceType = "hash"
	BalanceTypeRandom BalanceType = "random"
	BalanceTypeRoundRobin BalanceType = "roundrobin"
	BalanceTypeWRoundRobin BalanceType = "weight_roundrobin"
	BalanceTypeShuffle BalanceType =  "shuffle"
	BalanceTypeShuffle2 BalanceType = "shuffle2"
)

var (
	mgr = BalanceMgr{
		allBalance: make(map[string]Balancer),
	}
)

type BalanceMgr struct {
	allBalance map[string]Balancer
}

func (p *BalanceMgr) registerBalancer(balanceType string, b Balancer) {
	p.allBalance[balanceType] = b
}

func (p *BalanceMgr) getBalancer(balanceType string) (Balancer, bool) {
	bl, ok :=  p.allBalance[balanceType]
	return bl, ok
}

func RegisterBalancer(balanceType string, b Balancer) {
	mgr.registerBalancer(balanceType, b)
}

func GetBalance(balanceType BalanceType) (Balancer, bool) {
	return  mgr.getBalancer(string(balanceType))
}
