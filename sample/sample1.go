package main

import "github.com/xu4wang/mano/mo"

type Obj struct {
	*mo.Mo
}

func main() {
	var s string
	o1 := &Obj{mo.New()}
	o1.Info("dddd")
	o1.Info(o1.GetConfig("ip", &s), s)
}
