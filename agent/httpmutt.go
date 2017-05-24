package agent

import (
	"HttpMutt/configuration"
	"github.com/elazarl/goproxy"
	"github.com/golang/glog"
	"net/http"
	"regexp"
	"time"
)

const delay = 5000

func Log(configLocation string, port string, outputLog bool) {
	config, err := configuration.ParseConfig(configLocation)
	if err != nil {
		glog.Fatal(err)
	}

	glog.Error("Hello world!")
	glog.Infof("Starting proxy, listening on port %s", port)
	setProxy(config, port)

}

func setProxy(config configuration.ConfigurationOptions, port string) {

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	for _, endpoint := range config.Endpoints {
		regex, err := regexp.Compile(endpoint.Url)
		if err != nil {
			glog.Warning("Invalid regex format on URL.\n")
		}

		go proxy.OnRequest(goproxy.DstHostIs(endpoint.Host), goproxy.UrlMatches(regex)).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			return routeFactory(endpoint, req, ctx)
		})
	}

	glog.Fatal(http.ListenAndServe(":"+port, proxy))
}

func routeFactory(config configuration.Endpoint, req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Matched host '%s'", req.Host)
	return injectLatency(config, req, ctx)
}

func blockRequest(r *http.Request, c *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	if h, _, _ := time.Now().Clock(); h >= 8 && h <= 17 {
		return r, goproxy.NewResponse(r, goproxy.ContentTypeText, http.StatusForbidden, "Don't waste your time!")
	}
	return r, nil
}

func injectLatency(config configuration.Endpoint, req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Injecting %d milliseconds latency for '%s' on path '%s'", delay, req.Host, req.URL.Path)
	time.Sleep(time.Millisecond * delay)
	return req, nil
}
