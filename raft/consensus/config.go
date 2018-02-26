package consensus

import (
	"github.com/hashicorp/raft"
	"github.com/rs/xid"
)

var rafts map[string]*raft.Raft

func init() {
	rafts = make(map[string]*raft.Raft)
}

func Config(num int) {
	conf := raft.DefaultConfig()
	conf.LocalID = raft.ServerID(xid.New().String())
	snapshotStore := raft.NewDiscardSnapshotStore()

	addrs := []string{}
	transports := []*raft.InmemTransport{}
	for i := 0; i < num; i++ {
		addr, transport := raft.NewInmemTransport("")
		addrs = append(addrs, string(addr))
		transports = append(transports, transport)
	}
	//peerStore := &raft.StaticPeers{StaticPeers: addrs}
	memStore := raft.NewInmemStore()

	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i != j {
				transports[i].Connect(raft.ServerAddress(addrs[j]), transports[j])
			}

		}
		r, err := raft.NewRaft(conf, NewFSM(), memStore, memStore, snapshotStore, transports[i])
		if err != nil {
			panic(err)
		}
		//r.SetPeers(addrs)
		rafts[addrs[i]] = r
	}
}
