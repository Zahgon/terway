//go:build !linux

package daemon

import (
	"context"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/AliyunContainerService/terway/types"
)

func gcLeakedRules(existIP sets.Set[string]) { _ = "STUB: not implemented"; return }

func gcPolicyRoutes(ctx context.Context, mac string, containerIPNet *types.IPNetSet, namespace, name string) error {
	_ = "STUB: not implemented"
	return nil
}
