package dbusproxy

import (
	"github.com/godbus/dbus"
)

const (
	SettingsInterface  = NetworkManagerInterface + ".Settings"
	SettingsObjectPath = NetworkManagerObjectPath + "/Settings"

	/* Methods */
	SettingsAddConnection = SettingsInterface + ".AddConnection"
)

type SettingsProxy interface {
	AddConnection(ConnectionSettings) error
}

func NewSettings() (SettingsProxy, error) {
	var s settingsProxy
	return &s, s.init(NetworkManagerInterface, SettingsObjectPath)
}

type settingsProxy struct {
	dbusProxyCommon
}

// AddConnection adds new connection and save it to disk. It returns any D-Bus
// error encountered.
func (s *settingsProxy) AddConnection(settings ConnectionSettings) error {
	var dummy dbus.ObjectPath
	return s.obj.Call(SettingsAddConnection, 0, settings).Store(&dummy)
}
