package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
	// "gopkg.in/yaml.v3"
)

type MetaData struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	URCap      urCap
	Artifacts  artifacts
}

type urCap struct {
	Name    string `yaml:"name"`
	Vendor  string `yaml:"vendor"`
	Version string `yaml:"version"`
}

type artifacts struct {
	Containers  []container  `yaml:"containers", omitempty`
	Osgibundles []osgibundle `yaml:"osgibundles", omitempty`
	Webarchives []webarchive `yaml:"webarchives", omitempty`
}
type container struct {
	Name         string        `yaml:"name"`
	Image        string        `yaml:"image"`
	Ports        []port        `yaml:"ports"`
	VolumeMounts []volumeMount `yaml:"volumeMounts"`
}

type port struct {
	Name          string `yaml:"name"`
	ContainerPort int    `yaml:"containerPort"`
	Protocol      string `yaml:"protocol", omitempty`
}

type volumeMount struct {
	MountPath string `yaml:"mountPath"`
}

type osgibundle struct {
	Name   string `yaml:"name"`
	Bundle string `yaml:"bundle"`
}

type webarchive struct {
	Name   string `yaml:"name"`
	Bundle string `yaml:"bundle"`
}

func readMetaData(path string) (*MetaData, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	t := &MetaData{}

	err = yaml.UnmarshalStrict(buff, t)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", path, err)
	}
	// fmt.Printf("--- t:\n%v\n\n", *t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	/*
		m := make(map[interface{}]interface{})
		err = yaml.Unmarshal([]byte(buff), &m)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- m:\n%v\n\n", m)

		d, err = yaml.Marshal(&m)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Printf("--- m dump:\n%s\n\n", string(d))
	*/
	return t, nil
}

func main() {
	data, err := readMetaData("urcap.yaml")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%v", *data)

	fmt.Printf("ApiVersion: %#v\n", data.ApiVersion)
	fmt.Printf("URCap.Name: %#v\n", data.URCap.Name)
	fmt.Printf("Containers.Image: %#v\n", data.Artifacts.Containers[0].Image)
}
