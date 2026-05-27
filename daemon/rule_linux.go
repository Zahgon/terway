package daemon

import (
	"context"

	"github.com/AliyunContainerService/terway/types/daemon"
)

func ruleSync(ctx context.Context, res daemon.PodResources) error {
	_ = "STUB: not implemented"
	return nil
}

// check host veth ,make sure pod is present

// 1. route point to hostVeth

// default via 10.xx.xx.253 dev eth1 onlink table 1003

// 10.xx.xx.xx dev calixx scope link
