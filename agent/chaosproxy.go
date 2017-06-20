package chaosproxy

import (
	"chaosproxy/behaviours"
	"chaosproxy/config"
	"fmt"
	"net/http"
	"regexp"

	"github.com/elazarl/goproxy"
	"github.com/golang/glog"
)

var cfg config.ConfigurationOptions

func Proxy(args config.CommandLineArgs) {
	c, err := config.ParseYml(args.ConfigLocation)
	if err != nil {
		glog.Fatal(err)
	}

	cfg = c

	fmt.Printf("Starting proxy, listening on port :%s\n", cfg.Config.Port)
	setProxy()
}

func setProxy() {

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	for _, endpoint := range cfg.Endpoints {
		hostRegex, err := regexp.Compile(endpoint.Host)
		if err != nil {
			glog.Error("Invalid regex format on endpoint.host", err)
		}

		urlRegex, err := regexp.Compile(endpoint.Url)
		if err != nil {
			glog.Error("Invalid regex format on endpoint.url", err)
		}

		behavioursEnabled := cfg.Config.IsEnabled()
		go proxy.OnRequest(goproxy.ReqHostMatches(hostRegex), goproxy.UrlMatches(urlRegex)).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			if !behavioursEnabled {
				return behaviour.ForwardRequest(req, ctx)
			}

			return behaviour.GetBehaviour(endpoint, req, ctx)
		})
	}

	glog.Fatal(http.ListenAndServe(":"+cfg.Config.Port, proxy))
}
