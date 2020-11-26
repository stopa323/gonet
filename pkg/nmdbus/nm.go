package nmdbus

import (
	"github.com/godbus/dbus/v5"
)

const (
	NetworkManagerInterface  = "org.freedesktop.NetworkManager"
	NetworkManagerObjectPath = "/org/freedesktop/NetworkManager"

	/* Methods */
	NetworkManagerGetDevices = NetworkManagerInterface + ".GetDevices"
)

type NetworkManagerProxy interface {
	GetDevices() ([]DeviceProxy, error)
}

type networkManagerProxy struct {
	dbusBase
}

func NewNetworkManager() (NetworkManagerProxy, error) {
	var nm networkManagerProxy
	return &nm, nm.init(NetworkManagerInterface, NetworkManagerObjectPath)
}

func (nm *networkManagerProxy) GetDevices() (devices []DeviceProxy, err error) {
	var devicePaths []dbus.ObjectPath
	err = nm.obj.Call(NetworkManagerGetDevices, 0).Store(&devicePaths)

	if err != nil {
		return
	}

	devices = make([]DeviceProxy, len(devicePaths))

	for i, path := range devicePaths {
		devices[i], err = NewDevice(path)
		if err != nil {
			return
		}
	}
	return
}
