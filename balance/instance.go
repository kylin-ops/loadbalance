package balance

import "strconv"

type Instance struct {
	Host      string
	Port      int64
	Weight    int64
	CallTimes int64
	Type      string
	Meta      map[string]string
}

func NewInstance(host string, port int64, weigth int64, meta map[string]string) *Instance {
	return &Instance{
		Host:      host,
		Port:      port,
		Weight:    weigth,
		Meta:      meta,
		CallTimes: 0,
	}
}

func (i *Instance) GetPort() int64 {
	return i.Port
}

func (i *Instance) GetHost() string {
	return i.Host
}

func (i *Instance) GetCallTimes() int64 {
	return i.CallTimes
}

func (i *Instance) GetResult() string {
	return i.Host + ":" + strconv.FormatInt(i.Port, 10) + ";call times: " + strconv.FormatInt(i.CallTimes, 10)
}
