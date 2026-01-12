package options

import (
	"os"

	"github.com/iihsannlearn/corscan/internal/logger"
	"github.com/projectdiscovery/goflags"
)

type Options struct {
	URLs    string
	URLList string

	Threads     int
	Timeout     int
	Payload     string
	PayloadList string
	Verbose     bool
	RandomAgent bool

	CustomOrigin string
	Proxy        string

	Output    string
	RateLimit int
	Delay     int
	Version   bool
	VerifySSL bool
}

const Version = "1.0.0"

func ParseOptions() *Options {
	opts := &Options{}

	flagSet := goflags.NewFlagSet()

	flagSet.SetDescription("Corscan is a professional CORS vulnerability scanner written in Go")

	flagSet.CreateGroup("input", "Input",
		flagSet.StringVarP(&opts.URLs, "url", "u", "", "Target URL for scanning (comma separated)"),
		flagSet.StringVarP(&opts.URLList, "list", "l", "", "File containing list of target URLs"),
		flagSet.StringVarP(&opts.CustomOrigin, "origin", "or", "", "Custom origin for testing"),
	)

	flagSet.CreateGroup("config", "Configuration",
		flagSet.IntVarP(&opts.Threads, "threads", "t", 25, "Number of concurrent threads"),
		flagSet.IntVarP(&opts.Timeout, "timeout", "T", 10, "Timeout request (seconds)"),
		flagSet.IntVarP(&opts.RateLimit, "rate-limit", "rl", 10, "Maximum requests per second"),
		flagSet.IntVarP(&opts.Delay, "delay", "d", 0, "Delay between requests (milliseconds)"),
		flagSet.StringVarP(&opts.Proxy, "proxy", "x", "", "HTTP Proxy URL (e.g. http://127.0.0.1:8080)"),
	)

	flagSet.CreateGroup("output", "Output",
		flagSet.StringVarP(&opts.Output, "output", "o", "", "File for saving scan results"),
	)

	flagSet.CreateGroup("misc", "Optimization",
		flagSet.BoolVarP(&opts.RandomAgent, "random-agent", "ra", false, "Enable Random User-Agent"),
		flagSet.BoolVarP(&opts.Verbose, "verbose", "v", false, "Verbose mode (print error & debug)"),
		flagSet.BoolVarP(&opts.Version, "version", "V", false, "Display application version"),
	)

	if err := flagSet.Parse(); err != nil {
		logger.Error("Failed parsing flags: %s", err)
		os.Exit(1)
	}

	if opts.Version {
		logger.Info("Corscan v%s", Version)
		os.Exit(0)
	}

	if opts.URLs == "" && opts.URLList == "" && !hasStdin() {
		logger.Error("No target URL provided! Use -u, -l or pipe stdin.")
		logger.Info("Run 'corscan -h' for help.")
		os.Exit(1)
	}

	return opts
}

func hasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return stat.Mode()&os.ModeCharDevice == 0
}
