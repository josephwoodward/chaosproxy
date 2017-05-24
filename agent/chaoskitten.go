package chaoskitten

import (
	"chaos-kitten/behaviours"
	"chaos-kitten/config"
	"fmt"
	"github.com/elazarl/goproxy"
	"github.com/golang/glog"
	"net/http"
	"regexp"
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
			return routeFactory(endpoint, req, ctx)
		})
	}

	glog.Fatal(http.ListenAndServe(":"+cfg.Config.Port, proxy))
}

func routeFactory(config config.Endpoint, req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Matched host '%s'", req.Host)

	if config.ResponseStatusCode > 0 {
		r, _ := behaviour.InjectLatency(time.Duration(config.Delay), req, ctx)
		return behaviour.BlockRequest(config.ResponseStatusCode, r, ctx)
	}

	return behaviour.InjectLatency(time.Duration(config.Delay), req, ctx)
}
