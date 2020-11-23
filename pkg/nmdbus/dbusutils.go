package nmdbus

import (
	"fmt"

	"github.com/godbus/dbus"
	log "github.com/sirupsen/logrus"
)

type dbusBase struct {
	conn *dbus.Conn
	obj  dbus.BusObject
}

func (d *dbusBase) init(iface string, objectPath dbus.ObjectPath) error {
	var err error
	d.conn, err = dbus.SystemBus()
	if err != nil {
		return err
	}
	d.obj = d.conn.Object(iface, objectPath)
	log.Debugf("initialized %s", objectPath)

	return nil
}

func (d *dbusBase) getProperty(iface string) (interface{}, error) {
	variant, err := d.obj.GetProperty(iface)
	return variant.Value(), err
}

func (d *dbusBase) getStringProperty(iface string) (value string, err error) {
	prop, err := d.getProperty(iface)
	if err != nil {
		return
	}

	value, ok := prop.(string)
	if !ok {
		err = makeVariantTypeError(iface)
		return
	}
	return
}

func makeVariantTypeError(iface string) error {
	return fmt.Errorf("unexpected variant type for '%s'", iface)
}
