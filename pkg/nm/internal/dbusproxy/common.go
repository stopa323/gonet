package dbusproxy

import (
	"github.com/godbus/dbus"
)

type dbusProxyCommon struct {
	// Connection to system message bus.
	conn *dbus.Conn

	// Remote object on which methods can be invoked through D-Bus.
	obj dbus.BusObject
}

func (d *dbusProxyCommon) init(iface string, path dbus.ObjectPath) (err error) {
	// acquire connection to system bus. Connection is shared between multiple
	// objects - it is initialized once.
	d.conn, err = dbus.SystemBus()
	if err != nil {
		return
	}
	d.obj = d.conn.Object(iface, path)
	return
}
