package action

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/giantswarm/microerror"
)

func All(cluster string) ([]string, error) {
	var actions []string
	{
		_, err := os.Stat(fmt.Sprintf("cmd/%s", cluster))
		if os.IsNotExist(err) {
			err = os.Mkdir(fmt.Sprintf("cmd/%s", cluster), os.ModePerm)
			if err != nil {
				return nil, microerror.Mask(err)
			}
		} else {
			return nil, microerror.Mask(err)
		}

		path, err := filepath.Abs(fmt.Sprintf("cmd/%s", cluster))
		if err != nil {
			return nil, microerror.Mask(err)
		}

		files, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		for _, file := range files {
			if !file.IsDir() {
				continue
			}
			if !strings.HasPrefix(file.Name(), "ac") {
				continue
			}

			actions = append(actions, file.Name())
		}
	}

	return actions, nil
}
