package main

import "github.com/xu4wang/mano/mo"

type Obj struct {
	*mo.Mo
}

func OnStart(d []string) {
	o1 := &Obj{mo.New()}
	o1.Info("OnStart ", d)
}

func main() {
	var s string
	o1 := &Obj{mo.New()}
	o1.Info("dddd")
	o1.Info(o1.GetConfig("ip", &s), s)
	o1.Subscribe("start", OnStart)
	o1.Publish("start", []string{"On Start should be running?"})
	o1.Publish("start1", []string{"On Start should NOT be running?"})
}
