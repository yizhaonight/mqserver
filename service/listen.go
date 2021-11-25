package service

import (
	"gitea.com/lunny/tango"
	"kumact.com/gosdk/app/router"
)

type Config struct {
	Addr string `json:"addr"`
}

var (
	defaultConf = &Config{
		Addr: ":8999",
	}
)

func GetConfig() *Config {
	return defaultConf
}

func Run(addr string) {
	t := tango.Classic()
	methodList := []string{"POST", "PUT", "DELETE"}
	for _, v := range methodList {
		t.Route(v, "/*", Handler())
	}
	t.Run("0.0.0.0" + addr)
}

func Handler() tango.HandlerFunc {
	return func(ctx *tango.Context) {
		req := ctx.Req()
		token := router.GetRequestToken(req)
		ServerSubscribe(req.Method, req.URL, token, ctx)
		ctx.Action()
	}
}
