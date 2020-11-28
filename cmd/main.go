package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stopa323/gonet/pkg/nm"
)

func main() {
	// log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetLevel(log.DebugLevel)

	var (
		n   *nm.NetworkManager
		err error
	)

	n, err = nm.NewNetworkManager()
	if err != nil {
		panic(err)
	}

	ethConnection := nm.EthernetConnection{
		InterfaceName: "enp0s8",
		MTU:           1900,
	}
	if err = n.CreateConnection(&ethConnection); err != nil {
		panic(err)
	}
}
