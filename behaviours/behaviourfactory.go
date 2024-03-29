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

	trafficRange := config.RangeOrDefault()
	if !trafficInRange(trafficRange) {
		glog.Infof("Request out of range of %s", strconv.Itoa(trafficRange))
		return req, ctx.Resp
	}

	glog.Infof("Request in range of %s", strconv.Itoa(trafficRange))
	delay := time.Duration(config.Delay)

	if config.ResponseStatusCode > 0 {
		r, _ := InjectLatency(delay, req, ctx)
		return BlockRequest(config.ResponseStatusCode, r, ctx)
	}

	return InjectLatency(delay, req, ctx)
}

func trafficInRange(endpointRange int) bool {
	randomVal := rand.Intn(100-1) + 1
	return randomVal <= endpointRange
}
