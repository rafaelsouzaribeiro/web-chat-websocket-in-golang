package factory

import (
	"github.com/rafaelsouzaribeiro/web-chat-websocket-in-golang/configs"
)

type Factory struct {
	types string
	Conf  *configs.Conf
}

func NewFactory(types string, conf *configs.Conf) *Factory {
	return &Factory{
		types: types,
		Conf:  conf,
	}
}
