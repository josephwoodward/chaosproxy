package agent

import (
	"HttpMutt/behaviours"
	"HttpMutt/configuration"
	"github.com/elazarl/goproxy"
	"github.com/golang/glog"
	"net/http"
	"regexp"
	"time"
)

var config configuration.ConfigurationOptions

func Proxy(configLocation string) {
	cfg, err := configuration.ParseYml(configLocation)
	if err != nil {
		glog.Fatal(err)
	}

	config = cfg

	glog.Infof("Starting proxy, listening on port %s", config.Config.Port)
	setProxy()
}

func setProxy() {

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	for _, endpoint := range config.Endpoints {
		hostRegex, err := regexp.Compile(endpoint.Host)
		if err != nil {
			glog.Error("Invalid regex format on endpoint.host", err)
		}

		urlRegex, err := regexp.Compile(endpoint.Url)
		if err != nil {
			glog.Error("Invalid regex format on endpoint.url", err)
		}

		go proxy.OnRequest(goproxy.ReqHostMatches(hostRegex), goproxy.UrlMatches(urlRegex)).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			return routeFactory(endpoint, req, ctx)
		})
	}

	glog.Fatal(http.ListenAndServe(":"+config.Config.Port, proxy))
}

func routeFactory(config configuration.Endpoint, req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Matched host '%s'", req.Host)

	if config.ResponseStatusCode > 0 {
		r, _ := behaviour.InjectLatency(time.Duration(config.Delay), req, ctx)
		return behaviour.BlockRequest(config.ResponseStatusCode, r, ctx)
	}

	return behaviour.InjectLatency(time.Duration(config.Delay), req, ctx)
}
