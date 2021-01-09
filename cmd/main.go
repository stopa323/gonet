package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/stopa323/gonet/pkg/nm"
)

func main() {
	// log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetLevel(log.DebugLevel)

	// mft, err := manifest.Load("docs/nmcli-samples/eth-conn.yml")
	// if err != nil {
	// 	panic(err)
	// }

	var netManager *nm.NetworkManager
	netManager, err := nm.NewNetworkManager()
	if err != nil {
		panic(err)
	}

	// for _, c := range mft.Connections {
	// 	log.Debugf("Creating new connection: ")
	// 	if err = netManager.Connections.Create(c); err != nil {
	// 		log.Errorf("Failed to create connection : %s", err)
	// 		return
	// 	}
	// }

	c, e := netManager.Connections.List()
	if e != nil {
		panic(e)
	}
	for _, c2 := range c {
		log.Info(fmt.Sprintf("connection details:\n%s", c2.ToString()))
	}
}
