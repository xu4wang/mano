package config

import (
	"testing"

	"github.com/xu4wang/mano/mo/log"
)

func TestConfig_Get(t *testing.T) {
	d := molog.GetInstance()
	config := GetInstance("app.json")
	var s string
	var i int
	d.Info(config.GetConfig("ip", &s), s)
	d.Info(config.GetConfig("port", &i), i)
}
