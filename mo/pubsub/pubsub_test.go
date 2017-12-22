package pubsub

import (
	"testing"

	"github.com/xu4wang/mano/mo/log"
)

func On1(d []string) {
	l := molog.GetInstance()
	l.Info("On1 ", d)
}

func OnAll(d []string) {
	l := molog.GetInstance()
	l.Info("OnAll ", d)
}

func TestPubsub_1(t *testing.T) {
	r := GetInstance()
	r.Subscribe("On1", On1)
	r.Subscribe("On2", OnAll)
	r.Subscribe("On1", OnAll)
	r.Subscribe("On3", OnAll)
	r.Publish("On1", []string{"On1 Data"})
	r.Publish("On2", []string{"On2 Data"})
	r.Publish("On3", []string{"On3 Data"})
	r.Publish("On4", []string{"On4 Data"})

}
