package config

import (
	"os"

	"github.com/LeviyLokotb/light-automata/pkg/materials"
	"gopkg.in/yaml.v3"
)

type SuperConfig struct {
	Conf   Config   `yaml:"config"`
	OMConf OMConfig `yaml:"scene"`
}

func LoadSuperConfigFromYaml(path string) (*Config, *materials.ObjectsManager, error) {
	sconf, err := loadSuperConfigFromYaml(path)
	if err != nil {
		return nil, nil, err
	}

	objects, err := parseObjects(sconf.OMConf)
	if err != nil {
		return nil, nil, err
	}

	return &sconf.Conf, materials.NewObjectsManager(objects), nil
}

func loadSuperConfigFromYaml(path string) (*SuperConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf SuperConfig
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
