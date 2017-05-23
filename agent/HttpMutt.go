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

	//target := "http://localhost:5000"
	//remote, err := url.Parse(target)
	//if err != nil {
	//	panic(err)
	//}

	//proxy := httputil.NewSingleHostReverseProxy(remote)

	r := mux.NewRouter()
	for _, endpoint := range config.Endpoints {
		//r.HandleFunc(endpoint.Path, handler(proxy))
		r.HandleFunc(endpoint.Path, func(w http.ResponseWriter, request *http.Request) {

			fullPath := request.Host + request.URL.Path
			remote, err := url.Parse(fullPath)
			if err != nil {
				panic(err)
			}

			proxy := httputil.NewSingleHostReverseProxy(remote)
			r.HandleFunc(endpoint.Path, handler(proxy))

		})
	}

	http.Handle("/", r)
	http.ListenAndServe(":"+port, r)
}

//func Handler(w http.ResponseWriter, request *http.Request) {
//	fmt.Println(request.Host)
//}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	_ = p
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Body)
		//r.URL.Path = mux.Vars(r)["rest"]
		p.ServeHTTP(w, r)
	}
}
