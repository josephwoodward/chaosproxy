package chaoskoko

import (
	"chaoskoko/behaviours"
	"chaoskoko/config"
	"fmt"
	"github.com/elazarl/goproxy"
	"github.com/golang/glog"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"
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

		go proxy.OnRequest(goproxy.ReqHostMatches(hostRegex), goproxy.UrlMatches(urlRegex)).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			return behaviourFactory(endpoint, req, ctx)
		})
	}

	glog.Fatal(http.ListenAndServe(":"+cfg.Config.Port, proxy))
}

func behaviourFactory(config config.Endpoint, req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Matched host '%s'", req.Host)

	if trafficInRange(config.Range) {
		glog.Infof("Request is within range of %s", strconv.Itoa(config.Range))
		if config.ResponseStatusCode > 0 {
			r, _ := behaviour.InjectLatency(time.Duration(config.Delay), req, ctx)
			return behaviour.BlockRequest(config.ResponseStatusCode, r, ctx)
		}

		return behaviour.InjectLatency(time.Duration(config.Delay), req, ctx)
	}

	return req, ctx.Resp
}

func trafficInRange(endpointRange int) bool {
	randomVal := rand.Intn(100-1) + 1
	return randomVal <= endpointRange
}
