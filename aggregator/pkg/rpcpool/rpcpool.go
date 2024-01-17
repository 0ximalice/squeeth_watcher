package rpcpool

type RPCPoolList interface {
	Iter() string
}

type rpcPoolList struct {
	rpcs []string
	r    int
}

func NewRPCPoolList(rpcs []string) RPCPoolList {
	return &rpcPoolList{rpcs: rpcs}
}

func (r *rpcPoolList) Iter() string {
	defer func() {
		r.r++
	}()
	return r.rpcs[r.r%len(r.rpcs)]
}
