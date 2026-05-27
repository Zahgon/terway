/*
Copyright 2021-2022 Terway Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vswitch

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/singleflight"
	"k8s.io/apimachinery/pkg/util/cache"

	"github.com/AliyunContainerService/terway/pkg/aliyun/client"
)

var ErrNoAvailableVSwitch = errors.New("no available vSwitch")
var ErrIPNotEnough = errors.New("no ip left")

// Switch hole all switch info from both terway config and podNetworking
type Switch struct {
	ID   string
	Zone string

	AvailableIPCount int64 // for ipv4
	IPv4CIDR         string
	IPv6CIDR         string
}

type ByAvailableIP []Switch

func (a ByAvailableIP) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a ByAvailableIP) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (a ByAvailableIP) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// SwitchPool contain all vSwitches
type SwitchPool struct {
	cache *cache.LRUExpireCache
	ttl   time.Duration

	g singleflight.Group
}

// NewSwitchPool create pool and set vSwitches to pool
func NewSwitchPool(size int, ttl string) (*SwitchPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetOne get one vSwitch by zone and limit in ids
func (s *SwitchPool) GetOne(ctx context.Context, client client.VPC, zone string, ids []string, opts ...SelectOption) (*Switch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// lookup all vsw in cache and get one matched
// try sort the vsw

// keep the below logic untouched

// lookup all vsw in cache and get one matched

// GetByID will get vSwitch info from local store or openAPI
func (s *SwitchPool) GetByID(ctx context.Context, client client.VPC, id string) (*Switch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SwitchPool) Block(id string) { _ = "STUB: not implemented"; return }

// Add Switch to cache. Test purpose.
func (s *SwitchPool) Add(sw *Switch) { _ = "STUB: not implemented"; return }

// Del Switch from cache. Test purpose.
func (s *SwitchPool) Del(key string) { _ = "STUB: not implemented"; return }

type SelectionPolicy string

// VSwitch Selection Policy
const (
	VSwitchSelectionPolicyOrdered SelectionPolicy = "ordered"
	VSwitchSelectionPolicyRandom  SelectionPolicy = "random"
	VSwitchSelectionPolicyMost    SelectionPolicy = "most"
)

type SelectOption interface {
	// Apply applies this configuration to the given select options.
	Apply(*SelectOptions)
}

// SelectOptions contains options for requests.
type SelectOptions struct {
	IgnoreZone bool

	VSwitchSelectPolicy SelectionPolicy
}

// ApplyOptions applies the given select options on these options
func (o *SelectOptions) ApplyOptions(opts []SelectOption) *SelectOptions {
	_ = "STUB: not implemented"
	return nil
}

var _ SelectOption = &SelectOptions{}

// Apply implements SelectOption.
func (o *SelectOptions) Apply(so *SelectOptions) { _ = "STUB: not implemented"; return }
