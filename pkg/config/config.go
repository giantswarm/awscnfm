package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/ghodss/yaml"
)

func Cluster(scope string, cluster string) string {
	// Here we mean to prefer something like the cluster ID given via the
	// process environment.
	if cluster != "" {
		return cluster
	}

	// At this point the cluster ID is not provided with the process environment
	// so we read the config file information from the local file system.
	v := struct {
		Cluster *string `yaml:"cluster"`
	}{}

	mustFromFile(scope, &v)

	if v.Cluster == nil {
		return ""
	}

	return *v.Cluster
}

func SetCluster(scope string, cluster string) {
	// The config file structure here can be extended with more fields.
	v := struct {
		Cluster string `yaml:"cluster"`
	}{}

	mustFromFile(scope, &v)

	v.Cluster = cluster

	mustToFile(scope, &v)
}

func mustFromFile(scope string, v interface{}) {
	b, err := ioutil.ReadFile(mustName(scope))
	if os.IsNotExist(err) {
		return
	} else if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(b, v)
	if err != nil {
		panic(err)
	}
}

func mustToFile(scope string, v interface{}) {
	b, err := yaml.Marshal(v)
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Dir(mustName(scope)), os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(mustName(scope), b, 0644) // nolint:gosec
	if err != nil {
		panic(err)
	}
}

// mustName returns the config file name as absolute path according to the
// current OS User known to the running process.
//
//     ~/.config/awscnfm/cl001.yaml
//
func mustName(scope string) string {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	return filepath.Join(u.HomeDir, fmt.Sprintf(".config/awscnfm/%s.yaml", scope))
}
