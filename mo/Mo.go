package mo

import (
	"os"

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

func (m *Mo) PublishArgv1() {
	arg_num := len(os.Args)
	if arg_num > 1 {
		m.Publish(os.Args[1], os.Args)
	}
}
