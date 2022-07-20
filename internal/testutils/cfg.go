package testutils

import (
	"io/ioutil"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func loadFromYaml(t require.TestingT, path string, output interface{}) {
	buf, err := ioutil.ReadFile(path)
	require.NoError(t, err)

	err = yaml.Unmarshal([]byte(buf), &output)
	require.NoError(t, err)
}
