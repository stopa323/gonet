package nm

import (
	log "github.com/sirupsen/logrus"
)

type NetworkManager struct {
	Connections ConnectionController
}

func NewNetworkManager() (nm *NetworkManager, err error) {
	cc, err := NewConnectionController()
	if err != nil {
		log.Error("NetworkManager init failed")
		return
	}

	nm = &NetworkManager{
		Connections: cc,
	}
	return
}
