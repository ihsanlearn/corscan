package corscan

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/iihsannlearn/corscan/internal/options"
	"github.com/iihsannlearn/corscan/pkg/useragent"
)

type Result struct {
	VulnerableUrl string
	Origin        string
}

func ScanUrl(client *http.Client, target string, origins []string, opts *options.Options) []*Result {
	var results []*Result

	// Add dynamic payloads (shifting)
	expandedOrigins := append([]string{}, origins...)
	if opts.CustomOrigin != "" {
		expandedOrigins = append(expandedOrigins, opts.CustomOrigin)
	}

	for _, origin := range expandedOrigins {
		// Normal Scan
		if res := checkOrigin(client, target, origin, "", opts); res != nil {
			results = append(results, res)
		}

		// HPP Scan (Experimental)
		if res := checkOrigin(client, target, origin, "https://google.com", opts); res != nil {
			res.VulnerableUrl += " (HPP)"
			results = append(results, res)
		}
	}

	return results
}

func checkOrigin(client *http.Client, target, origin, secondOrigin string, opts *options.Options) *Result {
	req, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return nil
	}

	req.Header.Set("Origin", origin)
	if secondOrigin != "" {
		req.Header.Add("Origin", secondOrigin)
	}

	if opts.RandomAgent {
		req.Header.Set("User-Agent", useragent.GetRandom())
	}

	// Anti-WAF Headers (Reverse Header)
	req.Header.Set("X-Forwarded-For", fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255)))
	req.Header.Set("X-Real-IP", fmt.Sprintf("10.0.%d.%d", rand.Intn(255), rand.Intn(255)))

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	acao := resp.Header.Get("Access-Control-Allow-Origin")
	acac := resp.Header.Get("Access-Control-Allow-Credentials")

	if acao == origin {
		msg := fmt.Sprintf("Reflected Origin: %s", acao)
		if acac == "true" {
			msg += " (with Credentials!)"
		}
		return &Result{
			VulnerableUrl: fmt.Sprintf("%s [Origin: %s] %s", target, origin, msg),
			Origin:        origin,
		}
	} else if acao == "null" && origin == "null" {
		return &Result{
			VulnerableUrl: fmt.Sprintf("%s [Origin: null] (Reflected null)", target),
			Origin:        origin,
		}
	}
	return nil
}
