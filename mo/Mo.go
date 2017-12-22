package mo

import (
	"github.com/xu4wang/mano/mo/config"
	"github.com/xu4wang/mano/mo/log"
)

type Mo struct {
	*molog.Mo
	*config.Config
}

func New() *Mo {
	return &Mo{molog.New(), config.New("app.json")}
}
