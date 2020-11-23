package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/stopa323/gonet/pkg/nmdbus"
)

func main() {
	// log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetLevel(log.DebugLevel)

	var (
		nmp nmdbus.NetworkManagerProxy
		err error
	)

	nmp, err = nmdbus.NewNetworkManager()
	if err != nil {
		log.Error("could not initialize NetworkManager: ", err)
		panic(err)
	}

	devices, err := nmp.GetDevices()
	if err != nil {
		panic(err)
	}

	for _, d := range devices {
		var ifname string
		if ifname, err = d.GetPropertyInterface(); err != nil {
			log.Error(err)
		} else {
			log.Info(ifname)
		}
	}
}
