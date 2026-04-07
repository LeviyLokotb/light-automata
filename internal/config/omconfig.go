package config

import (
	"os"

	"github.com/LeviyLokotb/light-automata/pkg/materials"
	"gopkg.in/yaml.v3"
)

type OMConfig struct {
	Objects []ObjectConfig `yaml:"objects"`
}

type ObjectConfig struct {
	Shape    string `yaml:"shape"`
	Material string `yaml:"material"`
	Params   Params `yaml:"params"`
}

type Params map[string]int

func LoadSceneFromYaml(path string) (*materials.ObjectsManager, error) {
	omconf, err := loadOMConfigFromYaml(path)
	if err != nil {
		return nil, err
	}

	objects, err := parseObjects(omconf)
	if err != nil {
		return nil, err
	}

	return materials.NewObjectsManager(objects), nil
}

func loadOMConfigFromYaml(path string) (OMConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return OMConfig{}, err
	}

	var conf OMConfig
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return OMConfig{}, err
	}
	return conf, nil
}
