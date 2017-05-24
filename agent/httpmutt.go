package agent

import (
	"HttpMutt/configuration"
	"github.com/elazarl/goproxy"
	"github.com/golang/glog"
	"net/http"
	"regexp"
	"strconv"
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
		r, _ := injectLatency(time.Duration(config.Delay), req, ctx)
		return blockRequest(config.ResponseStatusCode, r, ctx)
	}

	return injectLatency(time.Duration(config.Delay), req, ctx)
}

func blockRequest(statusCode int, req *http.Request, c *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Returning status code '%s' for '%s' on path '%s'", strconv.Itoa(statusCode), req.Host, req.URL.Path)
	return req, goproxy.NewResponse(req, goproxy.ContentTypeText, statusCode, "")
}

func injectLatency(delay time.Duration, req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Injecting %d milliseconds latency for '%s' on path '%s'", delay, req.Host, req.URL.Path)
	time.Sleep(time.Millisecond * delay)
	return req, nil
}
