package behaviour

import (
	"chaosproxy/config"
	"github.com/elazarl/goproxy"
	"github.com/golang/glog"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var totalRequests int

func GetBehaviour(config config.Endpoint, req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	totalRequests++
	glog.Infof("Matched host '%s'", req.Host)
	glog.Infof("Creating request '%s'", strconv.Itoa(totalRequests))

	if !trafficInRange(config.RangeOrDefault()) {
		return req, ctx.Resp
	}

	glog.Infof("Request is within range of %s", config.RangeOrDefault())
	if config.ResponseStatusCode > 0 {
		r, _ := InjectLatency(time.Duration(config.Delay), req, ctx)
		return BlockRequest(config.ResponseStatusCode, r, ctx)
	}

	return InjectLatency(time.Duration(config.Delay), req, ctx)
}

func trafficInRange(endpointRange int) bool {
	randomVal := rand.Intn(100-1) + 1
	return randomVal <= endpointRange
}
