package config

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	t.Run("error with cached search result", func(t *testing.T) {
		cfg, err := GetConfig()
		log.Printf("%+v\n", cfg)
		assert.NoError(t, err)
	})
}
