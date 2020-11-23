package nmdbus

import "github.com/godbus/dbus"

const (
	DeviceInterface = NetworkManagerInterface + ".Device"

	/* Properties */
	DevicePropertyInterface = DeviceInterface + ".Interface"
)

type DeviceProxy interface {
	GetPropertyInterface() (string, error)
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
