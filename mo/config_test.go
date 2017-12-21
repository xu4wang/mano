package mo

import "testing"

func TestConfig_Get(t *testing.T) {
	d := &Mo{}
	config := New("app.json")
	var s string
	var i int
	d.Info(config.Get("ip", &s), s)
	d.Info(config.Get("port", &i), i)
}
