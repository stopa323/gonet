package nmdbus

type NMDeviceState uint32

const (
	NMDeviceStateUnknown      NMDeviceState = 0
	NMDeviceStateUnmanaged    NMDeviceState = 10
	NMDeviceStateUnavailable  NMDeviceState = 20
	NMDeviceStateDisconnected NMDeviceState = 30
	NMDeviceStatePrepare      NMDeviceState = 40
	NMDeviceStateConfig       NMDeviceState = 50
	NMDeviceStateNeedAuth     NMDeviceState = 60
	NMDeviceStateIPConfig     NMDeviceState = 70
	NMDeviceStateIPCheck      NMDeviceState = 80
	NMDeviceStateSecondaries  NMDeviceState = 90
	NMDeviceStateActivated    NMDeviceState = 100
	NMDeviceStateDeactivating NMDeviceState = 110
	NMDeviceStateFailed       NMDeviceState = 120
)
