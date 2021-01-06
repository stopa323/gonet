package manifest

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	obj "github.com/stopa323/gonet/pkg/objects"
	yaml "gopkg.in/yaml.v3"
)

type Manifest struct {
	Connections []obj.ConnectionBase
}

func Load(path string) (mft *Manifest, err error) {
	log.Debugf("reading manifest: %s", path)

	var yamlFile []byte
	yamlFile, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("manifest read: %w", err)
	}

	mft, err = unmarshalManifest(yamlFile)
	if err != nil {
		return
	}

	log.Infof("loaded manifest: %s", path)
	return
}

func unmarshalManifest(in []byte) (*Manifest, error) {
	type manifest struct {
		Connections []yaml.Node `yaml:"connections"`
	}
	var (
		err error
		m   manifest
		mft Manifest
	)

	err = yaml.Unmarshal(in, &m)
	if err != nil {
		return nil, fmt.Errorf("unmarshal manifest: %w", err)
	}

	for _, c := range m.Connections {
		err = unmarshalConnection(&c, &mft)
		if err != nil {
			return nil, err
		}
	}
	return &mft, nil
}

func unmarshalConnection(conn *yaml.Node, m *Manifest) (err error) {
	type connectionBase struct {
		Type obj.ConnectionType `yaml:"type"`
	}
	var cb connectionBase

	err = conn.Decode(&cb)
	if err != nil {
		return fmt.Errorf("unmarshal base connection: %w", err)
	}

	switch t := cb.Type; t {
	case obj.TypeEthernet:
		var o obj.EthernetConnection
		err = conn.Decode(&o)
		if err != nil {
			return fmt.Errorf("unmarshal ethernet connection: %w", err)
		}

		m.Connections = append(m.Connections, &o)
		log.Debugf("unmarshalled new ethernet connection: %s", o.InterfaceName)
		return
	default:
		return fmt.Errorf("unmarshal connection: unknown connection type (%s)", t)
	}
}
