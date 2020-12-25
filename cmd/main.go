package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stopa323/gonet/pkg/nm"
	"github.com/stopa323/gonet/pkg/nm/connection"
)

func main() {
	// log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetLevel(log.DebugLevel)

	var (
		netManager *nm.NetworkManager
		err        error
	)

	netManager, err = nm.NewNetworkManager()
	if err != nil {
		panic(err)
	}

	ethConnIntent := connection.EthernetConnection{
		InterfaceName: "enp0s8",
		MTU:           1600,
	}
	if err = netManager.Connections.Create(&ethConnIntent); err != nil {
		panic(err)
	}
}
