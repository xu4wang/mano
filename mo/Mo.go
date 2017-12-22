package mo

import (
	"github.com/xu4wang/mano/mo/config"
	"github.com/xu4wang/mano/mo/log"
	"github.com/xu4wang/mano/mo/pubsub"
)

type Mo struct {
	*molog.Mo
	*config.Config
	*pubsub.Router
}

func New() *Mo {
	return &Mo{molog.GetInstance(), config.GetInstance("app.json"), pubsub.GetInstance()}
}
