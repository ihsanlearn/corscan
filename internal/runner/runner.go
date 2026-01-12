package runner

import (
	"sync"
	"time"

	"github.com/iihsannlearn/corscan/internal/input"
	"github.com/iihsannlearn/corscan/internal/logger"
	"github.com/iihsannlearn/corscan/internal/options"
	"github.com/iihsannlearn/corscan/internal/output"
	"github.com/iihsannlearn/corscan/internal/payloads"
	"github.com/iihsannlearn/corscan/pkg/corscan"
	"github.com/iihsannlearn/corscan/pkg/httputils"
	"github.com/iihsannlearn/corscan/pkg/utils"
)

func Run(opts *options.Options) {
	logger.PrintBanner()

	inputProvider := input.NewInputProvider(opts)

	writer, err := output.NewWriter(opts.Output)
	if err != nil {
		logger.Error("%s", err)
		return
	}
	defer writer.Close()

	client := httputils.NewScannerClient(opts)
	urls := inputProvider.StreamURLs()

	var wg sync.WaitGroup

	var ticker *time.Ticker
	if opts.RateLimit > 0 {
		interval := time.Second / time.Duration(opts.RateLimit)
		ticker = time.NewTicker(interval)

		logger.Info("Rate limit set to %d requests per second", opts.RateLimit)
	}

	var targetPayloads []string
	if opts.PayloadList != "" {
		logger.Info("Using payload list %s", opts.PayloadList)
		var err error
		targetPayloads, err = utils.ReadFileLines(opts.PayloadList)
		if err != nil {
			logger.Error("%s", err)
			return
		}
	} else if opts.Payload != "" {
		targetPayloads = []string{opts.Payload}
	} else {
		logger.Info("Using default payloads")
		targetPayloads = payloads.GetDefaultPayloads()
	}

	for i := 0; i < opts.Threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urls {
				if ticker != nil {
					<-ticker.C
				}

				if opts.Delay > 0 {
					time.Sleep(time.Duration(opts.Delay) * time.Millisecond)
				}

				url, notOk := httputils.ProbeURL(url, opts.Timeout, opts)
				if notOk != "" {
					if opts.Verbose {
						logger.Dead("%s", notOk)
					}
					continue
				}

				results := corscan.ScanUrl(client, url, targetPayloads, opts)

				for _, result := range results {
					if result != nil {
						logger.Vulnerable("%s", result.VulnerableUrl)

						writer.Write(result.VulnerableUrl)
					} else {
						if opts.Verbose {
							logger.NotVulnerable("%s", url)
						}
					}
				}
			}
		}()
	}

	wg.Wait()

	if ticker != nil {
		ticker.Stop()
	}

	logger.Info("Scan completed")

}
