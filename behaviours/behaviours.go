package behaviour

import (
	"github.com/elazarl/goproxy"
	"github.com/golang/glog"
	"net/http"
	"strconv"
	"time"
)

func BlockRequest(statusCode int, req *http.Request, c *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Returning status code '%s' for '%s' on path '%s'", strconv.Itoa(statusCode), req.Host, req.URL.Path)
	return req, goproxy.NewResponse(req, goproxy.ContentTypeText, statusCode, "")
}

func InjectLatency(delay time.Duration, req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	glog.Infof("Injecting %d milliseconds latency for '%s' on path '%s'", delay, req.Host, req.URL.Path)
	time.Sleep(time.Millisecond * delay)
	return req, nil
}
