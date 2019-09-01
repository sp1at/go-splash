package args

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	Url  		string
	Endpoints
}

type Endpoints struct {
	Har 	bool
	Html 	bool
	Jpeg 	bool
	Json	bool
	Png 	bool
}

func ProcessArgs() *Args {
	args := Args{}
	var endpoints = Endpoints{}

	flag.BoolVar(&endpoints.Har, "har", false, "Return HTML output")

	flag.BoolVar(&endpoints.Html, "html", false, "Return HTML output")

	// both 'jpeg' and 'jpg' are accepted because who can remember which one it is??
	flag.BoolVar(&endpoints.Jpeg, "jpeg", false, "Return HTML output")
	flag.BoolVar(&endpoints.Jpeg, "jpg", false, "Return HTML output")

	flag.BoolVar(&endpoints.Json, "json", false, "Return HTML output")

	flag.BoolVar(&endpoints.Png,  "png", false, "Return HTML output")

	flag.Parse()

	args.Endpoints = endpoints
	args.Url = flag.Arg(0)
	return &args
}

func init() {
	flag.Usage = func() {
		h := "Golang Splash Wrapper\n\n"

		h += "Usage:\n"
		h += "  go-splash [flags] url\n"
		h += "  * default option is --json\n\n"

		h += "Example:\n"
		h += "  go-splash --jpg https://google.com\n\n"

		h += "Flags: (default: --har --html --jpg --json)\n"
		h += "-har"
		h += "-html"
		h += "-jpg"
		h += "-json"
		h += "-png"

		_, _ = fmt.Fprintf(os.Stderr, h)
	}
}
