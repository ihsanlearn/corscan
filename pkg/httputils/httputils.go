package httputils

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/iihsannlearn/corscan/internal/options"
	"github.com/iihsannlearn/corscan/pkg/useragent"
)

func NewScannerClient(opts *options.Options) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConnsPerHost: opts.Threads,
	}

	if opts.Proxy != "" {
		proxyURL, err := url.Parse(opts.Proxy)
		if err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}

	client := &http.Client{
		Timeout:   time.Duration(opts.Timeout) * time.Second,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return client
}

func ProbeURL(target string, timeout int, opts *options.Options) (string, string) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	
	if opts.Proxy != "" {
		proxyURL, err := url.Parse(opts.Proxy)
		if err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}

	client := &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: transport,
	}
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return target, err.Error()
	}

	if opts.RandomAgent {
		req.Header.Set("User-Agent", useragent.GetRandom())
	}
	
	resp, err := client.Do(req)
	if err != nil {
		return target, err.Error()
	}
	defer resp.Body.Close()
	return target, ""
}
