package agent

import (
	"HttpMutt/configuration"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Log(configLocation string, port string) {
	config, err := configuration.ParseConfig(configLocation)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on port " + port)
	setProxy(config, port)
}

func setProxy(config configuration.ConfigurationOptions, port string) {

	target := "https://requestb.in/wb042mwb"
	remote, err := url.Parse(target)
	if err != nil {
		panic(err)
	}

	_ = httputil.NewSingleHostReverseProxy(remote)

	r := mux.NewRouter()
	for _, endpoint := range config.Endpoints {
		r.HandleFunc(endpoint.Path, Handler)
	}

	http.Handle("/", r)
	http.ListenAndServe(":"+port, r)
}

func Handler(w http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Host)
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	_ = p
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = mux.Vars(r)["rest"]
		p.ServeHTTP(w, r)
	}
}
