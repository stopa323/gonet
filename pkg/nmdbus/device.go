package nmdbus

import "github.com/godbus/dbus"

const (
	DeviceInterface = NetworkManagerInterface + ".Device"

	/* Properties */
	DevicePropertyInterface = DeviceInterface + ".Interface"
	DevicePropertyState     = DeviceInterface + ".State"
)

type DeviceProxy interface {
	GetPropertyInterface() (string, error)
	GetPropertyState() (NMDeviceState, error)
}

func NewDevice(objectPath dbus.ObjectPath) (DeviceProxy, error) {
	var d deviceProxy
	return &d, d.init(NetworkManagerInterface, objectPath)
}

type deviceProxy struct {
	dbusBase
}

func (d *deviceProxy) GetPropertyInterface() (string, error) {
	return d.getStringProperty(DevicePropertyInterface)
}

func (d *deviceProxy) GetPropertyState() (NMDeviceState, error) {
	r, err := d.getUint32Property(DevicePropertyState)
	if err != nil {
		return NMDeviceStateFailed, err
	}
	return NMDeviceState(r), nil
}
