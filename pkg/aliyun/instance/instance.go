package instance

//go:generate mockery --name Interface

type Interface interface {
	GetRegionID() (string, error)
	GetZoneID() (string, error)
	GetVSwitchID() (string, error)
	GetPrimaryMAC() (string, error)
	GetInstanceID() (string, error)
	GetInstanceType() (string, error)
}

var defaultIns Interface = &ECS{}

func Init(in Interface) { _ = "STUB: not implemented"; return }

func GetInstanceMeta() Interface { _ = "STUB: not implemented"; return *new(Interface) }
