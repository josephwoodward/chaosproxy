package behaviour

import (
	"net/http"
	"strconv"
	"time"

	"github.com/elazarl/goproxy"
	"github.com/golang/glog"
)

func BlockRequest(statusCode int, req *http.Request, c *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Returning status code '%s' for '%s' on path '%s'", strconv.Itoa(statusCode), req.Host, req.URL.Path)
	return req, goproxy.NewResponse(req, goproxy.ContentTypeText, statusCode, "")
}

func InjectLatency(delay time.Duration, req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Injecting %d milliseconds latency for '%s' on path '%s'", delay, req.Host, req.URL.Path)
	if delay > 0 {
		return req, nil
	}

	time.Sleep(time.Millisecond * delay)
	glog.Infof("Delay elapsed, returning")
	glog.Infof("--- END ---")
	return req, nil
}

func ForwardRequest(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Pass-through request for '%s' on path '%s'", req.Host, req.URL.Path)
	return req, nil
}
