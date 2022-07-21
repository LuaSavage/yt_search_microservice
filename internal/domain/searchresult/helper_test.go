package searchresult

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func LoadFromYaml(path string, output interface{}) error {
	buf, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(buf), output)

	if err != nil {
		return err
	}

	return nil
}
