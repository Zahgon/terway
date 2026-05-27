package instance

import (
	"sync/atomic"
)

type ECS struct {
	regionID     atomic.Value
	zoneID       atomic.Value
	vSwitchID    atomic.Value
	primaryMAC   atomic.Value
	instanceID   atomic.Value
	instanceType atomic.Value
}

func (e *ECS) GetRegionID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *ECS) GetZoneID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *ECS) GetVSwitchID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *ECS) GetPrimaryMAC() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *ECS) GetInstanceID() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (e *ECS) GetInstanceType() (string, error) { _ = "STUB: not implemented"; return "", nil }
