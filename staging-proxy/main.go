package main

/**
 * Edited but from https://hackernoon.com/writing-a-reverse-proxy-in-just-one-line-with-go-c1edfa78c84b
 */

import (
	"fmt"
	log "log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

func sendError(w http.ResponseWriter, req *http.Request, error error) {
	log.Printf("Proxy Error: %q", error)
	fmt.Fprintf(w, "404 Staging Site Not Found")
}

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	// create the reverse proxy and attach error handler
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ErrorHandler = sendError;

	// modify host and attach forwarded host header
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile(`(?:\.staging\.test\.co\.uk)$`)
	s := r.ReplaceAllString(req.Host, "");

	check, _ := regexp.MatchString(`^[A-z0-9-]*$`, s)
	if (!check){
		log.Printf("Illgeal Branch")
		return;
	}
	host := "http://app." + s + ".svc.cluster.local";

	log.Printf("Updated Host = %q\n", host);

	serveReverseProxy(host, res, req)
}

func main() {
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}